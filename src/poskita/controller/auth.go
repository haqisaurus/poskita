package controller

import (
	"fmt"

	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/haqisaurus/poskita/config"
	"github.com/haqisaurus/poskita/dto"
	"github.com/haqisaurus/poskita/entity"
	"github.com/haqisaurus/poskita/util"
	"github.com/speps/go-hashids"
)

func GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return c.Status(fiber.StatusOK).JSON(util.GenerateResponse(claims))

}

func GetCompanyList(c *fiber.Ctx) error {
	claims := util.GetUser(c)

	db := config.Connect()
	defer config.Close(db)

	companies := []entity.AuthCompany{}
	result := db.Joins("inner JOIN auth_user_company_role aucr on aucr.company_id = auth_company.id ").Where("aucr.user_id", claims["id"]).Find(&companies)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(util.GenerateError(util.ErrUnauthorized, result.Error.Error()))
	}
	var response []dto.CompanyListRs
	for _, company := range companies {
		response = append(response, dto.CompanyListRs{
			ID:   company.ID,
			Name: company.Name,
		})
	}
	return c.Status(fiber.StatusOK).JSON(util.GenerateResponse(response))

}

func PostLoginCompany(c *fiber.Ctx) error {
	db := config.Connect()
	defer config.Close(db)

	var req dto.LoginCompanyRs
	var res dto.LoginRs
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrInvalidInput, err.Error()))
	}
	claimsToken := util.GetUser(c)
	user := entity.AuthUser{}
	result := db.Joins("inner JOIN auth_user_company_role aucr on aucr.company_id = auth_user.id ").Where("aucr.user_id", claimsToken["id"]).Where("company_id = ?", req.CompanyID).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(util.GenerateError(util.ErrUnauthorized, result.Error.Error()))
	}
	company := entity.AuthCompany{}
	result = db.Where("id = ?", req.CompanyID).First(&company)
	expiration := time.Now().Add(time.Minute * 60).Unix()
	claims := jwt.MapClaims{
		"email":       user.Email,
		"firstName":   user.FirstName,
		"lastName":    user.LastName,
		"username":    user.Username,
		"id":          user.ID,
		"companyId":   company.ID,
		"companyName": company.Name,
		"exp":         expiration,
	}
	privateKey, err := util.LoadPrivateKey(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		fmt.Printf("Error loading private key: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrInternal, err.Error()))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	t, errT := token.SignedString(privateKey)
	if errT != nil {
		log.Fatal(errT.Error())
	}
	// create refresh token
	hd := hashids.NewData()
	hd.Salt = os.Getenv("HASH_KEY")
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)

	currentTime := int(time.Now().UnixNano() / int64(time.Millisecond))
	refreshToken, _ := h.Encode([]int{int(user.ID), int(req.CompanyID), currentTime})

	res.Token = t
	res.RefreshToken = refreshToken
	res.Expiration = expiration
	return c.Status(fiber.StatusOK).JSON(util.GenerateResponse(res))

}
func PostLogin(c *fiber.Ctx) error {
	var req dto.LoginRq
	var res dto.LoginRs
	db := config.Connect()
	defer config.Close(db)

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrInvalidInput, err.Error()))
	}

	user := entity.AuthUser{}
	result := db.Where("username = ? OR email = ?", req.Username, req.Password).Where("is_active =  ? ", true).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(util.GenerateError(util.ErrUnauthorized, result.Error.Error()))
	}

	privateKey, err := util.LoadPrivateKey(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		fmt.Printf("Error loading private key: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrInternal, err.Error()))
	}

	expiration := time.Now().Add(time.Minute * 60).Unix()
	claims := jwt.MapClaims{
		"email":     user.Email,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"id":        user.ID,
		"exp":       expiration,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	t, errT := token.SignedString(privateKey)
	if errT != nil {
		log.Fatal(errT.Error())
	}
	userAudit := entity.AuthAudit{
		CreatedBy: user.Username,
		CreatedAt: time.Now(),
		UpdatedBy: user.Username,
		UserID:    user.ID,
		Username:  user.Username,
		UpdatedAt: time.Now(),
		IPAddress: strings.Join(c.IPs(), ","),
		ErrorCode: "0",
		ErrorDesc: "SUCESS",
	}

	resultInsert := db.Create(&userAudit)
	if resultInsert.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrInternal, result.Error.Error()))
	}

	res.Token = t
	res.RefreshToken = ""
	res.Expiration = expiration
	return c.Status(fiber.StatusOK).JSON(util.GenerateResponse(res))

}

func RefreshToken(c *fiber.Ctx) error {
	var req dto.RefreshTokenRq
	var res dto.LoginRs
	db := config.Connect()
	defer config.Close(db)

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrInvalidInput, err.Error()))
	}

	hd := hashids.NewData()
	hd.Salt = os.Getenv("HASH_KEY")
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)
	ids, e := h.DecodeWithError(req.RefreshToken)
	if e != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrInvalidInput, e.Error()))
	}
	userID := ids[0]
	companyID := ids[1]
	exp := time.Unix(int64(ids[2]), 0)
	fmt.Println(exp)
	// cek apakah refresh token sudah expired
	if exp.Before(time.Now()) {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrUnauthorized, "Refresh token is expired"))
	}

	userCompanyRole := entity.AuthUserCompanyRole{}
	result := db.Where("user_id = ? and role_id = ?", userID, companyID).First(&userCompanyRole)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(util.GenerateError(util.ErrUnauthorized, result.Error.Error()))
	}

	privateKey, err := util.LoadPrivateKey(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		fmt.Printf("Error loading private key: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrInternal, err.Error()))
	}
	user := entity.AuthUser{}
	result = db.Where("id = ? ", userID).First(&user)

	company := entity.AuthCompany{}
	result = db.Where("id = ? ", companyID).First(&company)

	expiration := time.Now().Add(time.Minute * 60).Unix()
	claims := jwt.MapClaims{
		"email":     user.Email,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"id":        user.ID,
		"companyId":   company.ID,
		"companyName": company.Name,
		"exp":       expiration,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	t, errT := token.SignedString(privateKey)
	if errT != nil {
		log.Fatal(errT.Error())
	}

	// create refresh token
	hd.MinLength = 30
	currentTime := int(time.Now().UnixNano() / int64(time.Millisecond))
	refreshToken, _ := h.Encode([]int{int(user.ID), int(company.ID), currentTime})

	res.Token = t
	res.RefreshToken = refreshToken
	res.Expiration = expiration

	return c.Status(fiber.StatusOK).JSON(util.GenerateResponse(res))

}

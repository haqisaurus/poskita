package controller

import (
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/haqisaurus/poskita/config"
	"github.com/haqisaurus/poskita/dto"
	"github.com/haqisaurus/poskita/entity"
	"github.com/haqisaurus/poskita/util"
)

func AddCompany(c *fiber.Ctx) error {
	db := config.Connect()
	defer config.Close(db)
	var req dto.AuthCompanyRq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrInvalidInput, err.Error()))
	}

	claimsToken := util.GetUser(c)

	company := entity.AuthCompany{
		CreatedBy:   claimsToken["username"].(string),
		CreatedAt:   time.Now(),
		UpdatedBy:   claimsToken["username"].(string),
		UpdatedAt:   time.Now(),
		UserID:      uint64(claimsToken["id"].(float64)),
		Name:        req.Name,
		Description: req.Description,
	}

	resultInsert := db.Create(&company)
	if resultInsert.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrInternal, resultInsert.Error.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(util.GenerateResponse(dto.InsertSuccess{
		ID: company.ID,
	}))
}
func EditCompany(c *fiber.Ctx) error {
	db := config.Connect()
	defer config.Close(db)
	var req dto.AuthCompanyRq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrInvalidInput, err.Error()))
	}

	company := entity.AuthCompany{}
	db.Model(&company).Find(&company, req.ID)
	if company.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrNotFound, "Entity not found"))
	}
	db.Model(&company).Updates(entity.AuthCompany{Name: req.Name, Description: req.Description})
 	return c.Status(fiber.StatusOK).JSON(util.GenerateResponse(dto.InsertSuccess{
		ID: company.ID,
	}))

}
func ListCompany(c *fiber.Ctx) error {
	page := c.Query("page", "1")
	size := c.Query("size", "10")
	offset, _ := strconv.Atoi(page)
	limit, _ := strconv.Atoi(size)
	if offset > 1 {
		offset = (offset - 1) * limit
	} else {
		offset = (offset - 1)
	}
	keyword := c.Query("keyword", "")
	db := config.Connect()
	defer config.Close(db)
	companies := []entity.AuthCompany{}
	query := db.Model(&entity.AuthCompany{})
	if keyword != "" {
		querySearch := "%" + strings.ToUpper(keyword) + "%"
		query.Where("UPPER(name) LIKE ? OR UPPER(description) LIKE ?", querySearch, querySearch)
	}

	column := c.Query("column", "id ")
	switch column {
	case "name":
		column = "name "
	default:
		column = "id "
	}

	order := c.Query("order", "DESC")
	switch order {
	case "ascend":
		order = "ASC"
	case "descend":
		order = "DESC"
	default:
		order = "DESC"
	}
	query.Order(column + order)
	var count int64
	query.Count(&count)
	query.Limit(limit).Offset(offset).Find(&companies)

	contents := []dto.AuthCompanyRs{}
	for _, company := range companies {
		contents = append(contents, dto.AuthCompanyRs{
			ID: company.ID,
			Name: company.Name,
			Description: company.Description,
			CreatedAt: company.CreatedAt,
			CreatedBy: company.CreatedBy,
			UpdatedAt: company.UpdatedAt,
			UpdatedBy: company.UpdatedBy,
		})
	}
	currentPage, _ := strconv.Atoi(page)
	return c.Status(fiber.StatusOK).JSON(util.GenerateResponse(dto.PaginationRs{
		Content: contents,
		Total:   int(count),
		Page:    currentPage,
		PerPage: limit,
	}))

}
func DetailCompany(c *fiber.Ctx) error {
	id_ := c.Params("id", "0")
	id, _ := strconv.Atoi(id_)

	company := entity.AuthCompany{}
	db := config.Connect()
	defer config.Close(db)
	db.Model(&company).Find(&company, id)
	if company.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrNotFound, "Entity not found"))
	}
	response := dto.AuthCompanyRs{
		ID: company.ID,
		Name: company.Name,
		Description: company.Description,
		CreatedAt: company.CreatedAt,
		CreatedBy: company.CreatedBy,
		UpdatedAt: company.UpdatedAt,
		UpdatedBy: company.UpdatedBy,
	}
	return c.Status(fiber.StatusOK).JSON(util.GenerateResponse(response))

}

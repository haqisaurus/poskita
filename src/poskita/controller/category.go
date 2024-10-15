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

func AddCategory(c *fiber.Ctx) error {
	db := config.Connect()
	defer config.Close(db)
	var req dto.ProductCategoryRq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrInvalidInput, err.Error()))
	}

	claimsToken := util.GetUser(c)

	category := entity.ProductCategory{
		CreatedBy:   claimsToken["username"].(string),
		CreatedAt:   time.Now(),
		UpdatedBy:   claimsToken["username"].(string),
		UpdatedAt:   time.Now(),
		CompanyID:   uint64(claimsToken["companyId"].(float64)),
		Name:        req.Name,
		Description: req.Description,
	}

	resultInsert := db.Create(&category)
	if resultInsert.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrInternal, resultInsert.Error.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(util.GenerateResponse(dto.InsertSuccess{
		ID: category.ID,
	}))
}

func EditCategory(c *fiber.Ctx) error {
	db := config.Connect()
	defer config.Close(db)
	var req dto.ProductCategoryRq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrInvalidInput, err.Error()))
	}

	category := entity.ProductCategory{}
	db.Model(&category).Find(&category, req.ID)
	if category.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrNotFound, "Entity not found"))
	}
	db.Model(&category).Updates(entity.AuthCompany{Name: req.Name, Description: req.Description})
 	return c.Status(fiber.StatusOK).JSON(util.GenerateResponse(dto.InsertSuccess{
		ID: category.ID,
	}))

}
func ListCategory(c *fiber.Ctx) error {
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
	claimsToken := util.GetUser(c)

	categories := []entity.ProductCategory{}
	query := db.Model(&entity.ProductCategory{})
	query.Where("company_id = ?", claimsToken["companyId"])
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
	query.Limit(limit).Offset(offset).Find(&categories)

	contents := []dto.ProductCategoryRs{}
	for _, category := range categories {
		contents = append(contents, dto.ProductCategoryRs{
			ID: category.ID,
			Name: category.Name,
			Description: category.Description,
			CreatedAt: category.CreatedAt,
			CreatedBy: category.CreatedBy,
			UpdatedAt: category.UpdatedAt,
			UpdatedBy: category.UpdatedBy,
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
func DetailCategory(c *fiber.Ctx) error {
	id_ := c.Params("id", "0")
	id, _ := strconv.Atoi(id_)

	category := entity.ProductCategory{}
	db := config.Connect()
	defer config.Close(db)
	claimsToken := util.GetUser(c)
	db.Model(&category) .Where("company_id = ?", claimsToken["companyId"]).Find(&category, id)
	if category.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrNotFound, "Entity not found"))
	}
	response := dto.ProductCategoryRs{
		ID: category.ID,
		Name: category.Name,
		Description: category.Description,
		CreatedAt: category.CreatedAt,
		CreatedBy: category.CreatedBy,
		UpdatedAt: category.UpdatedAt,
		UpdatedBy: category.UpdatedBy,
	}
	return c.Status(fiber.StatusOK).JSON(util.GenerateResponse(response))

}
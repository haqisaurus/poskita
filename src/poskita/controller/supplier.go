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

func AddSupplier(c *fiber.Ctx) error {
	db := config.Connect()
	defer config.Close(db)
	var req dto.ProductSupplierRq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrInvalidInput, err.Error()))
	}

	claimsToken := util.GetUser(c)

	supplier := entity.ProductSupplier{
		CreatedBy:   claimsToken["username"].(string),
		CreatedAt:   time.Now(),
		UpdatedBy:   claimsToken["username"].(string),
		UpdatedAt:   time.Now(),
		CompanyID:   uint64(claimsToken["companyId"].(float64)),
		Name:        req.Name,
		Description: req.Description,
		Phone:       req.Phone,
		Address:     req.Address,
	}

	resultInsert := db.Create(&supplier)
	if resultInsert.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrInternal, resultInsert.Error.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(util.GenerateResponse(dto.InsertSuccess{
		ID: supplier.ID,
	}))
}

func EditSupplier(c *fiber.Ctx) error {
	db := config.Connect()
	defer config.Close(db)
	var req dto.ProductSupplierRq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrInvalidInput, err.Error()))
	}

	supplier := entity.ProductSupplier{}
	db.Model(&supplier).Find(&supplier, req.ID)
	if supplier.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrNotFound, "Entity not found"))
	}
	db.Model(&supplier).Updates(entity.AuthCompany{
		Name:        req.Name,
		Description: req.Description,
		Phone:       req.Phone,
		Address:     req.Address})
	return c.Status(fiber.StatusOK).JSON(util.GenerateResponse(dto.InsertSuccess{
		ID: supplier.ID,
	}))

}
func ListSupplier(c *fiber.Ctx) error {
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

	categories := []entity.ProductSupplier{}
	query := db.Model(&entity.ProductSupplier{})
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

	contents := []dto.ProductSupplierRs{}
	for _, supplier := range categories {
		contents = append(contents, dto.ProductSupplierRs{
			ID:          supplier.ID,
			Name:        supplier.Name,
			Description: supplier.Description,
			Phone:       supplier.Phone,
			Address:     supplier.Address,
			CreatedAt:   supplier.CreatedAt,
			CreatedBy:   supplier.CreatedBy,
			UpdatedAt:   supplier.UpdatedAt,
			UpdatedBy:   supplier.UpdatedBy,
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
func DetailSupplier(c *fiber.Ctx) error {
	id_ := c.Params("id", "0")
	id, _ := strconv.Atoi(id_)

	supplier := entity.ProductSupplier{}
	db := config.Connect()
	defer config.Close(db)
	claimsToken := util.GetUser(c)
	db.Model(&supplier).Where("company_id = ?", claimsToken["companyId"]).Find(&supplier, id)
	if supplier.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(util.GenerateError(util.ErrNotFound, "Entity not found"))
	}
	response := dto.ProductSupplierRs{
		ID:          supplier.ID,
		Name:        supplier.Name,
		Description: supplier.Description,
		Phone:       supplier.Phone,
		Address:     supplier.Address,
		CreatedAt:   supplier.CreatedAt,
		CreatedBy:   supplier.CreatedBy,
		UpdatedAt:   supplier.UpdatedAt,
		UpdatedBy:   supplier.UpdatedBy,
	}
	return c.Status(fiber.StatusOK).JSON(util.GenerateResponse(response))

}

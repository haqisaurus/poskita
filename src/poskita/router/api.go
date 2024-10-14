package router

import (

	"github.com/gofiber/fiber/v2"
	"github.com/haqisaurus/poskita/controller"

)

func SetupRouter(app *fiber.App) {
	app.Post("/api/v1/login", controller.PostLogin)

	app.Get("/api/v1/my-companies", FirstAuthMiddleware, controller.GetCompanyList)
	app.Post("/api/v1/login-company", FirstAuthMiddleware, controller.PostLoginCompany)
	app.Get("/api/v1/me", CompanyAuthMiddleware, controller.GetMe)
	app.Post("/api/v1/refresh-token", controller.RefreshToken)

	app.Post("/api/v1/company-add", CompanyAuthMiddleware, controller.AddCompany)
	app.Post("/api/v1/company-edit", CompanyAuthMiddleware, controller.EditCompany)
	app.Get("/api/v1/company-list", CompanyAuthMiddleware, controller.ListCompany)
	app.Get("/api/v1/company-detail/:id", CompanyAuthMiddleware, controller.DetailCompany)
	 
}
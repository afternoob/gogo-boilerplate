package app

import (
	"github.com/afternoob/gogo-boilerplate/service/company"
	"github.com/afternoob/gogo-boilerplate/service/staff"
	"github.com/gin-gonic/gin"
)

type App struct {
	staffService   staff.Service
	companyService company.Service
}

func New(staffService staff.Service, companyService company.Service) *App {
	return &App{
		staffService:   staffService,
		companyService: companyService,
	}
}

func (app *App) RegisterRoute(router *gin.Engine) *App {
	router.POST("/staff", app.CreateStaff)
	router.PUT("/staff", app.UpdateStaff)
	router.GET("/staffsByCompany", app.GetStaffsByCompany)
	router.POST("/company", app.CreateCompany)

	return app
}

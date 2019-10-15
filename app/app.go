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

func (app *App) RegisterRoute(ge *gin.Engine) *App {
	ge.POST("/staff", app.CreateStaff)
	ge.PUT("/staff", app.UpdateStaff)
	ge.GET("/staffsByCompany", app.GetStaffsByCompany)
	ge.POST("/company", app.CreateCompany)

	return app
}

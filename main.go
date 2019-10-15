package main

import (
	"github.com/afternoob/gogo-boilerplate/app"
	companyRepo "github.com/afternoob/gogo-boilerplate/repository/company/store"
	staffRepo "github.com/afternoob/gogo-boilerplate/repository/staff/store"
	companyService "github.com/afternoob/gogo-boilerplate/service/company"
	staffService "github.com/afternoob/gogo-boilerplate/service/staff"
	"github.com/devit-tel/goxid"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	_ = newApp().RegisterRoute(g)
	g.Run()
}

func newApp() *app.App {
	xid := goxid.New()

	companyStore := companyRepo.New()
	company := companyService.New(xid, companyStore)

	staffStore := staffRepo.New()
	staff := staffService.New(xid, staffStore, companyStore)

	return app.New(staff, company)
}

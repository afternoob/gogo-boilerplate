package app

import (
	"net/http"

	serviceCompany "github.com/afternoob/gogo-boilerplate/service/company"
	"github.com/devit-tel/goerror/ginresp"
	"github.com/gin-gonic/gin"
)

type Company struct {
	Name string `json:"name"`
}

type CreateCompanyInput struct {
	Name string `json:"name" binding:"required"`
}

type CreateCompanyOutput struct {
	Company *Company `json:"company"`
}

func (app *App) CreateCompany(c *gin.Context) {
	input := &CreateCompanyInput{}
	if err := c.ShouldBind(input); err != nil {
		ginresp.RespValidateError(c, err)
		return
	}

	company, err := app.companyService.CreateCompany(&serviceCompany.CreateCompanyInput{Name: input.Name})
	if err != nil {
		ginresp.RespWithError(c, err)
		return
	}

	// warning: maybe create new function for map response
	c.JSON(http.StatusOK, &CreateCompanyOutput{
		Company: &Company{
			Name: company.Name,
		},
	})
}

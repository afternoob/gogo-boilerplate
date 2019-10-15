package app

import (
	"net/http"

	domainStaff "github.com/afternoob/gogo-boilerplate/domain/staff"
	serviceStaff "github.com/afternoob/gogo-boilerplate/service/staff"

	"github.com/devit-tel/goerror/ginresp"
	"github.com/gin-gonic/gin"
)

type Staff struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CompanyId string `json:"companyId"`
	Tel       string `json:"tel"`
}

type CreateStaffInput struct {
	Name      string `json:"name" binding:"required"`
	CompanyId string `json:"companyId" binding:"required"`
	Tel       string `json:"tel"`
}

type CreateStaffOutput struct {
	Staff *Staff `json:"staff"`
}

func (app *App) CreateStaff(c *gin.Context) {
	input := &CreateStaffInput{}
	if err := c.ShouldBind(input); err != nil {
		ginresp.RespValidateError(c, err)
		return
	}

	staff, err := app.staffService.CreateStaff(&serviceStaff.CreateStaffInput{
		Name:      input.Name,
		CompanyId: input.CompanyId,
		Tel:       input.Tel,
	})
	if err != nil {
		ginresp.RespWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, &CreateStaffOutput{
		Staff: toStaffOutput(staff),
	})
}

func toStaffOutput(staff *domainStaff.Staff) *Staff {
	return &Staff{
		Id:        staff.Id,
		Name:      staff.Name,
		CompanyId: staff.CompanyId,
		Tel:       staff.Tel,
	}
}

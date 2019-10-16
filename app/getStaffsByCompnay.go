package app

import (
	"fmt"
	"net/http"

	"github.com/afternoob/gogo-boilerplate/app/inout/staff"
	serviceStaff "github.com/afternoob/gogo-boilerplate/service/staff"
	"github.com/devit-tel/goerror/ginresp"
	"github.com/gin-gonic/gin"
)

func (app *App) GetStaffsByCompany(c *gin.Context) {
	input := &staff.GetStaffsByCompanyInput{}
	if err := c.ShouldBind(input); err != nil {
		fmt.Println(err)
		ginresp.RespValidateError(c, err)
		return
	}

	staffs, err := app.staffService.GetStaffsByCompany(c.Request.Context(),
		&serviceStaff.GetStaffsByCompanyInput{
			CompanyId: input.CompanyId,
			Offset:    input.Offset,
			Limit:     input.Limit,
		})
	if err != nil {
		ginresp.RespWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, &staff.GetStaffsByCompanyOutput{
		Staffs: staff.ToStaffsOutput(staffs),
	})
}

package app

import (
	"net/http"

	staff2 "github.com/afternoob/gogo-boilerplate/app/inout/staff"
	serviceStaff "github.com/afternoob/gogo-boilerplate/service/staff"

	"github.com/devit-tel/goerror/ginresp"
	"github.com/gin-gonic/gin"
)

func (app *App) CreateStaff(c *gin.Context) {
	input := &staff2.CreateStaffInput{}
	if err := c.ShouldBind(input); err != nil {
		ginresp.RespValidateError(c, err)
		return
	}

	staff, err := app.staffService.CreateStaff(c.Request.Context(),
		&serviceStaff.CreateStaffInput{
			Name:      input.Name,
			CompanyId: input.CompanyId,
			Tel:       input.Tel,
		})
	if err != nil {
		ginresp.RespWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, &staff2.CreateStaffOutput{
		Staff: staff2.ToStaffOutput(staff),
	})
}

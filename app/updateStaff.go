package app

import (
	"net/http"

	staff2 "github.com/afternoob/gogo-boilerplate/app/inout/staff"
	serviceStaff "github.com/afternoob/gogo-boilerplate/service/staff"

	"github.com/devit-tel/goerror/ginresp"
	"github.com/gin-gonic/gin"
)

func (app *App) UpdateStaff(c *gin.Context) {
	input := &staff2.UpdateStaffInput{}
	if err := c.ShouldBind(input); err != nil {
		ginresp.RespValidateError(c, err)
		return
	}

	staff, err := app.staffService.UpdateStaff(c.Request.Context(),
		&serviceStaff.UpdateStaffInput{
			StaffId: input.Id,
			Name:    input.Name,
			Tel:     input.Tel,
		})
	if err != nil {
		ginresp.RespWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, &staff2.UpdateStaffOutput{
		Staff: staff2.ToStaffOutput(staff),
	})
}

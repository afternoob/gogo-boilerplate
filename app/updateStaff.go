package app

import (
	"net/http"

	serviceStaff "github.com/afternoob/gogo-boilerplate/service/staff"

	"github.com/devit-tel/goerror/ginresp"
	"github.com/gin-gonic/gin"
)

type UpdateStaffInput struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Tel  string `json:"tel" binding:"required"`
}

type UpdateStaffOutput struct {
	Staff *Staff `json:"staff"`
}

func (app *App) UpdateStaff(c *gin.Context) {
	input := &UpdateStaffInput{}
	if err := c.ShouldBind(input); err != nil {
		ginresp.RespValidateError(c, err)
		return
	}

	staff, err := app.staffService.UpdateStaff(&serviceStaff.UpdateStaffInput{
		StaffId: input.Id,
		Name:    input.Name,
		Tel:     input.Tel,
	})
	if err != nil {
		ginresp.RespWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, &UpdateStaffOutput{
		Staff: toStaffOutput(staff),
	})
}

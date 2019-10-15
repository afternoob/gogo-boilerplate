package app

import (
	"fmt"
	"net/http"

	domainStaff "github.com/afternoob/gogo-boilerplate/domain/staff"
	serviceStaff "github.com/afternoob/gogo-boilerplate/service/staff"

	"github.com/devit-tel/goerror/ginresp"
	"github.com/gin-gonic/gin"
)

type GetStaffsByCompanyInput struct {
	CompanyId string `json:"companyId" binding:"required"`
	Limit     int    `json:"limit,default=20"`
	Offset    int    `json:"offset"`
}

type GetStaffsByCompanyOutput struct {
	Staffs []*Staff `json:"staffs"`
}

func (app *App) GetStaffsByCompany(c *gin.Context) {
	input := &GetStaffsByCompanyInput{}
	if err := c.ShouldBind(input); err != nil {
		fmt.Println(err)
		ginresp.RespValidateError(c, err)
		return
	}

	staffs, err := app.staffService.GetStaffsByCompany(&serviceStaff.GetStaffsByCompanyInput{
		CompanyId: input.CompanyId,
		Offset:    input.Offset,
		Limit:     input.Limit,
	})
	if err != nil {
		ginresp.RespWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, &GetStaffsByCompanyOutput{
		Staffs: toStaffsOutput(staffs),
	})
}

func toStaffsOutput(staffs []*domainStaff.Staff) []*Staff {
	outputStaffs := make([]*Staff, len(staffs))

	for index, staff := range staffs {
		outputStaffs[index] = toStaffOutput(staff)
	}

	return outputStaffs
}

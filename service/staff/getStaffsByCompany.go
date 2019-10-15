package staff

import (
	domainStaff "github.com/afternoob/gogo-boilerplate/domain/staff"
	"github.com/devit-tel/goerror"
)

type GetStaffsByCompanyInput struct {
	CompanyId string
	Offset    int
	Limit     int
}

func (service *StaffService) GetStaffsByCompany(input *GetStaffsByCompanyInput) ([]*domainStaff.Staff, goerror.Error) {
	staffs, err := service.staffRepository.GetStaffsByCompany(input.CompanyId, input.Offset, input.Limit)
	if err != nil {
		return nil, err
	}

	return staffs, nil
}

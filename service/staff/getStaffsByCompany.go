package staff

import (
	"context"

	domainStaff "github.com/afternoob/gogo-boilerplate/domain/staff"
	"github.com/devit-tel/goerror"
)

type GetStaffsByCompanyInput struct {
	CompanyId string
	Offset    int64
	Limit     int64
}

func (service *StaffService) GetStaffsByCompany(ctx context.Context, input *GetStaffsByCompanyInput) ([]*domainStaff.Staff, goerror.Error) {
	staffs, err := service.staffRepository.GetStaffsByCompany(ctx, input.CompanyId, input.Offset, input.Limit)
	if err != nil {
		return nil, err
	}

	return staffs, nil
}

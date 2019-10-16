package staff

import (
	"context"

	domainStaff "github.com/afternoob/gogo-boilerplate/domain/staff"
	"github.com/devit-tel/goerror"
)

type UpdateStaffInput struct {
	StaffId string
	Name    string
	Tel     string
}

func (service *StaffService) UpdateStaff(ctx context.Context, input *UpdateStaffInput) (*domainStaff.Staff, goerror.Error) {
	staff, err := service.staffRepository.Get(ctx, input.StaffId)
	if err != nil {
		return nil, err
	}

	staff.Update(input.Name, input.Tel)

	if err := service.staffRepository.Save(ctx, staff); err != nil {
		return nil, err
	}

	return staff, nil
}

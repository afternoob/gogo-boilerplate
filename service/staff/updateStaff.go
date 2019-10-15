package staff

import (
	domainStaff "github.com/afternoob/gogo-boilerplate/domain/staff"
	"github.com/devit-tel/goerror"
)

type UpdateStaffInput struct {
	StaffId string
	Name    string
	Tel     string
}

func (service *StaffService) UpdateStaff(input *UpdateStaffInput) (*domainStaff.Staff, goerror.Error) {
	staff, err := service.staffRepository.Get(input.StaffId)
	if err != nil {
		return nil, err
	}

	staff.Update(input.Name, input.Tel)

	if err := service.staffRepository.Save(staff); err != nil {
		return nil, err
	}

	return staff, nil
}

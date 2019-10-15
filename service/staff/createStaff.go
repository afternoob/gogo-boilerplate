package staff

import (
	domainStaff "github.com/afternoob/gogo-boilerplate/domain/staff"
	"github.com/devit-tel/goerror"
)

type CreateStaffInput struct {
	Name      string
	CompanyId string
	Tel       string
}

func (service *StaffService) CreateStaff(input *CreateStaffInput) (*domainStaff.Staff, goerror.Error) {
	_, err := service.companyRepository.Get(input.CompanyId)
	if err != nil {
		return nil, err
	}

	newStaff := domainStaff.Create(service.xid.Gen(), input.CompanyId, input.Name, input.Tel)

	if err := service.staffRepository.Save(newStaff); err != nil {
		return nil, err
	}

	return newStaff, nil
}

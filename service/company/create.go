package company

import (
	domainCompany "github.com/afternoob/gogo-boilerplate/domain/company"
	"github.com/devit-tel/goerror"
)

type CreateCompanyInput struct {
	Name string
}

func (service *CompanyService) CreateCompany(input *CreateCompanyInput) (*domainCompany.Company, goerror.Error) {
	newCompany := domainCompany.Create(service.xid.Gen(), input.Name)

	if err := service.companyRepository.Save(newCompany); err != nil {
		return nil, err
	}

	return newCompany, nil
}

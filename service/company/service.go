package company

import (
	domainCompany "github.com/afternoob/gogo-boilerplate/domain/company"
	"github.com/afternoob/gogo-boilerplate/repository/company"
	"github.com/devit-tel/goerror"
	"github.com/devit-tel/goxid"
)

//go:generate mockery -name=Service
type Service interface {
	CreateCompany(input *CreateCompanyInput) (*domainCompany.Company, goerror.Error)
}

type CompanyService struct {
	companyRepository company.Repository
	xid               *goxid.ID
}

func New(xid *goxid.ID, c company.Repository) *CompanyService {
	return &CompanyService{
		companyRepository: c,
		xid:               xid,
	}
}

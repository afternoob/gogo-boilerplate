package company

import (
	"github.com/afternoob/gogo-boilerplate/domain/company"
	"github.com/devit-tel/goerror"
)

var (
	ErrCompanyNotFound   = goerror.DefineNotFound("CompanyNotFound", "company not found")
	ErrUnableGetCompany  = goerror.DefineNotFound("UnableGetCompany", "unable to get company")
	ErrUnableSaveCompany = goerror.DefineNotFound("UnableSaveCompany", "unable to save company")
)

//go:generate mockery -name=Repository
type Repository interface {
	Save(company *company.Company) goerror.Error
	Get(companyId string) (*company.Company, goerror.Error)
}

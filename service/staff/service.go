package staff

import (
	"context"

	domainStaff "github.com/afternoob/gogo-boilerplate/domain/staff"
	"github.com/afternoob/gogo-boilerplate/repository/company"
	"github.com/afternoob/gogo-boilerplate/repository/staff"
	"github.com/devit-tel/goerror"
	"github.com/devit-tel/goxid"
)

//go:generate mockery -name=Service
type Service interface {
	GetStaffsByCompany(ctx context.Context, input *GetStaffsByCompanyInput) ([]*domainStaff.Staff, goerror.Error)
	CreateStaff(ctx context.Context, input *CreateStaffInput) (*domainStaff.Staff, goerror.Error)
	UpdateStaff(ctx context.Context, input *UpdateStaffInput) (*domainStaff.Staff, goerror.Error)
}

type StaffService struct {
	xid               *goxid.ID
	staffRepository   staff.Repository
	companyRepository company.Repository
}

func New(xid *goxid.ID, r staff.Repository, c company.Repository) *StaffService {
	return &StaffService{
		xid:               xid,
		staffRepository:   r,
		companyRepository: c,
	}
}

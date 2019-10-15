package staff

import (
	"github.com/afternoob/gogo-boilerplate/domain/staff"
	"github.com/devit-tel/goerror"
)

var (
	ErrStaffNotFound   = goerror.DefineNotFound("StaffNotFound", "staff not found")
	ErrUnableGetStaff  = goerror.DefineNotFound("UnableGetStaff", "unable to get staff")
	ErrUnableGetStaffs = goerror.DefineNotFound("UnableGetStaffs", "unable to get staffs by company")
	ErrUnableSaveStaff = goerror.DefineNotFound("UnableSaveStaff", "unable to save staff")
)

//go:generate mockery -name=Repository
type Repository interface {
	Get(staffId string) (*staff.Staff, goerror.Error)
	GetStaffsByCompany(companyId string, offset, limit int) ([]*staff.Staff, goerror.Error)
	Save(staff *staff.Staff) goerror.Error
}

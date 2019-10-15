package store

import (
	domain "github.com/afternoob/gogo-boilerplate/domain/staff"
	repoStaff "github.com/afternoob/gogo-boilerplate/repository/staff"
	"github.com/devit-tel/goerror"
)

func New() *Store {
	return &Store{Data: map[string]*domain.Staff{}}
}

type Store struct {
	Data map[string]*domain.Staff
}

func (s *Store) Get(staffId string) (*domain.Staff, goerror.Error) {
	staff, ok := s.Data[staffId]
	if !ok {
		return nil, repoStaff.ErrStaffNotFound.WithInput(staffId)
	}

	return staff, nil
}

func (s *Store) Save(staff *domain.Staff) goerror.Error {
	s.Data[staff.Id] = staff

	return nil
}

func (s *Store) GetStaffsByCompany(companyId string, offset, limit int) ([]*domain.Staff, goerror.Error) {
	staffs := make([]*domain.Staff, 0)
	for _, staff := range s.Data {
		if staff.CompanyId == companyId {
			staffs = append(staffs, staff)
		}
	}

	if offset > len(staffs) {
		return []*domain.Staff{}, nil
	}

	if limit > len(staffs) {
		return staffs[offset:], nil
	}

	return staffs[offset:limit], nil
}

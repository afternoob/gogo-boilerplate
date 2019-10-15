package store

import (
	domain "github.com/afternoob/gogo-boilerplate/domain/company"
	repoCompany "github.com/afternoob/gogo-boilerplate/repository/company"
	"github.com/devit-tel/goerror"
)

func New() *Store {
	return &Store{Data: map[string]*domain.Company{}}
}

type Store struct {
	Data map[string]*domain.Company
}

func (s *Store) Get(companyId string) (*domain.Company, goerror.Error) {
	company, ok := s.Data[companyId]
	if !ok {
		return nil, repoCompany.ErrCompanyNotFound.WithInput(companyId)
	}

	return company, nil
}

func (s *Store) Save(company *domain.Company) goerror.Error {
	s.Data[company.Id] = company

	return nil
}

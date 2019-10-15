package staff

import (
	"github.com/devit-tel/gotime"
)

type Staff struct {
	Id        string
	CompanyId string
	Name      string
	Tel       string
	CreatedAt int64
	UpdatedAt int64
}

func Create(id, companyId, name, tel string) *Staff {
	return &Staff{
		Id:        id,
		CompanyId: companyId,
		Name:      name,
		Tel:       tel,
		CreatedAt: gotime.NowUnix(),
		UpdatedAt: gotime.NowUnix(),
	}
}

func (s *Staff) Update(name, tel string) {
	s.Name = name
	s.Tel = tel
	s.UpdatedAt = gotime.NowUnix()
}

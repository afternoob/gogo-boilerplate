package company

type Company struct {
	Id   string
	Name string
}

func Create(id, name string) *Company {
	return &Company{
		Id:   id,
		Name: name,
	}
}

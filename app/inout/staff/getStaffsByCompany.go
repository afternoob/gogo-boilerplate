package staff

type GetStaffsByCompanyInput struct {
	CompanyId string `json:"companyId" binding:"required"`
	Limit     int64  `json:"limit,default=20"`
	Offset    int64  `json:"offset"`
}

type GetStaffsByCompanyOutput struct {
	Staffs []*Staff `json:"staffs"`
}

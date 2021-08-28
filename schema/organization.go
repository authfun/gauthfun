package schema

func (Organization) TableName() string {
	return "organization"
}

type Organization struct {
	Id       string `gorm:"id" json:"id"`
	ParentId *string `gorm:"parent_id" json:"parentId"`
	TenantId string  `gorm:"tenant_id" json:"tenantId"`
	Name     string `gorm:"name" json:"name"`
}

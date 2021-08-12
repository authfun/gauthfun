package schema

func (Tenant) TableName() string {
	return "tenant"
}

type Tenant struct {
	Id   string `gorm:"id" json:"id"`
	Name string `gorm:"name" json:"name"`
}

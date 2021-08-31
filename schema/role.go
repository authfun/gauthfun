package schema

func (Role) TableName() string {
	return "role"
}

type Role struct {
	Id   string `gorm:"id" json:"id"`
	Name string `gorm:"name" json:"name"`
}

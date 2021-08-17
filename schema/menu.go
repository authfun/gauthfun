package schema

func (Menu) TableName() string {
	return "menu"
}

type Menu struct {
	Id   string `gorm:"id" json:"id"`
	Name string `gorm:"name" json:"name"`
}

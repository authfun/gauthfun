package schema

func (User) TableName() string {
	return "user"
}

type User struct {
	Id   string `gorm:"id" json:"id"`
	Name string `gorm:"name" json:"name"`
}

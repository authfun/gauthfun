package schema

func (Feature) TableName() string {
	return "feature"
}

type Feature struct {
	Id   string `gorm:"id" json:"id"`
	Name string `gorm:"name" json:"name"`
}

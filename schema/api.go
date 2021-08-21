package schema

func (Api) TableName() string {
	return "api"
}

type Api struct {
	Id     string `gorm:"id" json:"id"`
	Name   string `gorm:"name" json:"name"`
	Group  string `gorm:"group" json:"group"`
	Method string `gorm:"method" json:"method"`
	Route  string `gorm:"route" json:"route"`
}

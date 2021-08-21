package model

type ApiForm struct {
	Name   string `json:"name"`
	Group  string `json:"group"`
	Method string `json:"method"`
	Route  string `json:"route"`
}

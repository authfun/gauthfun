package model

type OrganizationNode struct {
	Id       string               `json:"id"`
	ParentId *string              `json:"parentId"`
	Name     string              `json:"name"`
	Children []*OrganizationNode `json:"children"`
}

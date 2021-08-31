package model

import (
	"github.com/authfun/gauthfun/schema"
)

type OrganizationDetail struct {
	Id       string  `json:"id"`
	ParentId *string `json:"parentId"`
	TenantId *string `json:"tenantId"`
	Name     string  `json:"name"`

	Roles []schema.Role `json:"roles"`
	Users []schema.User `json:"users"`
}

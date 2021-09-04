package service

import (
	"fmt"

	"github.com/authfun/gauthfun/casbin"
	"github.com/authfun/gauthfun/consts"
)

func GetUsersForOrganization(organizationId string, tenantId string) ([]string, error) {
	name := consts.OrganizationPrefix + fmt.Sprintf("%v", organizationId)
	domain := consts.TenantPrefix + fmt.Sprintf("%v", tenantId)
	return casbin.Enforcer.GetUsersForRole(name, domain)
}

func GetRolesForOrganization(organizationId string, tenantId string) ([]string, error) {
	name := consts.OrganizationPrefix + fmt.Sprintf("%v", organizationId)
	domain := consts.TenantPrefix + fmt.Sprintf("%v", tenantId)
	return casbin.Enforcer.GetRolesForUser(name, domain)
}

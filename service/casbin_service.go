package service

import (
	"fmt"
	"strings"

	linq "github.com/ahmetb/go-linq/v3"
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

func GetObjectsForFeature(featureId string, implicit bool) ([][]string, error) {
	sub := consts.FeaturePrefix + fmt.Sprintf("%v", featureId)

	if implicit {
		return casbin.Enforcer.GetImplicitPermissionsForUser(sub, consts.Domain_Pattern_All)
	} else {
		permissions := casbin.Enforcer.GetPermissionsForUser(sub, consts.Domain_Pattern_All)
		return permissions, nil
	}
}

func GetFeaturesForFeature(featureId string, implicit bool) ([]string, error) {
	sub := consts.FeaturePrefix + fmt.Sprintf("%v", featureId)

	if implicit {
		return casbin.Enforcer.GetImplicitRolesForUser(sub, consts.Domain_Pattern_All)
	} else {
		return casbin.Enforcer.GetRolesForUser(sub, consts.Domain_Pattern_All)
	}
}

func GetIds(raws []string, prefix string) []string {
	var ids []string
	linq.From(raws).SelectT(func(raw string) string {
		id := strings.ReplaceAll(raw, prefix, "")
		return id
	}).ToSlice(&ids)

	return ids
}

func FilterIds(raws []string, prefix string) []string {
	var ids []string
	linq.From(raws).WhereT(func(raw string) bool {
		return strings.Contains(raw, prefix)
	}).SelectT(func(raw string) string {
		id := strings.ReplaceAll(raw, prefix, "")
		return id
	}).ToSlice(&ids)

	return ids
}

func FilterObjects(objects [][]string, prefix string) []string {
	var ids []string
	linq.From(objects).WhereT(func(obj []string) bool {
		return strings.Contains(obj[2], prefix)
	}).SelectT(func(obj []string) string {
		id := strings.ReplaceAll(obj[2], prefix, "")
		return id
	}).ToSlice(&ids)

	return ids
}

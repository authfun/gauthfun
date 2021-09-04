package model

type FeatureForm struct {
	Name       string   `json:"name"`
	MenuIds    []string `json:"menuIds"`
	ApiIds     []string `json:"apiIds"`
	FeatureIds []string `json:"featureIds"`
}

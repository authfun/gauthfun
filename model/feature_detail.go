package model

import "github.com/authfun/gauthfun/schema"

type FeatureDetail struct {
	Name     string           `json:"name"`
	Menus    []schema.Menu    `json:"menus"`
	Apis     []schema.Api     `json:"apis"`
	Features []schema.Feature `json:"features"`
}

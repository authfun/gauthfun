// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/apis": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Api"
                ],
                "summary": "Add api",
                "parameters": [
                    {
                        "description": "Api info",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ApiForm"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/schema.Api"
                        }
                    }
                }
            }
        },
        "/api/apis/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Api"
                ],
                "summary": "Get api list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schema.Api"
                            }
                        }
                    }
                }
            }
        },
        "/api/apis/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Api"
                ],
                "summary": "Get api",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Api"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Api"
                ],
                "summary": "Update api",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Api info",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ApiForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Api"
                ],
                "summary": "Delete api",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        },
        "/api/features": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Feature"
                ],
                "summary": "Add feature",
                "parameters": [
                    {
                        "description": "Feature info",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.FeatureForm"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/schema.Feature"
                        }
                    }
                }
            }
        },
        "/api/features/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Feature"
                ],
                "summary": "Get feature list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schema.Feature"
                            }
                        }
                    }
                }
            }
        },
        "/api/features/options": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Feature"
                ],
                "summary": "Get feature option",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Option"
                            }
                        }
                    }
                }
            }
        },
        "/api/features/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Feature"
                ],
                "summary": "Get feature by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Feature ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Whether to get the implicit info",
                        "name": "implicit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.FeatureDetail"
                        }
                    }
                }
            }
        },
        "/api/menus": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "Add menu",
                "parameters": [
                    {
                        "description": "Menu info",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.MenuForm"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/schema.Menu"
                        }
                    }
                }
            }
        },
        "/api/menus/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "Get menu list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schema.Menu"
                            }
                        }
                    }
                }
            }
        },
        "/api/menus/options": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "Get menu option",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Option"
                            }
                        }
                    }
                }
            }
        },
        "/api/menus/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "Get menu",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Menu Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Menu"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "Update menu",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Menu Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Menu info",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.MenuForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "Delete menu",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Menu Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        },
        "/api/organizations/tree": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Organization"
                ],
                "summary": "Get organization tree",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tenant ID",
                        "name": "tenantId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.OrganizationNode"
                            }
                        }
                    }
                }
            }
        },
        "/api/organizations/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Organization"
                ],
                "summary": "Get organization",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Organization ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.OrganizationDetail"
                        }
                    }
                }
            }
        },
        "/api/tenants": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tenant"
                ],
                "summary": "Add tenant",
                "parameters": [
                    {
                        "description": "Tenant info",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TenantForm"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/schema.Tenant"
                        }
                    }
                }
            }
        },
        "/api/tenants/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tenant"
                ],
                "summary": "Get tenant list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schema.Tenant"
                            }
                        }
                    }
                }
            }
        },
        "/api/tenants/options": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tenant"
                ],
                "summary": "Get tenant option",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Option"
                            }
                        }
                    }
                }
            }
        },
        "/api/tenants/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tenant"
                ],
                "summary": "Get tenant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tenant Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Tenant"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tenant"
                ],
                "summary": "Update tenant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tenant Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Tenant info",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TenantForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tenant"
                ],
                "summary": "Delete tenant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tenant Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ApiForm": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "method": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "route": {
                    "type": "string"
                }
            }
        },
        "model.FeatureDetail": {
            "type": "object",
            "properties": {
                "apis": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Api"
                    }
                },
                "features": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Feature"
                    }
                },
                "menus": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Menu"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.FeatureForm": {
            "type": "object",
            "properties": {
                "apiIds": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "featureIds": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "menuIds": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.MenuForm": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Option": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.OrganizationDetail": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parentId": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Role"
                    }
                },
                "tenantId": {
                    "type": "string"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.User"
                    }
                }
            }
        },
        "model.OrganizationNode": {
            "type": "object",
            "properties": {
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.OrganizationNode"
                    }
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parentId": {
                    "type": "string"
                }
            }
        },
        "model.TenantForm": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "schema.Api": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "method": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "route": {
                    "type": "string"
                }
            }
        },
        "schema.Feature": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "schema.Menu": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "schema.Role": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "schema.Tenant": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "schema.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}

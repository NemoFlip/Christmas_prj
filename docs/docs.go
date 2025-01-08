// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/api/users/data": {
            "post": {
                "description": "post request from user for present recommendation",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "API"
                ],
                "summary": "User request",
                "parameters": [
                    {
                        "description": "user's preferences",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user's preferences has been accepted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "invalid user's data",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.User": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "interests": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "job": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Christmas prj",
	Description:      "Simple project with recommendation system",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

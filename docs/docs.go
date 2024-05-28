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
        "/fill/{height}/{width}/{image}": {
            "get": {
                "description": "Resize image by URL address",
                "consumes": [
                    "multipart/form"
                ],
                "tags": [
                    "fill"
                ],
                "summary": "Fill",
                "operationId": "fill",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Height",
                        "name": "height",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Width",
                        "name": "width",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Image URL address",
                        "name": "image",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ok",
                        "schema": {
                            "$ref": "#/definitions/http.ResponseForm"
                        }
                    },
                    "400": {
                        "description": "Client error",
                        "schema": {
                            "$ref": "#/definitions/http.BadRequestForm"
                        }
                    },
                    "502": {
                        "description": "Gateway error",
                        "schema": {
                            "$ref": "#/definitions/http.BadRequestForm"
                        }
                    },
                    "503": {
                        "description": "Server does not available\".",
                        "schema": {
                            "$ref": "#/definitions/http.ServerErrorForm"
                        }
                    }
                }
            }
        },
        "/hello/": {
            "get": {
                "description": "Check service is available",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hello"
                ],
                "summary": "Hello",
                "operationId": "hello",
                "responses": {
                    "200": {
                        "description": "Ok",
                        "schema": {
                            "$ref": "#/definitions/http.ResponseForm"
                        }
                    },
                    "503": {
                        "description": "Server does not available\".",
                        "schema": {
                            "$ref": "#/definitions/http.ServerErrorForm"
                        }
                    }
                }
            }
        },
        "/tests/tests.jpg": {
            "get": {
                "description": "Resize image by URL address",
                "produces": [
                    "multipart/form"
                ],
                "tags": [
                    "tests"
                ],
                "summary": "TestDownload",
                "operationId": "tests-download",
                "responses": {
                    "200": {
                        "description": "Ok",
                        "schema": {
                            "$ref": "#/definitions/http.ResponseForm"
                        }
                    },
                    "400": {
                        "description": "Client error",
                        "schema": {
                            "$ref": "#/definitions/http.BadRequestForm"
                        }
                    },
                    "502": {
                        "description": "Gateway error",
                        "schema": {
                            "$ref": "#/definitions/http.BadRequestForm"
                        }
                    },
                    "503": {
                        "description": "Server does not available\".",
                        "schema": {
                            "$ref": "#/definitions/http.ServerErrorForm"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.BadRequestForm": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Bad Request message"
                },
                "status": {
                    "type": "integer",
                    "example": 400
                }
            }
        },
        "http.ResponseForm": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Done"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "http.ServerErrorForm": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Server Error message"
                },
                "status": {
                    "type": "integer",
                    "example": 503
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

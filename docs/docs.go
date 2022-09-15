// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/api/books": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book"
                ],
                "summary": "Get all",
                "responses": {
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controllers.ResponseHTTP"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create new book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book"
                ],
                "summary": "Book's creation",
                "parameters": [
                    {
                        "description": "1, Some-Book, 200",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/book.Book"
                        }
                    }
                ],
                "responses": {
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controllers.ResponseHTTP"
                        }
                    }
                }
            }
        },
        "/api/books/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get book by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book"
                ],
                "summary": "Get book by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controllers.ResponseHTTP"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "sign-in process",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "sign-in",
                "parameters": [
                    {
                        "description": "AAA 123",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                ],
                "responses": {
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controllers.ResponseHTTP"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "Sign-Up process",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "sign-up",
                "parameters": [
                    {
                        "description": "AAA 123",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                ],
                "responses": {
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controllers.ResponseHTTP"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "book.Book": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                }
            }
        },
        "controllers.ResponseHTTP": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "user.User": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "Jwt"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Book App",
	Description:      "This is an API for Book Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

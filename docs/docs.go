// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Flastor"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/songs": {
            "get": {
                "description": "get musics",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "music"
                ],
                "summary": "List musics",
                "parameters": [
                    {
                        "type": "string",
                        "description": "song filter",
                        "name": "song",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "group filter",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "release_date filter",
                        "name": "release_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "link filter",
                        "name": "link",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "text filter",
                        "name": "text",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "set output limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "set offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Music"
                            }
                        }
                    },
                    "418": {
                        "description": "I'm a teapot"
                    }
                }
            },
            "post": {
                "description": "create music",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "music"
                ],
                "summary": "Create music",
                "parameters": [
                    {
                        "description": "Create music",
                        "name": "music",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateMusicDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Music"
                        }
                    },
                    "418": {
                        "description": "I'm a teapot"
                    }
                }
            }
        },
        "/songs/{id}": {
            "get": {
                "description": "get music by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "music"
                ],
                "summary": "Show a music",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Music ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Music"
                        }
                    },
                    "418": {
                        "description": "I'm a teapot"
                    }
                }
            },
            "put": {
                "description": "update music",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "music"
                ],
                "summary": "Update music",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Music ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update music",
                        "name": "music",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateMusicDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "418": {
                        "description": "I'm a teapot"
                    }
                }
            },
            "delete": {
                "description": "delete music by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "music"
                ],
                "summary": "delete a music",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Music ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "418": {
                        "description": "I'm a teapot"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CreateMusicDTO": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                }
            }
        },
        "model.Music": {
            "type": "object",
            "properties": {
                "artist": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "model.UpdateMusicDTO": {
            "type": "object",
            "properties": {
                "artist": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Songius",
	Description:      "This is a simple music library.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

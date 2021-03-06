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
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Root path to give hint about available methods",
                "summary": "Show possible API routes"
            }
        },
        "/blocks/{blockParameter}": {
            "get": {
                "description": "get all header information of the block",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get block's header information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "blockNumber/blockHash/latest",
                        "name": "blockParameter",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Header"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/groups": {
            "get": {
                "description": "get all groups in smart contract",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get groups",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/big.Int"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/groups/{groupId}": {
            "get": {
                "description": "get detail information of specific group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get group's information",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Group ID",
                        "name": "groupId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Group"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/indexes/{indexId}": {
            "get": {
                "description": "get detail information of specific index",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get index's information",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Index ID",
                        "name": "indexId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Index"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "big.Int": {
            "type": "object"
        },
        "controller.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "model.Group": {
            "type": "object",
            "properties": {
                "indexes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/big.Int"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Index": {
            "type": "object",
            "properties": {
                "ethpriceinwei": {
                    "$ref": "#/definitions/big.Int"
                },
                "name": {
                    "type": "string"
                },
                "percentagechange": {
                    "$ref": "#/definitions/big.Int"
                },
                "usdcapitalization": {
                    "$ref": "#/definitions/big.Int"
                },
                "usdpriceincents": {
                    "$ref": "#/definitions/big.Int"
                }
            }
        },
        "types.Header": {
            "type": "object",
            "properties": {
                "baseFeePerGas": {
                    "description": "BaseFee was added by EIP-1559 and is ignored in legacy headers.",
                    "$ref": "#/definitions/big.Int"
                },
                "difficulty": {
                    "$ref": "#/definitions/big.Int"
                },
                "extraData": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "gasLimit": {
                    "type": "integer"
                },
                "gasUsed": {
                    "type": "integer"
                },
                "logsBloom": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "miner": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "mixHash": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "nonce": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "number": {
                    "$ref": "#/definitions/big.Int"
                },
                "parentHash": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "receiptsRoot": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "sha3Uncles": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "stateRoot": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "timestamp": {
                    "type": "integer"
                },
                "transactionsRoot": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "not used in current API's version": {
            "type": "basic"
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
	Version:     "1.0",
	Host:        "localhost:3000",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Ethereum smart contract API",
	Description: "This is a server to get information from smart contract 0x4f7f1380239450AAD5af611DB3c3c1bb51049c29 and general block information.",
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
	swag.Register("swagger", &s{})
}

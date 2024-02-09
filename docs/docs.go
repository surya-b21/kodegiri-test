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
        "/login": {
            "post": {
                "description": "Get token for auth",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login function",
                "parameters": [
                    {
                        "description": "Login Payload",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/loyalty-program": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all loyalty program list",
                "tags": [
                    "LoyaltyProgram"
                ],
                "summary": "LoyaltiProgramList function",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.LoyaltyProgram"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Store a loyalty program",
                "tags": [
                    "LoyaltyProgram"
                ],
                "summary": "LoyaltyPrograStore function",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/member-activity": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get token for auth",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Community"
                ],
                "summary": "Login function",
                "parameters": [
                    {
                        "description": "Member Activity Payload",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/community.MemberActivityPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/member-get-member": {
            "post": {
                "description": "Store a Member Get Member",
                "tags": [
                    "Community"
                ],
                "summary": "StoreMemberGetMember function",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/membership": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all membership list",
                "tags": [
                    "Membership"
                ],
                "summary": "Membership List function",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/membership/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get membership detail",
                "tags": [
                    "Membership"
                ],
                "summary": "Membership Detail function",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/redeem-point": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Store a reedemed point",
                "tags": [
                    "ReedemedPointStore"
                ],
                "summary": "ReedemedPointStore function",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/report": {
            "get": {
                "description": "Get list of report",
                "tags": [
                    "Report"
                ],
                "summary": "ListReport function",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/tier-management": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all tier list",
                "tags": [
                    "TierManagement"
                ],
                "summary": "Tier List function",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Tier"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Store a tier",
                "tags": [
                    "TierManagement"
                ],
                "summary": "Tier Store function",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Tier"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/tier-management/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get a tier detail",
                "tags": [
                    "TierManagement"
                ],
                "summary": "Tier Detail function",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Tier"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update a tier",
                "tags": [
                    "TierManagement"
                ],
                "summary": "Tier Update function",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Tier"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete a tier",
                "tags": [
                    "TierManagement"
                ],
                "summary": "Tier Delete function",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/transaction": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Store a transaction",
                "tags": [
                    "StoreTransaction"
                ],
                "summary": "StoreTransaction function",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.LoginPayload": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "remember_token": {
                    "type": "string"
                }
            }
        },
        "community.MemberActivityPayload": {
            "type": "object",
            "properties": {
                "activity_name": {
                    "type": "string"
                },
                "member_id": {
                    "type": "string"
                }
            }
        },
        "model.LoyaltyProgram": {
            "type": "object",
            "properties": {
                "benefit_community_fixed_point": {
                    "type": "integer"
                },
                "benefit_transaction_fixed_point": {
                    "type": "integer"
                },
                "benefit_transaction_precentage": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string",
                    "format": "date-time"
                },
                "id": {
                    "type": "string",
                    "format": "uuid"
                },
                "loyaltyProgramTier": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.LoyaltyProgramTier"
                    }
                },
                "loyalty_end": {
                    "type": "string",
                    "format": "date-time"
                },
                "loyalty_start": {
                    "type": "string",
                    "format": "date-time"
                },
                "name": {
                    "type": "string"
                },
                "policy_birthday_bonus": {
                    "type": "boolean"
                },
                "policy_is_community_activity": {
                    "type": "boolean"
                },
                "policy_is_community_member_get_member": {
                    "type": "boolean"
                },
                "policy_is_transaction_first_purchase": {
                    "type": "boolean"
                },
                "policy_transaction_amount": {
                    "type": "integer"
                },
                "policy_transaction_qty": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string",
                    "format": "date-time"
                }
            }
        },
        "model.LoyaltyProgramTier": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "format": "date-time"
                },
                "id": {
                    "type": "string",
                    "format": "uuid"
                },
                "loyaltyProgram": {
                    "$ref": "#/definitions/model.LoyaltyProgram"
                },
                "loyalty_program_id": {
                    "type": "string"
                },
                "tier": {
                    "$ref": "#/definitions/model.Tier"
                },
                "tier_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string",
                    "format": "date-time"
                }
            }
        },
        "model.Tier": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "format": "date-time"
                },
                "id": {
                    "type": "string",
                    "format": "uuid"
                },
                "maximal_point": {
                    "type": "integer"
                },
                "minimal_point": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string",
                    "format": "date-time"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Kodegiri App",
	Description:      "for test purpose",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

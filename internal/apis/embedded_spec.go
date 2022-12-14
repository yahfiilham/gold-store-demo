// Code generated by go-swagger; DO NOT EDIT.

package apis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a simple gold store API.",
    "title": "Gold Store Demo",
    "contact": {
      "name": "gold store maintainers",
      "url": "http://localhost:8081/docs",
      "email": "gold.store@email.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "https://www.apache.org/licenses/LICENSE-2.0"
    },
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "paths": {
    "/balance": {
      "get": {
        "description": "api for get balance",
        "tags": [
          "balance"
        ],
        "summary": "get balance",
        "operationId": "GetBalance",
        "parameters": [
          {
            "$ref": "#/parameters/accountNumber"
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "$ref": "#/definitions/account"
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          }
        }
      }
    },
    "/buyback": {
      "post": {
        "description": "api for buyback",
        "tags": [
          "buyback"
        ],
        "summary": "buyback",
        "operationId": "SaveBuyback",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/baseRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          }
        }
      }
    },
    "/health": {
      "get": {
        "tags": [
          "health check"
        ],
        "summary": "check server",
        "operationId": "getHealthCheck",
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          }
        }
      }
    },
    "/mutation": {
      "get": {
        "description": "api for get data transaction",
        "tags": [
          "transaction"
        ],
        "summary": "get data transaction",
        "operationId": "GetMutation",
        "parameters": [
          {
            "$ref": "#/parameters/accountNumber"
          },
          {
            "$ref": "#/parameters/startDate"
          },
          {
            "$ref": "#/parameters/endDate"
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/transaction"
                  }
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          }
        }
      }
    },
    "/price": {
      "get": {
        "description": "api for get current price",
        "tags": [
          "price"
        ],
        "summary": "get price",
        "operationId": "GetPrice",
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "$ref": "#/definitions/price"
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          }
        }
      },
      "post": {
        "description": "api for input price",
        "tags": [
          "price"
        ],
        "summary": "input price",
        "operationId": "SavePrice",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/priceData"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "success",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "$ref": "#/definitions/price"
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          }
        }
      }
    },
    "/topup": {
      "post": {
        "description": "api for topup",
        "tags": [
          "topup"
        ],
        "summary": "topup",
        "operationId": "SaveTopupGold",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/baseRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "account": {
      "allOf": [
        {
          "$ref": "#/definitions/modelID"
        },
        {
          "$ref": "#/definitions/modelLogTime"
        },
        {
          "$ref": "#/definitions/accountData"
        }
      ]
    },
    "accountData": {
      "type": "object",
      "properties": {
        "accountNumber": {
          "type": "string"
        },
        "balance": {
          "type": "number",
          "format": "double",
          "x-go-custom-tag": "gorm:\"type:decimal(10,3)\""
        }
      }
    },
    "baseRequest": {
      "type": "object",
      "properties": {
        "accountNumber": {
          "type": "string"
        },
        "gold": {
          "type": "number",
          "format": "double"
        },
        "price": {
          "type": "number",
          "format": "double"
        }
      }
    },
    "baseResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "modelID": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "x-go-custom-tag": "gorm:\"type:int primary key auto_increment\""
        }
      }
    },
    "modelLogTime": {
      "type": "object",
      "properties": {
        "createdAt": {
          "type": "integer",
          "format": "date-time",
          "x-go-custom-tag": "gorm:\"type:int;autoCreateTime\""
        },
        "updatedAt": {
          "type": "integer",
          "format": "date-time",
          "x-go-custom-tag": "gorm:\"type:int;autoUpdateTime\""
        }
      }
    },
    "price": {
      "allOf": [
        {
          "$ref": "#/definitions/modelID"
        },
        {
          "$ref": "#/definitions/modelLogTime"
        },
        {
          "$ref": "#/definitions/priceData"
        }
      ]
    },
    "priceData": {
      "type": "object",
      "properties": {
        "buybackPrice": {
          "type": "number",
          "format": "double"
        },
        "topupPrice": {
          "type": "number",
          "format": "double"
        }
      }
    },
    "transaction": {
      "allOf": [
        {
          "$ref": "#/definitions/modelID"
        },
        {
          "$ref": "#/definitions/modelLogTime"
        },
        {
          "$ref": "#/definitions/accountData"
        },
        {
          "$ref": "#/definitions/priceData"
        },
        {
          "$ref": "#/definitions/transactionData"
        }
      ]
    },
    "transactionData": {
      "type": "object",
      "properties": {
        "gold": {
          "type": "number",
          "format": "double",
          "x-go-custom-tag": "gorm:\"type:decimal(10,3)\""
        },
        "type": {
          "type": "string",
          "enum": [
            "topup",
            "buyback"
          ]
        }
      }
    }
  },
  "parameters": {
    "accountNumber": {
      "type": "string",
      "name": "account_no",
      "in": "query",
      "required": true
    },
    "endDate": {
      "type": "integer",
      "name": "end_date",
      "in": "query",
      "required": true
    },
    "startDate": {
      "type": "integer",
      "name": "start_date",
      "in": "query",
      "required": true
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a simple gold store API.",
    "title": "Gold Store Demo",
    "contact": {
      "name": "gold store maintainers",
      "url": "http://localhost:8081/docs",
      "email": "gold.store@email.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "https://www.apache.org/licenses/LICENSE-2.0"
    },
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "paths": {
    "/balance": {
      "get": {
        "description": "api for get balance",
        "tags": [
          "balance"
        ],
        "summary": "get balance",
        "operationId": "GetBalance",
        "parameters": [
          {
            "type": "string",
            "name": "account_no",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "$ref": "#/definitions/account"
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          }
        }
      }
    },
    "/buyback": {
      "post": {
        "description": "api for buyback",
        "tags": [
          "buyback"
        ],
        "summary": "buyback",
        "operationId": "SaveBuyback",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/baseRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          }
        }
      }
    },
    "/health": {
      "get": {
        "tags": [
          "health check"
        ],
        "summary": "check server",
        "operationId": "getHealthCheck",
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          }
        }
      }
    },
    "/mutation": {
      "get": {
        "description": "api for get data transaction",
        "tags": [
          "transaction"
        ],
        "summary": "get data transaction",
        "operationId": "GetMutation",
        "parameters": [
          {
            "type": "string",
            "name": "account_no",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "name": "start_date",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "name": "end_date",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/transaction"
                  }
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          }
        }
      }
    },
    "/price": {
      "get": {
        "description": "api for get current price",
        "tags": [
          "price"
        ],
        "summary": "get price",
        "operationId": "GetPrice",
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "$ref": "#/definitions/price"
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          }
        }
      },
      "post": {
        "description": "api for input price",
        "tags": [
          "price"
        ],
        "summary": "input price",
        "operationId": "SavePrice",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/priceData"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "success",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "$ref": "#/definitions/price"
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          }
        }
      }
    },
    "/topup": {
      "post": {
        "description": "api for topup",
        "tags": [
          "topup"
        ],
        "summary": "topup",
        "operationId": "SaveTopupGold",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/baseRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/baseResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "account": {
      "allOf": [
        {
          "$ref": "#/definitions/modelID"
        },
        {
          "$ref": "#/definitions/modelLogTime"
        },
        {
          "$ref": "#/definitions/accountData"
        }
      ]
    },
    "accountData": {
      "type": "object",
      "properties": {
        "accountNumber": {
          "type": "string"
        },
        "balance": {
          "type": "number",
          "format": "double",
          "x-go-custom-tag": "gorm:\"type:decimal(10,3)\""
        }
      }
    },
    "baseRequest": {
      "type": "object",
      "properties": {
        "accountNumber": {
          "type": "string"
        },
        "gold": {
          "type": "number",
          "format": "double"
        },
        "price": {
          "type": "number",
          "format": "double"
        }
      }
    },
    "baseResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "modelID": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "x-go-custom-tag": "gorm:\"type:int primary key auto_increment\""
        }
      }
    },
    "modelLogTime": {
      "type": "object",
      "properties": {
        "createdAt": {
          "type": "integer",
          "format": "date-time",
          "x-go-custom-tag": "gorm:\"type:int;autoCreateTime\""
        },
        "updatedAt": {
          "type": "integer",
          "format": "date-time",
          "x-go-custom-tag": "gorm:\"type:int;autoUpdateTime\""
        }
      }
    },
    "price": {
      "allOf": [
        {
          "$ref": "#/definitions/modelID"
        },
        {
          "$ref": "#/definitions/modelLogTime"
        },
        {
          "$ref": "#/definitions/priceData"
        }
      ]
    },
    "priceData": {
      "type": "object",
      "properties": {
        "buybackPrice": {
          "type": "number",
          "format": "double"
        },
        "topupPrice": {
          "type": "number",
          "format": "double"
        }
      }
    },
    "transaction": {
      "allOf": [
        {
          "$ref": "#/definitions/modelID"
        },
        {
          "$ref": "#/definitions/modelLogTime"
        },
        {
          "$ref": "#/definitions/accountData"
        },
        {
          "$ref": "#/definitions/priceData"
        },
        {
          "$ref": "#/definitions/transactionData"
        }
      ]
    },
    "transactionData": {
      "type": "object",
      "properties": {
        "gold": {
          "type": "number",
          "format": "double",
          "x-go-custom-tag": "gorm:\"type:decimal(10,3)\""
        },
        "type": {
          "type": "string",
          "enum": [
            "topup",
            "buyback"
          ]
        }
      }
    }
  },
  "parameters": {
    "accountNumber": {
      "type": "string",
      "name": "account_no",
      "in": "query",
      "required": true
    },
    "endDate": {
      "type": "integer",
      "name": "end_date",
      "in": "query",
      "required": true
    },
    "startDate": {
      "type": "integer",
      "name": "start_date",
      "in": "query",
      "required": true
    }
  }
}`))
}

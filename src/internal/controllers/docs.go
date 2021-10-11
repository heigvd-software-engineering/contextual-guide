// Package docs Contextual Guide
//
// Documentation of the open API of the Contextual-guide project
//
//     Schemes: http
//     BasePath: /api
//     Version: 0.0.1
//     Host: localhost:3000
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
//     Security:
//     - apikey
//
//    SecurityDefinitions:
//    apikey:
//      type: apiKey
//      in: header
//      name: x-api-key
//
// swagger:meta
package controllers

import "main/src/internal/models"

// swagger:response resourceList
type ResourceListResp struct {
	// in:body
	Body []models.Resource
}

// swagger:response resource
type ResourceResp struct {
	// in:body
	Body models.Resource
}

// swagger:parameters getResource
type GetResourceParams struct {
	// name: uuid
	// in: path
	// description: The uuid of the resource
	// required: true
	// type: string
	Uuid string
}

// swagger:parameters postResource
type PostResourceParams struct {
	// in:body
	Body models.Resource
}

// swagger:parameters putResource
type PutResourceParams struct {
	// name: uuid
	// in: path
	// description: The uuid of the resource
	// required: true
	// type: string
	Uuid string

	// in:body
	Body models.Resource
}

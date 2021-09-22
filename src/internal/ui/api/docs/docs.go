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
package docs

import "main/src/internal/controllers"

// swagger:parameters resourceSaveCommand
type ResourceSaveCommandWrapper struct {
	// in:body
	Body controllers.ResourceSaveCommand
}

// An JSON representation of the resource
// swagger:response resourceDTO
type ResourceDTOWrapper struct {
	// in:body
	Body controllers.ResourceDTO
}

// A list of resourceDTO
// swagger:response resourceDTOList
type ResourceDTOListWrapper struct {
	// in:body
	Body []controllers.ResourceDTO
}
// swagger:parameters resourceGetById
type ResourceGetByIdWrapper struct {

	// name: uuid
	// in: path
	// description: The uuid of the resource
	// required: true
	// type: string
	Uuid string
}
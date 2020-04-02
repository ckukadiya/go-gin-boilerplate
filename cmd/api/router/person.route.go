package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-boilerplate/internal/modules/person"
)

func NewPerson(p *person.PersonController, r *gin.RouterGroup) {
	p.Prepare()

	personRoute := r.Group("/person")

	// swagger:operation GET /person/list person getPersons
	//
	// Returns list of all people.
	//
	// ---
	// consumes:
	//   - "application/json"
	// produces:
	//   - "application/json"
	// responses:
	//   '200':
	//     description: person list response
	//     schema:
	//	        "$ref": "#/definitions/PersonListResponse"
	//
	//   default:
	//      description: General error
	//      schema:
	//	        "$ref": "#/definitions/GeneralError"
	//
	personRoute.GET("/list", p.GetAll)

	// swagger:operation GET /person/get/{id} person getPerson
	//
	// Returns person details of given person id.
	//
	// ---
	// consumes:
	//   - "application/json"
	// produces:
	//   - "application/json"
	// parameters:
	//   -
	//     in: "path"
	//     name: "id"
	//     description: "Person id which is require to fetch person details."
	//     required: true
	//     schema:
	//       type: string
	// responses:
	//   '200':
	//     description: person get response
	//     schema:
	//	        "$ref": "#/definitions/PersonResponse"
	//
	//   default:
	//      description: General error
	//      schema:
	//	        "$ref": "#/definitions/GeneralError"
	//
	personRoute.GET("/get/:id", p.Get)

	// swagger:operation POST /person/new person addPerson
	//
	// Insert given new person details in people.
	//
	// ---
	// consumes:
	//   - "application/json"
	// produces:
	//   - "application/json"
	// parameters:
	//   -
	//     in: "body"
	//     name: "body"
	//     description: "Person object that needs to be added to the people"
	//     required: true
	//     schema:
	//          "$ref": "#/definitions/Person"
	// responses:
	//   '200':
	//     description: person add response
	//     schema:
	//       type: object
	//       required:
	//         - id
	//       properties:
	//         id:
	//           type: string
	//
	//   default:
	//      description: General error
	//      schema:
	//	        "$ref": "#/definitions/GeneralError"
	//
	personRoute.POST("/new", p.Post)

	// swagger:operation PUT /person/update/{id} person updatePerson
	//
	// Update given person details in people.
	//
	// ---
	// consumes:
	//   - "application/json"
	// produces:
	//   - "application/json"
	// parameters:
	//   -
	//     in: "path"
	//     name: "id"
	//     description: "Person id which is require to fetch person details."
	//     required: true
	//     schema:
	//       type: string
	//   -
	//     in: "body"
	//     name: "body"
	//     description: "Person object that needs to be update in the people"
	//     required: true
	//     schema:
	//          "$ref": "#/definitions/Person"
	// responses:
	//   '200':
	//     description: person update response
	//     schema:
	//	        "$ref": "#/definitions/PersonResponse"
	//
	//   default:
	//      description: General error
	//      schema:
	//	        "$ref": "#/definitions/GeneralError"
	//
	personRoute.PUT("/update/:id", p.Put)

	// swagger:operation DELETE /person/remove/{id} person deletePerson
	//
	// Delete person details of given person id.
	//
	// ---
	// consumes:
	//   - "application/json"
	// produces:
	//   - "application/json"
	// parameters:
	//   -
	//     in: "path"
	//     name: "id"
	//     description: "Person id which is require to delete person details."
	//     required: true
	//     schema:
	//       type: string
	// responses:
	//   '200':
	//     description: person delete response
	//     schema:
	//       type: string
	//
	//   default:
	//      description: General error
	//      schema:
	//	        "$ref": "#/definitions/GeneralError"
	//
	personRoute.DELETE("/remove/:id", p.Delete)
}

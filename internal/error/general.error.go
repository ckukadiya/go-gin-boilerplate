package apperror

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GeneralError is the default error struct containing detailed information about the error
type GeneralError struct {
	// HTTP Status code to be set in response
	Status int `json:"-"`
	// Message is the error message that may be displayed to end users
	Message string `json:"message,omitempty"`
}

var (
	// Generic error
	Generic = NewStatus(http.StatusInternalServerError)
	// DB represents database related errors
	DB = NewStatus(http.StatusInternalServerError)
	// Forbidden represents access to forbidden resource error
	Forbidden = NewStatus(http.StatusForbidden)
	// BadRequest represents error for bad requests
	BadRequest = NewStatus(http.StatusBadRequest)
	// NotFound represents errors for not found resources
	NotFound = NewStatus(http.StatusNotFound)
	// Unauthorized represents errors for unauthorized requests
	Unauthorized = NewStatus(http.StatusUnauthorized)
)

// NewStatus generates new error containing only http status code
func NewStatus(status int) *GeneralError {
	return &GeneralError{Status: status}
}

// New generates an application error
func New(status int, msg string) *GeneralError {
	return &GeneralError{Status: status, Message: msg}
}

// Error returns the error message.
func (e GeneralError) Error() string {
	return e.Message
}

var validationErrors = map[string]string{
	"required": " is required, but was not received",
	"min":      "'s value or length is less than allowed",
	"max":      "'s value or length is bigger than allowed",
}

func getVldErrorMsg(s string) string {
	if v, ok := validationErrors[s]; ok {
		return v
	}
	return " failed on " + s + " validation"
}

// Response writes an error response to client
func Response(c *gin.Context, err error) {
	switch err.(type) {
	case *GeneralError:
		e := err.(*GeneralError)
		if e.Message == "" {
			log.Println("Error from Response : ", e)
			c.AbortWithStatus(e.Status)
		} else {
			log.Println("Error from Response : ", e.Message)
			c.AbortWithStatusJSON(e.Status, e)
		}
	default:
		log.Println("Error from Response : ", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}

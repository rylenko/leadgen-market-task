package ginapi

import "github.com/gin-gonic/gin"

// Handlers error type.
type Error struct {
	Code int       `json:"code"`
	Message string `json:"message"`
}

// Pushes error's JSON to the passed context.
func (e *Error) Push(c *gin.Context) {
	c.JSON(e.Code, e)
}

// Creates a new error with passed code and message.
func NewError(code int, message string) *Error {
	return &Error{
		Code: code,
		Message: message,
	}
}

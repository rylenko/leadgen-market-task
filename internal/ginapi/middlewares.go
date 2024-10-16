package ginapi

import (
	"log"

	"github.com/gin-gonic/gin"
)

func printErrorsMiddleware(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		for _, err := range c.Errors {
			log.Printf("Error: %v", err)
		}
	}
}

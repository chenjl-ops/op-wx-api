package header

import (
	"github.com/chenjl-ops/go-lib/uuid"
	"github.com/gin-gonic/gin"
	"strings"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get("X-Request-Id")
		if requestID == "" {
			newRequestID, err := uuid.NewV4()
			if err == nil {
				c.Writer.Header().Set("X-Request-Id", strings.Replace(newRequestID.String(), "-", "", -1))
			}
		} else {
			c.Writer.Header().Set("X-Request-Id", requestID)
		}
		c.Next()
	}
}

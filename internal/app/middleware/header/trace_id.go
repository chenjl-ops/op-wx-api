package header

import (
	"github.com/chenjl-ops/go-lib/uuid"
	"github.com/gin-gonic/gin"
	"strings"
)

func TraceIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.Request.Header.Get("TraceId")
		if traceID == "" {
			newTraceID, err := uuid.NewV1()
			if err == nil {
				c.Writer.Header().Set("TraceId", strings.Replace(newTraceID.String(), "-", "", -1))
			}
		} else {
			c.Writer.Header().Set("TraceId", traceID)
		}
		c.Next()
	}
}

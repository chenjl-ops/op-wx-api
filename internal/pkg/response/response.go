package response

import "github.com/gin-gonic/gin"

func WithError(c *gin.Context, code int, message any) {
	requestID := c.Writer.Header().Get("X-Request-Id")

	c.AbortWithStatusJSON(code, gin.H{
		"request_id":    requestID,
		"code":          code,
		"error_message": message,
	})
	return
}

func WithResult(c *gin.Context, code int, message string, data any) {
	requestID := c.Writer.Header().Get("X-Request-Id")

	c.JSON(code, Response{
		RequestId: requestID,
		Code:      code,
		Data:      data,
		Message:   message,
	})
}

func WithRequestResult(c *gin.Context, code int, message any, requestData any, responseData any) {
	requestID := c.Writer.Header().Get("X-Request-Id")

	c.JSON(code, gin.H{
		"code":          code,
		"request_id":    requestID,
		"message":       message,
		"request_data":  requestData,
		"response_data": responseData,
	})
}

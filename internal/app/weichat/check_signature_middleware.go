package weichat

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"op-wx-api/internal/pkg/response"
)

func CheckSignatureMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		WxSignature := c.Query("signature")
		Timestamp := c.Query("timestamp")
		Nonce := c.Query("nonce")
		//EchoString := c.Query("echostr")

		if checkWxGetSignature(Timestamp, Nonce, WxSignature) {
			c.Next()
		} else {
			response.WithError(c, http.StatusUnauthorized, "Signature check error, not from WX")
			return
		}
	}
}

package weichat

import "github.com/gin-gonic/gin"

func Url(r *gin.Engine) {
	v1 := r.Group("/v1/wx")
	{
		v1.GET("", CheckWxCallBack)
		v1.POST("", CheckSignatureMiddleware(), WxCallBackMain)
	}
}

package weichat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"op-wx-api/internal/pkg/response"
)

func CheckWxCallBack(c *gin.Context) {
	WxSignature := c.Query("signature")
	Timestamp := c.Query("timestamp")
	Nonce := c.Query("nonce")
	EchoString := c.Query("echostr")

	if checkWxGetSignature(Timestamp, Nonce, WxSignature) {
		fmt.Println("check is ok")
		c.String(http.StatusOK, EchoString)
	} else {
		fmt.Println("check is not ok")
		response.WithResult(c, http.StatusOK, "Signature check error", nil)
	}
}

func WxCallBackMain(c *gin.Context) {
	var wxBase WxBase
	if err := c.ShouldBindBodyWith(&wxBase, binding.XML); err != nil {
		fmt.Println("WX Base data error: ", err)
		c.XML(http.StatusOK, nil)
	} else {
		switch wxBase.MsgType {
		case "text":
			var data WxMessage
			if err = c.ShouldBindBodyWith(&data, binding.XML); err != nil {
				fmt.Println("WX POST data: ", data)
				fmt.Println("WX Message data error: ", err)
				c.XML(http.StatusOK, nil)
			} else {
				c.XML(http.StatusOK, wxMessageHandler(&data))
			}
		case "image":
			var data WxImageMessage
			if err = c.ShouldBindBodyWith(&data, binding.XML); err != nil {
				fmt.Println("WX Message image data error: ", err)
			} else {
				c.XML(http.StatusOK, wxImageHandler(&data))
			}
		case "voice":
			var data WxVoiceMessage
			if err = c.ShouldBindBodyWith(&data, binding.XML); err != nil {
				fmt.Println("WX Message voice data error: ", err)
			} else {
				c.XML(http.StatusOK, wxVoiceHandler(&data))
			}
		default:
			fmt.Println("Other MsgType: ", wxBase.MsgType)
			c.XML(http.StatusOK, nil)
		}
	}
}

func wxMessageHandler(message *WxMessage) *ReplyMessage {
	return getMessage(message)
}

func wxImageHandler(message *WxImageMessage) *ReplyImage {
	return getImageMessage(message)
}

func wxVoiceHandler(message *WxVoiceMessage) *ReplyVoice {
	return getVoiceMessage(message)
}

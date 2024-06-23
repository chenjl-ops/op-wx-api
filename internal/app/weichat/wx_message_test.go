package weichat

import (
	"fmt"
	"testing"
	"time"
)

func TestGetMessage(t *testing.T) {
	message := &WxMessage{
		MsgType:      "text",
		FromUserName: "from_test",
		ToUserName:   "to_test",
		Content:      "test content",
		CreateTime:   time.Now().Unix(),
	}
	data := getMessage(message)
	fmt.Println(data)
}

func TestGetImageMessage(t *testing.T) {
	message := &WxImageMessage{
		MsgType:      "image",
		FromUserName: "from_test",
		ToUserName:   "to_test",
		MediaId:      "testImageId",
		PicUrl:       "https://test.com/a.jpeg",
		CreateTime:   time.Now().Unix(),
	}
	data := getImageMessage(message)
	fmt.Println(data)
}

func TestGetVoiceMessage(t *testing.T) {
	message := &WxVoiceMessage{
		MsgType:      "voice",
		FromUserName: "from_test",
		ToUserName:   "to_test",
		MediaId:      "testVoiceId",
		CreateTime:   time.Now().Unix(),
		Format:       "VoiceTest",
	}
	data := getVoiceMessage(message)
	fmt.Println(data)
}

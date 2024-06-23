package weichat

import "time"

func getMessage(message *WxMessage) *ReplyMessage {
	return &ReplyMessage{
		ToUserName:   message.FromUserName,
		FromUserName: message.ToUserName,
		MsgType:      "text",
		Content:      message.Content,
		CreateTime:   time.Now().Unix(),
	}
}

func getImageMessage(message *WxImageMessage) *ReplyImage {
	imageMedia := ImageMedia{
		MediaId: message.MediaId,
	}

	return &ReplyImage{
		ToUserName:   message.FromUserName,
		FromUserName: message.ToUserName,
		MsgType:      "image",
		CreateTime:   time.Now().Unix(),
		Image:        imageMedia,
	}

}

func getVoiceMessage(message *WxVoiceMessage) *ReplyVoice {
	voiceMedia := VoiceMedia{
		MediaId: message.MediaId,
	}

	return &ReplyVoice{
		ToUserName:   message.FromUserName,
		FromUserName: message.ToUserName,
		MsgType:      "voice",
		CreateTime:   time.Now().Unix(),
		Voice:        voiceMedia,
	}
}

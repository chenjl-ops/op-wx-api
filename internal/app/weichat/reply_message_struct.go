package weichat

import "encoding/xml"

// ReplyMessage
/*
参数	是否必须	描述
ToUserName	是	接收方账号（收到的OpenID）
FromUserName	是	开发者微信号
CreateTime	是	消息创建时间 （整型）
MsgType	是	消息类型，文本为text
Content	是	回复的消息内容（换行：在content中能够换行，微信客户端就支持换行显示）
*/
type ReplyMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `json:"to_user" xml:"ToUserName"`
	FromUserName string   `json:"from_user" xml:"FromUserName"`
	CreateTime   int64    `json:"create_time" xml:"CreateTime"`
	MsgType      string   `json:"msg_type" xml:"MsgType"`
	Content      string   `json:"content" xml:"Content"`
}

// ReplyImage
/*
参数	是否必须	说明
ToUserName	是	接收方账号（收到的OpenID）
FromUserName	是	开发者微信号
CreateTime	是	消息创建时间戳 （整型）
MsgType	是	消息类型，语音为voice
MediaId	是	通过素材管理中的接口上传多媒体文件，得到的id
*/
type ReplyImage struct {
	XMLName      xml.Name   `xml:"xml"`
	ToUserName   string     `json:"to_user" xml:"ToUserName"`
	FromUserName string     `json:"from_user" xml:"FromUserName"`
	CreateTime   int64      `json:"create_time" xml:"CreateTime"`
	MsgType      string     `json:"msg_type" xml:"MsgType"`
	Image        ImageMedia `json:"media_id" xml:"Image"`
}

type ImageMedia struct {
	XMLName xml.Name `xml:"Image"`
	MediaId string   `json:"media_id" xml:"MediaId"`
}

// ReplyVoice
/*
参数	是否必须	说明
ToUserName	是	接收方账号（收到的OpenID）
FromUserName	是	开发者微信号
CreateTime	是	消息创建时间戳 （整型）
MsgType	是	消息类型，语音为voice
MediaId	是	通过素材管理中的接口上传多媒体文件，得到的id
*/
type ReplyVoice struct {
	XMLName      xml.Name   `xml:"xml"`
	ToUserName   string     `json:"to_user" xml:"ToUserName"`
	FromUserName string     `json:"from_user" xml:"FromUserName"`
	CreateTime   int64      `json:"create_time" xml:"CreateTime"`
	MsgType      string     `json:"msg_type" xml:"MsgType"`
	Voice        VoiceMedia `json:"media_id" xml:"Voice"`
}

type VoiceMedia struct {
	XMLName xml.Name `xml:"Voice"`
	MediaId string   `json:"media_id" xml:"MediaId"`
}

// ReplyVideo
/*
参数	是否必须	说明
ToUserName	是	接收方账号（收到的OpenID）
FromUserName	是	开发者微信号
CreateTime	是	消息创建时间 （整型）
MsgType	是	消息类型，视频为video
MediaId	是	通过素材管理中的接口上传多媒体文件，得到的id
Title	否	视频消息的标题
Description	否	视频消息的描述
*/
type ReplyVideo struct {
	XMLName      xml.Name   `xml:"xml"`
	ToUserName   string     `json:"to_user" xml:"ToUserName"`
	FromUserName string     `json:"from_user" xml:"FromUserName"`
	CreateTime   int64      `json:"create_time" xml:"CreateTime"`
	MsgType      string     `json:"msg_type" xml:"MsgType"`
	Video        VideoMedia `json:"video" xml:"Video"`
}

type VideoMedia struct {
	XMLName     xml.Name `xml:"Video"`
	MediaId     string   `json:"media_id" xml:"MediaId"`
	Title       string   `json:"title" xml:"Title"`
	Description string   `json:"description" xml:"Description"`
}

// ReplyMusic
/*
参数	是否必须	说明
ToUserName	是	接收方账号（收到的OpenID）
FromUserName	是	开发者微信号
CreateTime	是	消息创建时间 （整型）
MsgType	是	消息类型，音乐为music
Title	否	音乐标题
Description	否	音乐描述
MusicURL	否	音乐链接
HQMusicUrl	否	高质量音乐链接，WIFI环境优先使用该链接播放音乐
ThumbMediaId	是	缩略图的媒体id，通过素材管理中的接口上传多媒体文件，得到的id
*/
type ReplyMusic struct {
	XMLName      xml.Name   `xml:"xml"`
	ToUserName   string     `json:"to_user" xml:"ToUserName"`
	FromUserName string     `json:"from_user" xml:"FromUserName"`
	CreateTime   int64      `json:"create_time" xml:"CreateTime"`
	MsgType      string     `json:"msg_type" xml:"MsgType"`
	Music        MusicMedia `json:"music" xml:"Music"`
}

type MusicMedia struct {
	XMLName      xml.Name `xml:"Music"`
	Title        string   `json:"title" xml:"Title"`
	Description  string   `json:"description" xml:"Description"`
	MusicUrl     string   `json:"music_url" xml:"MusicUrl"`
	HQMusicUrl   string   `json:"hq_music_url" xml:"HQMusicUrl"`
	ThumbMediaId string   `json:"thumb_media_id" xml:"ThumbMediaId"`
}

// ReplyNews
/*
参数	是否必须	说明
ToUserName	是	接收方账号（收到的OpenID）
FromUserName	是	开发者微信号
CreateTime	是	消息创建时间 （整型）
MsgType	是	消息类型，图文为news
ArticleCount	是	图文消息个数；当用户发送文本、图片、语音、视频、图文、地理位置这六种消息时，开发者只能回复1条图文消息；其余场景最多可回复8条图文消息
Articles	是	图文消息信息，注意，如果图文数超过限制，则将只发限制内的条数
Title	是	图文消息标题
Description	是	图文消息描述
PicUrl	是	图片链接，支持JPG、PNG格式，较好的效果为大图360*200，小图200*200
Url	是	点击图文消息跳转链接
*/
type ReplyNews struct {
	XMLName      xml.Name  `xml:"xml"`
	ToUserName   string    `json:"to_user" xml:"ToUserName"`
	FromUserName string    `json:"from_user" xml:"FromUserName"`
	CreateTime   int64     `json:"create_time" xml:"CreateTime"`
	MsgType      string    `json:"msg_type" xml:"MsgType"`
	ArticleCount int       `json:"article_count" xml:"ArticleCount"`
	Articles     []Article `json:"articles" xml:"Articles"`
	Title        string    `json:"title" xml:"Title"`
	Description  string    `json:"description" xml:"Description"`
	PicUrl       string    `json:"pic_url" xml:"PicUrl"`
	URL          string    `json:"url" xml:"Url"`
}

type Article struct {
	XMlName     xml.Name `json:"xml_name" xml:"item"`
	Title       string   `json:"title" xml:"Title"`
	Description string   `json:"description" xml:"Description"`
	PicUrl      string   `json:"pic_url" xml:"PicUrl"`
	URL         string   `json:"url" xml:"Url"`
}

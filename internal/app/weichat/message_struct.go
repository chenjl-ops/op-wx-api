package weichat

import "encoding/xml"

/*
Developer Url：
https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Receiving_standard_messages.html
*/

// WxBase
/*
主要用于验证微信MsgType类型 后面好处理类型绑定数据
*/
type WxBase struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
}

//WxMessage
/*
参数	描述
ToUserName	开发者微信号
FromUserName	发送方账号（一个OpenID）
CreateTime	消息创建时间 （整型）
MsgType	消息类型，文本为text
Content	文本消息内容
MsgId	消息id，64位整型
MsgDataId	消息的数据ID（消息如果来自文章时才有）
Idx	多图文时第几篇文章，从1开始（消息如果来自文章时才有）
*/
type WxMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
	MsgId        string   `xml:"MsgId"`
	MsgDataId    string   `xml:"MsgDataId"`
	Idx          string   `xml:"Idx"`
}

// WxImageMessage
/*
参数	描述
ToUserName	开发者微信号
FromUserName	发送方账号（一个OpenID）
CreateTime	消息创建时间 （整型）
MsgType	消息类型，图片为image
PicUrl	图片链接（由系统生成）
MediaId	图片消息媒体id，可以调用获取临时素材接口拉取数据。
MsgId	消息id，64位整型
MsgDataId	消息的数据ID（消息如果来自文章时才有）
Idx	多图文时第几篇文章，从1开始（消息如果来自文章时才有）
*/
type WxImageMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	PicUrl       string   `xml:"PicUrl"`
	MediaId      string   `xml:"MediaId"`
	MsgId        string   `xml:"MsgId"`
	MsgDataId    string   `xml:"MsgDataId"`
	Idx          string   `xml:"Idx"`
}

// WxVoiceMessage
/*
参数	描述
ToUserName	开发者微信号
FromUserName	发送方账号（一个OpenID）
CreateTime	消息创建时间 （整型）
MsgType	语音为voice
MediaId	语音消息媒体id，可以调用获取临时素材接口拉取数据，Format为amr时返回8K采样率amr语音。
Format	语音格式，如amr，speex等
MsgId	消息id，64位整型
MsgDataId	消息的数据ID（消息如果来自文章时才有）
Idx	多图文时第几篇文章，从1开始（消息如果来自文章时才有）
MediaId16K	16K采样率语音消息媒体id，可以调用获取临时素材接口拉取数据，返回16K采样率amr/speex语音。
*/
type WxVoiceMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	MediaId      string   `xml:"MediaId"`
	Format       string   `xml:"Format"`
	MsgId        string   `xml:"MsgId"`
	MsgDataId    string   `xml:"MsgDataId"`
	Idx          string   `xml:"Idx"`
	MediaId16K   string   `xml:"MediaId16K"`
}

// WxVideoMessage
/*
参数	描述
ToUserName	开发者微信号
FromUserName	发送方账号（一个OpenID）
CreateTime	消息创建时间 （整型）
MsgType	视频为video
MediaId	视频消息媒体id，可以调用获取临时素材接口拉取数据。
ThumbMediaId	视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
MsgId	消息id，64位整型
MsgDataId	消息的数据ID（消息如果来自文章时才有）
Idx	多图文时第几篇文章，从1开始（消息如果来自文章时才有）
*/
type WxVideoMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	MediaId      string   `xml:"MediaId"`
	ThumbMediaId string   `xml:"ThumbMediaId"`
	MsgId        string   `xml:"MsgId"`
	MsgDataId    string   `xml:"MsgDataId"`
	Idx          string   `xml:"Idx"`
}

// WxShortVideoMessage
/*
参数	描述
ToUserName	开发者微信号
FromUserName	发送方账号（一个OpenID）
CreateTime	消息创建时间 （整型）
MsgType	小视频为shortvideo
MediaId	视频消息媒体id，可以调用获取临时素材接口拉取数据。
ThumbMediaId	视频消息缩略图的媒体id，可以调用获取临时素材接口拉取数据。
MsgId	消息id，64位整型
MsgDataId	消息的数据ID（消息如果来自文章时才有）
Idx	多图文时第几篇文章，从1开始（消息如果来自文章时才有）
*/
type WxShortVideoMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	MediaId      string   `xml:"MediaId"`
	ThumbMediaId string   `xml:"ThumbMediaId"`
	MsgId        string   `xml:"MsgId"`
	MsgDataId    string   `xml:"MsgDataId"`
	Idx          string   `xml:"Idx"`
}

// WxLocationMessage
/*
参数	描述
ToUserName	开发者微信号
FromUserName	发送方账号（一个OpenID）
CreateTime	消息创建时间 （整型）
MsgType	消息类型，地理位置为location
Location_X	地理位置纬度
Location_Y	地理位置经度
Scale	地图缩放大小
Label	地理位置信息
MsgId	消息id，64位整型
MsgDataId	消息的数据ID（消息如果来自文章时才有）
Idx	多图文时第几篇文章，从1开始（消息如果来自文章时才有）
*/
type WxLocationMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	LocationX    float64  `xml:"LocationX"`
	LocationY    float64  `xml:"LocationY"`
	Scale        int64    `xml:"Scale"`
	Label        string   `xml:"Label"`
	MsgId        string   `xml:"MsgId"`
	MsgDataId    string   `xml:"MsgDataId"`
	Idx          string   `xml:"Idx"`
}

// WxLinkMessage
/*
参数	描述
ToUserName	接收方微信号
FromUserName	发送方微信号，若为普通用户，则是一个OpenID
CreateTime	消息创建时间
MsgType	消息类型，链接为link
Title	消息标题
Description	消息描述
Url	消息链接
MsgId	消息id，64位整型
MsgDataId	消息的数据ID（消息如果来自文章时才有）
Idx	多图文时第几篇文章，从1开始（消息如果来自文章时才有）
*/
type WxLinkMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Title        string   `xml:"Title"`
	Description  string   `xml:"Description"`
	Url          string   `xml:"Url"`
	MsgId        string   `xml:"MsgId"`
	MsgDataId    string   `xml:"MsgDataId"`
	Idx          string   `xml:"Idx"`
}

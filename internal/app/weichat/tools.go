package weichat

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"op-wx-api/internal/pkg/conf"
	"sort"
	"strings"
	"time"
)

// NewStableAccessToken 根据wx token 生成带创建时间token
func NewStableAccessToken(token *AccessToken) *StableAccessToken {
	sat := &StableAccessToken{
		AccessToken: token.AccessToken,
		ExpiresIn:   token.ExpiresIn,
		CreateTime:  time.Now().Unix(),
	}
	return sat
}

// getToken 获取token信息
func (wToken *AccessToken) getToken() string {
	return wToken.AccessToken
}

// getToken 获取token信息
func (sToken *StableAccessToken) getToken() string {
	return sToken.AccessToken
}

// getExpires 获取有效期
func (wToken *AccessToken) getExpires() int64 {
	return wToken.ExpiresIn
}

// getExpires 获取有效期
func (sToken *StableAccessToken) getExpires() int64 {
	return sToken.ExpiresIn
}

// getCreateTime 获取token创建时间
func (sToken *StableAccessToken) getCreateTime() int64 {
	return sToken.CreateTime
}

// isExpires 判断token是否过期
func (sToken *StableAccessToken) isExpires() bool {
	if time.Now().Unix()-(sToken.ExpiresIn+sToken.CreateTime) > 0 {
		return true
	} else {
		return false
	}
}

// DecodeWxMessage 微信消息解密
func DecodeWxMessage(encodeMessage string) any {
	AESKey, err := base64.StdEncoding.DecodeString(fmt.Sprintf("%s=", conf.NacosConfig.WXEncodeingAESKey))
	if err != nil {
		fmt.Println("base64 decode key err:", err)
	}

	TmpMsg, err := base64.StdEncoding.DecodeString(encodeMessage)
	if err != nil {
		fmt.Println("base64 decode message err:", err)
	}

	text := string(AesDecryptCBC(TmpMsg, AESKey))
	padStr := []rune(text[len(text):])
	pad := int(padStr[0])

	content := text[16 : len(text)-pad] // 真实消息内容
	//msgLen := content[:4]
	//msg := content[4 : msgLen+4]
	//fromAppid := content[msgLen+4:]

	return content
}

// EncodeWxMessage 微信消息加密
func EncodeWxMessage() any {
	return nil
}

// AesDecryptCBC AES CBC decrypt
func AesDecryptCBC(encrypted []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	decrypted := make([]byte, len(encrypted))
	blockMode.CryptBlocks(decrypted, encrypted)
	decrypted = pkcs5UnPadding(decrypted)
	return decrypted
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// 校验签名来源是否来自微信
func checkWxGetSignature(timestamp string, nonce string, wxSignature string) bool {
	list := []string{conf.NacosConfig.WXToken, timestamp, nonce}
	sort.Strings(list)

	s := sha1.New()
	s.Write([]byte(strings.Join(list, "")))
	hashcode := hex.EncodeToString(s.Sum(nil))

	if hashcode != wxSignature {
		fmt.Println("is not ok")
		return false
	} else {
		fmt.Println("is ok")
		return true
	}
}

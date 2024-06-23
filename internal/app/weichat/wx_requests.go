package weichat

import (
	"fmt"
	"github.com/chenjl-ops/go-lib/requests"
	"op-wx-api/internal/pkg/conf"
)

func getAccessToken() (*AccessToken, error) {
	url := "https://api.weixin.qq.com/cgi-bin/stable_token"

	requestData := map[string]string{"grant_type": "client_credential", "appid": conf.NacosConfig.WXAppId, "secret": conf.NacosConfig.WXAppSecret}
	var responseData *AccessToken

	err := requests.Request(url, "POST", nil, requestData, &responseData)
	fmt.Println("access_token: ", responseData)
	return responseData, err
}

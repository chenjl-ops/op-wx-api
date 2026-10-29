package weichat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GeminiApi struct {
	Url       string
	ApiKey    string
	Prompt    string
	ModelName string
}

type GeminiApiResponse struct {
	Code      int    `json:"code"`
	RequestId string `json:"request_id"`
	Data      string `json:"data"`
	Message   string `json:"message"`
}

func (g *GeminiApi) getGemini() (message string, err error) {
	// 构造请求数据
	requestData := map[string]string{
		"prompt":     g.Prompt,
		"api_key":    g.ApiKey,
		"model_name": g.ModelName,
	}

	reqBody, _ := json.Marshal(requestData)

	resp, err := http.Post(
		g.Url,
		"application/json",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		fmt.Println("request error:", err)
		return "", err
	}

	defer resp.Body.Close()

	// 解析返回的 JSON
	body, _ := io.ReadAll(resp.Body)

	// 打印 HTTP 状态码和原始响应
	fmt.Println("HTTP status:", resp.Status)
	fmt.Println("Response headers:", resp.Header)
	fmt.Println("Response body:", string(body))

	var result *GeminiApiResponse
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("json load error:", err)
		return "", err
	}

	return result.Data, nil
}

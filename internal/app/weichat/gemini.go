package weichat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//// InteractionRequest Interactions API 请求体
//type InteractionRequest struct {
//	Model string `json:"model"`
//	Input string `json:"input"`
//}
//
//// StepContent 2. Interactions API 响应体结构
//type StepContent struct {
//	Type string `json:"type"`
//	Text string `json:"text"`
//}
//
//type InteractionStep struct {
//	Type    string        `json:"type"`
//	Content []StepContent `json:"content,omitempty"`
//}
//
//type InteractionResponse struct {
//	ID         string            `json:"id"`
//	Status     string            `json:"status"`
//	Model      string            `json:"model"`
//	OutputText string            `json:"output_text,omitempty"`
//	Steps      []InteractionStep `json:"steps,omitempty"`
//	Error      *struct {
//		Code    int    `json:"code"`
//		Message string `json:"message"`
//	} `json:"error,omitempty"`
//}
//
//// Text 提取 AI 回复内容的便捷方法
//func (r *InteractionResponse) Text() string {
//	if r.OutputText != "" {
//		return r.OutputText
//	}
//	for _, s := range r.Steps {
//		if s.Type == "model_output" {
//			for _, c := range s.Content {
//				if c.Text != "" {
//					return c.Text
//				}
//			}
//		}
//	}
//	return ""
//}
//
//// CallGeminiInteraction 核心方法
//func CallGeminiInteraction(apiKey, prompt string) (string, error) {
//	if apiKey == "" {
//		return "", fmt.Errorf("GEMINI_API_KEY 不能为空")
//	}
//
//	client := &http.Client{
//		Timeout: 20 * time.Second,
//	}
//
//	apiURL := "https://generativelanguage.googleapis.com/v1beta/interactions"
//
//	// 使用 InteractionRequest 结构体
//	reqPayload := InteractionRequest{
//		Model: "gemini-3-flash-preview",
//		Input: prompt,
//	}
//
//	jsonData, err := json.Marshal(reqPayload)
//	if err != nil {
//		return "", fmt.Errorf("JSON 序列化失败: %w", err)
//	}
//
//	req, err := http.NewRequestWithContext(context.Background(), "POST", apiURL, bytes.NewBuffer(jsonData))
//	if err != nil {
//		return "", fmt.Errorf("创建 Request 失败: %w", err)
//	}
//
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("x-goog-api-key", apiKey)
//
//	resp, err := client.Do(req)
//	if err != nil {
//		return "", fmt.Errorf("请求失败: %w", err)
//	}
//
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//
//		}
//	}(resp.Body)
//
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return "", fmt.Errorf("读取响应失败: %w", err)
//	}
//
//	if resp.StatusCode != http.StatusOK {
//		return "", fmt.Errorf("API 状态码 [%d]: %s", resp.StatusCode, string(body))
//	}
//
//	var result InteractionResponse
//	if err := json.Unmarshal(body, &result); err != nil {
//		return "", fmt.Errorf("解析 JSON 失败: %w", err)
//	}
//
//	if result.Error != nil {
//		return "", fmt.Errorf("业务报错: %s", result.Error.Message)
//	}
//
//	return result.Text(), nil
//}

type GeminiApi struct {
	Url    string
	ApiKey string
	Prompt string
}

type GeminiApiResponse struct {
	Code      int    `json:"code"`
	RequestId string `json:"request_id"`
	Data      any    `json:"data"`
	Message   string `json:"message"`
}

func (g *GeminiApi) getGemini() (message string, err error) {
	// 构造请求数据
	requestData := map[string]string{
		"prompt":  g.Prompt,
		"api_key": g.ApiKey,
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
	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("json load error:", err)
		return "", err
	}

	return result["data"].(string), nil
}

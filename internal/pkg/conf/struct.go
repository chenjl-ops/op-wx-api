package conf

type Specification struct {
	ServerRunPort         int64  `envconfig:"SERVER_RUN_PORT" mapstructure:"server_run_port" json:"SERVER_RUN_PORT"`
	WXToken               string `envconfig:"WX_TOKEN" mapstructure:"wx_token" json:"wx_token"`
	WXAppId               string `envconfig:"WX_APP_ID" mapstructure:"wx_app_id" json:"wx_app_id"`
	WXAppSecret           string `envconfig:"WX_APP_SECRET" mapstructure:"wx_app_secret" json:"wx_app_secret"`
	WXEncodeingAESKey     string `envconfig:"WX_ENCODEING_AES_KEY" mapstructure:"wx_encodeing_aes_key" json:"wx_encodeing_aes_key"`
	GeminiUrl             string `envconfig:"GEMINI_URL" mapstructure:"gemini_url" json:"gemini_url"`
	GeminiToken           string `envconfig:"GEMINI_TOKEN" mapstructure:"gemini_token" json:"gemini_token"`
	GeminiDefaultResponse string `envconfig:"GEMINI_DEFAULT_RESPONSE" mapstructure:"gemini_default_response" json:"gemini_default_response"`
	GeminiModelName       string `envconfig:"GEMINI_MODEL_NAME" mapstructure:"gemini_model_name" json:"gemini_model_name"`
}

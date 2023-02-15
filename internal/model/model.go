package model

// Config
// @Description: 配置
type Config struct {
	MsgEnabled   bool
	MsgType      string
	BarkUrl      string
	BarkKey      string
	MailHost     string
	MailProtocol string
	MailPort     int
	MailUser     string
	MailPwd      string
	MailFromName string
	MailTo       []string
}

// BarkRequest
// @Description: Bark请求
type BarkRequest struct {
	Body      string `json:"body"`
	DeviceKey string `json:"device_key"`
	Title     string `json:"title"`
	Badge     int    `json:"badge"`
	Category  string `json:"category"`
	Sound     string `json:"sound"`
	Icon      string `json:"icon"`
	Group     string `json:"group"`
	Url       string `json:"url"`
}

// BarkResponse
// @Description: Bark回复
type BarkResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
}

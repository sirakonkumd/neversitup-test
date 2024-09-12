package models

type StandardResponse struct {
	Code     string      `json:"code"`
	BizError string      `json:"bizError"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data,omitempty"`
}

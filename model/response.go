package model

type ResponseData struct {
	Status bool    `json:"status,omitempty"`
	Result int     `json:"result,omitempty"`
	Error  *string `json:"error,omitempty"`
}

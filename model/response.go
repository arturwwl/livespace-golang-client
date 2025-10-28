package model

// ResponseData represents basic api response data
type ResponseData struct {
	Status bool    `json:"status,omitempty"`
	Result int     `json:"result,omitempty"`
	Error  *string `json:"error,omitempty"`
}

package model

// GetToken represents api request data
type GetToken struct {
	ApiKey  string `form:"_api_key,omitempty"`
	ApiAuth string `form:"_api_auth,omitempty"`
}

type Token struct {
	Data *TokenData `json:"data,omitempty"`
	ResponseData
}

type TokenData struct {
	Token     string `json:"token,omitempty"`
	SessionID string `json:"session_id,omitempty"`
}

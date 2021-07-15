package model

type AuthorizedRequest struct {
	ApiKey     string `json:"_api_key,omitempty" form:"_api_key,omitempty"`
	ApiSHA     string `json:"_api_sha,omitempty" form:"_api_sha,omitempty"`
	ApiSession string `json:"_api_session,omitempty" form:"_api_session,omitempty"`
	Type       string `json:"type,omitempty" form:"type,omitempty"`
}

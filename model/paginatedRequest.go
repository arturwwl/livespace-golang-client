package model

type PaginatedRequest struct {
	Offset *string `json:"offset,omitempty" form:"offset,omitempty"`
	Limit  int     `json:"limit,omitempty" form:"limit,omitempty"`
}

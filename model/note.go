package model

// CreateNote stores create note api request data
type CreateNote struct {
	AuthorizedRequest
	Contact NoteData `json:"contact,omitempty"`
}

// NoteData stores create note api response data
type NoteData struct {
	PersonID string  `json:"id,omitempty"`
	Note     string  `json:"note,omitempty"`
	Tags     *string `json:"type,omitempty"`
}

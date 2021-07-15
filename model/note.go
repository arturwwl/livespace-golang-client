package model

type CreateNote struct {
	AuthorizedRequest
	Contact NoteData `json:"contact,omitempty"`
}

type NoteData struct {
	PersonID string  `json:"id,omitempty"`
	Note     string  `json:"note,omitempty"`
	Tags     *string `json:"type,omitempty"`
}

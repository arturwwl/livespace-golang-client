package livespaceclient

import (
	"encoding/json"
	"github.com/arturwwl/livespace-golang-client/model"
)

// CreateNote creates new note using api
func (c *LivespaceClient) CreateNote(noteM model.NoteData) error {
	var err error
	request := model.CreateNote{
		Contact: noteM,
	}
	request.AuthorizedRequest, err = c.prepareAuthorizedRequest()
	if err != nil {
		return err
	}

	var responseBytes []byte
	responseBytes, err = c.makeRequest("Contact/addContactNote", request, true)
	if err != nil {
		return err
	}

	contactSingle := model.ContactSingle{}
	err = json.Unmarshal(responseBytes, &contactSingle)
	if err != nil {
		return err
	}

	return nil
}

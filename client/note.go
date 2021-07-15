package livespace_client

import (
	"encoding/json"
	"github.com/arturwwl/livespace-golang-client/model"
)

func (c *LivespaceClient) CreateNote(noteM *model.NoteData) (err error) { //TODO: custom error
	request := model.CreateNote{
		Contact: *noteM,
	}
	request.AuthorizedRequest, err = c.prepareAuthorizedRequest()
	if err != nil {
		return err
	}

	if responseBytes, err := c.makeRequest("Contact/addContactNote", request, true); err != nil {
		return err
	} else {
		contactSingle := model.ContactSingle{}
		_ = json.Unmarshal(responseBytes, &contactSingle)
	}

	return nil
}

package livespace_client

import (
	"encoding/json"
	"github.com/arturwwl/livespace-golang-client/model"
)

func (c *LivespaceClient) CreateContact(contactM *model.ContactData) (err error) { //TODO: custom error
	request := model.CreateContact{
		Contact: *contactM,
	}
	request.AuthorizedRequest, err = c.prepareAuthorizedRequest()
	if err != nil {
		return err
	}

	if responseBytes, err := c.makeRequest("Contact/addContact", request, true); err != nil {
		return err
	} else {
		contactSingle := model.ContactSingle{}
		_ = json.Unmarshal(responseBytes, &contactSingle)
		contactM = &contactSingle.Data.Contact
	}

	return nil
}

func (c *LivespaceClient) GetContact(emails *string, firstNames *string, lastNames *string) (contactM *model.ContactData, err error) { //TODO: custom error
	request := model.GetContact{
		AuthorizedRequest: model.AuthorizedRequest{},
		PaginatedRequest:  model.PaginatedRequest{},
		Emails:            emails,
		Firstnames:        firstNames,
		Lastnames:         lastNames,
	}

	request.AuthorizedRequest, err = c.prepareAuthorizedRequest()
	if err != nil {
		return nil, err
	}
	request.AuthorizedRequest.Type = "contact"
	request.Limit = 1

	if responseBytes, err := c.makeRequest("Contact/getAll", request, false); err != nil {
		return contactM, err
	} else {
		contactList := model.ContactList{}
		_ = json.Unmarshal(responseBytes, &contactList)
		if len(contactList.Data.Contact) > 0 {
			return &contactList.Data.Contact[0], nil
		}
	}
	return nil, nil
}

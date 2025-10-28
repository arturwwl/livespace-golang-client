package livespaceclient

import (
	"encoding/json"
	"github.com/arturwwl/livespace-golang-client/model"
)

// CreateContact creates new contact using api
func (c *LivespaceClient) CreateContact(contactM *model.ContactData) error {
	var err error
	request := model.CreateContact{
		Contact: *contactM,
	}
	request.AuthorizedRequest, err = c.prepareAuthorizedRequest()
	if err != nil {
		return err
	}

	var responseBytes []byte
	responseBytes, err = c.makeRequest("Contact/addContact", request, true)
	if err != nil {
		return err
	}

	contactSingle := model.ContactSingle{}
	err = json.Unmarshal(responseBytes, &contactSingle)
	if err != nil {
		return err
	}

	contactM = &contactSingle.Data.Contact

	return nil
}

// GetContact searches for existing contact using api
func (c *LivespaceClient) GetContact(request model.ContactData) (contactM *model.ContactData, err error) {
	list, err := c.ListContact(request)
	if err != nil {
		return
	}

	if len(list) > 0 {
		return &list[0], nil
	}

	return nil, nil
}

// ListContact gets list for existing contacts using api
func (c *LivespaceClient) ListContact(request model.ContactData) (contactM []model.ContactData, err error) {
	aRequest := model.ListContact{
		AuthorizedRequest: model.AuthorizedRequest{},
		PaginatedRequest:  model.PaginatedRequest{},
		ContactData:       request,
	}

	aRequest.AuthorizedRequest, err = c.prepareAuthorizedRequest()
	if err != nil {
		return nil, err
	}

	aRequest.AuthorizedRequest.Type = "contact"
	aRequest.Limit = 1

	responseBytes, err := c.makeRequest("Contact/getAll", request, false)
	if err != nil {
		return contactM, err
	}

	var contactList model.ContactList
	err = json.Unmarshal(responseBytes, &contactList)
	if err != nil {
		return nil, err
	}

	return contactList.Data.Contact, nil
}

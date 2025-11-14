package livespaceclient

import (
	"encoding/json"
	"fmt"
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
func (c *LivespaceClient) GetContact(filters model.ListContactFilters) (*model.ContactData, error) {
	list, err := c.ListContact(filters)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return &list[0], nil
	}

	return nil, nil
}

// ListContact gets list for existing contacts using api
func (c *LivespaceClient) ListContact(filters model.ListContactFilters) ([]model.ContactData, error) {
	var err error
	request := model.ListContact{
		AuthorizedRequest:  model.AuthorizedRequest{},
		PaginatedRequest:   model.PaginatedRequest{},
		ListContactFilters: filters,
	}

	request.AuthorizedRequest, err = c.prepareAuthorizedRequest()
	if err != nil {
		return nil, err
	}

	request.AuthorizedRequest.Type = "contact"

	responseBytes, err := c.makeRequest("Contact/getAll", request, true)
	if err != nil {
		return nil, err
	}

	var contactList model.ContactList
	err = json.Unmarshal(responseBytes, &contactList)
	if err != nil {
		return nil, fmt.Errorf("json: %s error: %v", responseBytes, err)
	}

	return contactList.Data.Contact, nil
}

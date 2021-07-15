package test

import (
	livespace_client "github.com/arturwwl/livespace-golang-client/client"
	"github.com/arturwwl/livespace-golang-client/helper"
	"github.com/arturwwl/livespace-golang-client/model"
	"testing"
)

func TestGetContact(t *testing.T) {
	client, err := livespace_client.New("conf/cfg.ini")
	if err != nil {
		t.Fatal(err)
	}

	email := "testnil@example.com"
	contactM, err := client.GetContact(&email, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if contactM != nil {
		t.Fatal(err)
	}
}

func TestCreateContactNote(t *testing.T) {
	client, err := livespace_client.New("conf/cfg.ini")
	if err != nil {
		t.Fatal(err)
	}
	noteM := model.NoteData{
		PersonID: "00000000-aaaa-0000-aaaa-bbbbbbbbbbbb",
		Note:     "test note",
		Tags:     nil,
	}
	err = client.CreateNote(&noteM)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateContact(t *testing.T) {
	client, err := livespace_client.New("conf/cfg.ini")
	if err != nil {
		t.Fatal(err)
	}
	contactM := model.ContactData{
		FirstName:     "Lorem",
		LastName:      "Ipsum",
		ContactSource: helper.PointString("Contact Source"),
		Emails: &[]model.Email{
			{
				Email: "test@example.com",
			},
		},
		Phones: &[]model.Phone{
			{
				PhoneNo: "123123123",
			},
		},
	}
	err = client.CreateContact(&contactM)
	if err != nil {
		t.Fatal(err)
	}
}

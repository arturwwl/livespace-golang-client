package model

// ListContact represents api request for listing contacts
type ListContact struct {
	AuthorizedRequest
	PaginatedRequest
	ContactData
}

// CreateContact represents api request for creating contact
type CreateContact struct {
	AuthorizedRequest
	Contact ContactData `json:"contact,omitempty"`
}

// ContactSingle represents api response for creating contact
type ContactSingle struct {
	Data ContactSingleItem `json:"data,omitempty"`
	ResponseData
}

// ContactList represents api response for listing contacts
type ContactList struct {
	Data ContactListItem `json:"data,omitempty"`
	ResponseData
}

type ContactListItem struct {
	Contact []ContactData `json:"contact,omitempty"`
}

type ContactSingleItem struct {
	Contact ContactData `json:"contact,omitempty"`
}

type ContactData struct {
	ID               *string                 `json:"id,omitempty"`
	ContactID        *string                 `json:"contact_id,omitempty"`
	Type             *string                 `json:"type,omitempty"`
	Url              *string                 `json:"url,omitempty"`
	CompanyName      *string                 `json:"company_name,omitempty"`
	Phone            *string                 `json:"phone,omitempty"`
	Name             *string                 `json:"name,omitempty"`
	Groups           *[]string               `json:"groups,omitempty"`
	GroupsID         *[]string               `json:"groups_id,omitempty"`
	Dataset          *map[string]interface{} `json:"dataset,omitempty"`
	DatasetFieldName *map[string]*string     `json:"dataset_field_name,omitempty"`
	DatasetFieldType *map[string]*string     `json:"dataset_field_type,omitempty"`

	FirstName     string     `json:"firstname,omitempty"`
	LastName      string     `json:"lastname,omitempty"`
	Note          *string    `json:"note,omitempty"`
	Position      *string    `json:"position,omitempty"`
	ContactSource *string    `json:"contact_source,omitempty"`
	WWW           *string    `json:"www,omitempty"`
	Emails        *[]Email   `json:"emails,omitempty"`
	Phones        *[]Phone   `json:"phones,omitempty"`
	Addresses     *[]Address `json:"addresses,omitempty"`
}

type Address struct {
	Street     *string `json:"street,omitempty"`
	City       *string `json:"city,omitempty"`
	PostalCode *string `json:"postcode,omitempty"`
	Country    *string `json:"country,omitempty"`
}
type Email struct {
	Email     string `json:"email,omitempty"`
	IsDefault *bool  `json:"is_default,omitempty"`
}

type PhoneType int

const PhoneTypeTelephone PhoneType = 1
const PhoneTypeMobile PhoneType = 2
const PhoneTypeFax PhoneType = 3

type Phone struct {
	PhoneNo string    `json:"phone_no,omitempty"`
	Type    PhoneType `json:"type,omitempty"`
}

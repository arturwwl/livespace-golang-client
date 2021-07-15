package model

type GetContact struct {
	AuthorizedRequest
	PaginatedRequest
	Emails     *string `json:"emails,omitempty" form:"emails,omitempty"`
	Firstnames *string `json:"firstnames,omitempty" form:"firstnames,omitempty"`
	Lastnames  *string `json:"lastnames,omitempty" form:"lastnames,omitempty"`
	Phones     *string `json:"phones,omitempty" form:"phones,omitempty"`
}

type CreateContact struct {
	AuthorizedRequest
	Contact ContactData `json:"contact,omitempty"`
}

type ContactSingle struct {
	Data ContactSingleItem `json:"data,omitempty"`
	ResponseData
}

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

	FirstName     string   `json:"firstname,omitempty"`
	LastName      string   `json:"lastname,omitempty"`
	Note          *string  `json:"note,omitempty"`
	Position      *string  `json:"position,omitempty"`
	ContactSource *string  `json:"contact_source,omitempty"`
	WWW           *string  `json:"www,omitempty"`
	Emails        *[]Email `json:"emails,omitempty"`
	Phones        *[]Phone `json:"phones,omitempty"`
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

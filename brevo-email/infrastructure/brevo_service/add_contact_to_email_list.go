package brevoservice

import (
	"encoding/json"

	brevorequest "gitlab.com/eliotandelon/brevo-email/infrastructure/brevo_request"
	brevoresponse "gitlab.com/eliotandelon/brevo-email/infrastructure/brevo_response"
)

func (client BrevoMailClient) AddContactToEmailList(email string, listId int64, attributtes brevorequest.CreateContactAtributtesRequest) (*brevoresponse.ContactInUserListResponse, error) {
	contact := brevorequest.CreateContactRequest{
		Email:         email,
		ListIds:       []int64{listId},
		UpdateEnabled: true,
		Attributes:    attributtes,
	}
	jsonData, err := json.Marshal(contact)
	if err != nil {
		return nil, err
	}

	var result brevoresponse.ContactInUserListResponse
	err = client.post(client.urlBase, "/v3/contacts", jsonData, &result, getHeaders(client.brevoApiKey))
	if err != nil {
		return nil, err
	}

	return &result, nil
}

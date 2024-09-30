package brevointerface

import (
	brevorequest "gitlab.com/eliotandelon/brevo-email/infrastructure/brevo_request"
	brevoresponse "gitlab.com/eliotandelon/brevo-email/infrastructure/brevo_response"
)

type BrevoServiceInterface interface {
	AddContactToEmailList(email string, listId int64, attributtes brevorequest.CreateContactAtributtesRequest) (*brevoresponse.ContactInUserListResponse, error)
	GetAllList() (*brevoresponse.EmailListResponse, error)
	GetIdOfListByName(name string) (float64, error)
	SendMail(reqMail brevorequest.MailRequest, uriEndpoint string, data map[string]string) (*string, error)
}

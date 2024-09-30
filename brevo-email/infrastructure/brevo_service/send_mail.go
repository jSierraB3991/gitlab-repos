package brevoservice

import (
	"encoding/json"

	brevomapper "gitlab.com/eliotandelon/brevo-email/domain/brevo_mapper"
	brevorequest "gitlab.com/eliotandelon/brevo-email/infrastructure/brevo_request"
	brevoresponse "gitlab.com/eliotandelon/brevo-email/infrastructure/brevo_response"
)

func (client *BrevoMailClient) SendMail(reqMail brevorequest.MailRequest, uriEndpoint string, data map[string]string) (*string, error) {

	smtpEMail := brevomapper.ToBrevoMailByMailRequest(reqMail, client.mailTemplatePath+"/"+reqMail.HtmlName, data)
	jsonData, err := json.Marshal(smtpEMail)
	if err != nil {
		return nil, err
	}
	var result brevoresponse.MailSendResponse

	err = client.post(client.urlBase, uriEndpoint, jsonData, &result, getHeaders(client.brevoApiKey))
	if err != nil {
		if err.Error() == "201 Created" {
			return &result.MessageId, nil
		}
		return nil, err
	}
	return &result.MessageId, nil
}

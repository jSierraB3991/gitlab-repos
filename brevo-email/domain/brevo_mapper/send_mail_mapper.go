package brevomapper

import (
	brevorequest "gitlab.com/eliotandelon/brevo-email/infrastructure/brevo_request"
	eliotlibs "gitlab.com/eliotandelon/eliot-libs"
)

func ToBrevoMailByMailRequest(requestMail brevorequest.MailRequest, fileContenttHtml string, data map[string]string) brevorequest.BrevoRequestSendMail {
	result := brevorequest.BrevoRequestSendMail{
		Subject:     requestMail.Subject,
		HtmlContent: eliotlibs.ReplaceTextInFile(fileContenttHtml, data),
		Sender: brevorequest.DataBrevoRequestMail{
			Name:  requestMail.SenderDataMail.Name,
			Email: requestMail.SenderDataMail.Email,
		},
		To: []brevorequest.DataBrevoRequestMail{
			{
				Email: requestMail.ToDataMail.Email,
				Name:  requestMail.ToDataMail.Name,
			},
		},
	}

	return result
}

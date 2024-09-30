package brevoservice

import brevorequest "gitlab.com/eliotandelon/brevo-email/infrastructure/brevo_request"

type BrevoMailClient struct {
	brevoApiKey      string
	mailTemplatePath string
	urlBase          string
}

func NewBrevoMailClient(brevoApiKey, mailTemplatePath, urlBase string) *BrevoMailClient {
	return &BrevoMailClient{
		brevoApiKey:      brevoApiKey,
		mailTemplatePath: mailTemplatePath,
		urlBase:          urlBase,
	}
}

func getHeaders(apiKey string) []brevorequest.HeaderRequest {
	return []brevorequest.HeaderRequest{
		{Key: "api-key", Value: apiKey},
	}
}

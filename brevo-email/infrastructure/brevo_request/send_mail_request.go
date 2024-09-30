package brevorequest

type BrevoRequestSendMail struct {
	Subject     string                 `json:"subject"`
	HtmlContent string                 `json:"htmlContent"`
	Sender      DataBrevoRequestMail   `json:"sender"`
	To          []DataBrevoRequestMail `json:"to"`
}

type DataBrevoRequestMail struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

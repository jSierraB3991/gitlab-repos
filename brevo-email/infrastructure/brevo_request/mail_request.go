package brevorequest

type MailRequest struct {
	Subject        string
	HtmlName       string
	ToDataMail     DataMail
	SenderDataMail DataMail
}

type DataMail struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

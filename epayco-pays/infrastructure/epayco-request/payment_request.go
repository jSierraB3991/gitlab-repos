package epaycorequest

type PaymentRequest struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"i"`
	ID         int    `json:"id"`
	TypeSell   string `json:"typeSell"`

	Amount             string `json:"amount"`
	Quantity           string `json:"quantity"`
	Currency           string `json:"currency"`
	Description        string `json:"description"`
	Title              string `json:"title"`
	Country            string `json:"country"`
	Test               string `json:"test"`
	UrlResponse        string `json:"urlResponse"`
	UrlConfirmation    string `json:"urlConfirmation"`
	NameBilling        string `json:"name_billing"`
	AddressBilling     string `json:"address_billing"`
	TypeDocBilling     string `json:"type_doc_billing"`
	MobilephoneBilling string `json:"mobilephone_billing"`
	NumberDocBilling   string `json:"number_doc_billing"`
}

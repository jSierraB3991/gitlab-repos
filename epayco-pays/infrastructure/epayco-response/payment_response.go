package epaycoresponse

import "time"

type PaymentDataResponse struct {
	Success       bool            `json:"success"`
	TitleResponse string          `json:"titleResponse"`
	TextResponse  string          `json:"textResponse"`
	LastAction    string          `json:"lastAction"`
	Data          PaymentResponse `json:"data"`
}

type PaymentResponse struct {
	Date            string     `json:"date"`
	State           int        `json:"state"`
	TxtCode         string     `json:"txtCode"`
	ClientID        uint       `json:"clientId"`
	OnePayment      int        `json:"onePayment"`
	Quantity        uint       `json:"quantity"`
	BaseTax         float64    `json:"baseTax"`
	Description     string     `json:"description"`
	Title           string     `json:"title"`
	Currency        string     `json:"currency"`
	TypeSell        string     `json:"typeSell"`
	URLConfirmation string     `json:"urlConfirmation"`
	URLResponse     string     `json:"urlResponse"`
	Tax             float64    `json:"tax"`
	Amount          uint       `json:"amount"`
	InvoiceNumber   string     `json:"invoceNumber"`
	ExpirationDate  *time.Time `json:"expirationDate"`
	RouteQr         string     `json:"routeQr"`
	RouteLink       string     `json:"routeLink"`
	ID              uint       `json:"id"`
}

type InvalidRequestPayResponse struct {
	TitleResponse   string `json:"titleResponse"`
	MessageResponse string `json:"textResponse"`
}

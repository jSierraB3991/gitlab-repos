package epaycorequest

type ClientRequest struct {
	CardNumber       string `json:"card[number]"`
	CardExpMonth     string `json:"card[exp_month]"`
	CardExpYear      string `json:"card[exp_year]"`
	CardCVC          string `json:"card[cvc]"`
	Name             string `json:"name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Invoice          string `json:"invoice"`
	Description      string `json:"description"`
	Value            string `json:"value"`
	Currency         string `json:"currency"`
	Dues             string `json:"dues"`
	EpaycoPublicKey  string `json:"public_key"`
	EpaycoPrivateKey string `json:"private_key"`
	CustomerID       string `json:"cust_id_cliente"`
}

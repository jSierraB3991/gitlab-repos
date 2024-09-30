package epaycoresponse

type PaginationData struct {
	TotalCount int `json:"totalCount"`
	Limit      int `json:"limit"`
	Page       int `json:"page"`
}

type DataTransactionListResponse struct {
	ReferencePayco      int     `json:"referencePayco"`
	ReferenceClient     string  `json:"referenceClient"`
	TransactionDate     string  `json:"transactionDate"`
	Description         string  `json:"description"`
	PaymentMethod       string  `json:"paymentMethod"`
	Amount              float64 `json:"amount"`
	Status              string  `json:"status"`
	Test                bool    `json:"test"`
	Currency            string  `json:"currency"`
	TransactionDateTime string  `json:"transactionDateTime"`
	Iva                 float64 `json:"iva"`
	Bank                string  `json:"bank"`
	Card                string  `json:"card"`
	Receipt             string  `json:"receipt"`
	Authorization       string  `json:"authorization"`
	Response            string  `json:"response"`
	Trmdia              float64 `json:"trmdia"`
	DocType             string  `json:"docType"`
	Document            string  `json:"document"`
	Names               string  `json:"names"`
	Lastnames           string  `json:"lastnames"`
	CicloPse            *string `json:"cicloPse"`
}

type ErrorsDataResponse struct {
	CodeError uint   `json:"codError"`
	Message   string `json:"errorMessage"`
}

type TransactionListResponse struct {
	Pagination PaginationData                `json:"pagination"`
	Data       []DataTransactionListResponse `json:"data"`
	/*Aggregations struct {
		Status []struct {
			Key      string `json:"key"`
			DocCount int    `json:"doc_count"`
		} `json:"status"`
		TransactionType struct {
			Produccion struct {
				DocCount int `json:"doc_count"`
			} `json:"produccion"`
			Pruebas struct {
				DocCount int `json:"doc_count"`
			} `json:"pruebas"`
		} `json:"transactionType"`
		TransactionFranchises map[string]struct {
			DocCount int `json:"doc_count"`
		} `json:"transactionFranchises"`
		TransactionStatus map[string]struct {
			DocCount int `json:"doc_count"`
		} `json:"transactionStatus"`
	} `json:"aggregations"`*/
	TotalErrors uint                 `json:"totalerrors"`
	Errors      []ErrorsDataResponse `json:"errors"`
}

type EpaycoTransactionResponse struct {
	Success       bool                    `json:"success"`
	TitleResponse string                  `json:"title_response"`
	TextResponse  string                  `json:"text_response"`
	LastAction    string                  `json:"last_action"`
	Data          TransactionListResponse `json:"data"`
}

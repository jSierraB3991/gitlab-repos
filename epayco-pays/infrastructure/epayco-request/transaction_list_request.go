package epaycorequest

type TransactionListRequest struct {
	Filter FilterData `json:"filter"`
}

type FilterData struct {
	ReferenceClient string `json:"referenceClient"`
}

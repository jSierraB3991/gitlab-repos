package infrastructure_test

import (
	"testing"

	epaycorequest "gitlab.com/eliotandelon/epayco-pays/infrastructure/epayco-request"
)

func TestValidateTrasactionByInvoceNumberOrReferenceClient(t *testing.T) {
	testCases := []struct {
		referenceClient string
		LengthValueData int
	}{
		{referenceClient: "149442366e228aa64383", LengthValueData: 1},
		{referenceClient: "38346aa822e663244941", LengthValueData: 0},
		{referenceClient: "", LengthValueData: 2}, //Replace 2, with all number of transactions
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.referenceClient, func(t *testing.T) {
			t.Parallel()
			response, err := s.TransactionList(epaycorequest.TransactionListRequest{
				Filter: epaycorequest.FilterData{
					ReferenceClient: tc.referenceClient,
				},
			}, token)
			if err != nil {
				t.Errorf("Transaction List Error %v", err)
			}

			if len(response) != tc.LengthValueData {
				t.Errorf("Error, I Want %v, Got %v", tc.LengthValueData, len(response))
			}
		})
	}
}

package infrastructure_test

import (
	"log"
	"os"
	"testing"

	epaycorequest "gitlab.com/eliotandelon/epayco-pays/infrastructure/epayco-request"
)

func TestPaymentEpayco(t *testing.T) {
	testCases := []struct {
		Name        string
		req         epaycorequest.PaymentRequest
		ExpectError error
	}{
		{
			req: epaycorequest.PaymentRequest{
				Amount:             "5",
				Quantity:           "10",
				Currency:           "USD",
				Description:        "Plan anuela of Test of testing",
				Country:            "CO",
				Test:               "true",
				Title:              "Plan anual",
				UrlResponse:        os.Getenv("BACK_URL") + "/public/epayco/confirmate",
				UrlConfirmation:    os.Getenv("FRONT_URL") + "/dashboard/pays",
				NameBilling:        "John Doe",
				AddressBilling:     "Kra 14a N2 a 19",
				TypeDocBilling:     "CC",
				MobilephoneBilling: "3005123535",
				NumberDocBilling:   "93072221280",
			},
			ExpectError: nil,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			response, err := s.Pay(tc.req, token)
			if err != tc.ExpectError {
				t.Errorf("Error, I Want %v, Got %v", tc.ExpectError, err)
			}
			log.Println(response)
		})
	}
}

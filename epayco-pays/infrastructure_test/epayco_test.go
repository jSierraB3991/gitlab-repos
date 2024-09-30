package infrastructure_test

import (
	"log"
	"os"
	"testing"

	epaycoserviceinterface "gitlab.com/eliotandelon/epayco-pays/domain/epayco-service-interface"
	epaycoservice "gitlab.com/eliotandelon/epayco-pays/infrastructure/epayco-service"
)

var s epaycoserviceinterface.EpaycoServiceInterface
var token string

func TestMain(m *testing.M) {
	s = epaycoservice.NewEpaycoService("https://apify.epayco.co", "3c8d81af687b42275e982ad014d56d88", "2a7204fba66470e362402c22a20b1d4b")

	tokenModel, err := s.LoginEpayco()
	if err != nil {
		log.Fatal("GET TOKEN")
	}
	token = tokenModel.Token
	code := m.Run()
	os.Exit(code)
}

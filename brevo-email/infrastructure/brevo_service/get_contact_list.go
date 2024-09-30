package brevoservice

import brevoresponse "gitlab.com/eliotandelon/brevo-email/infrastructure/brevo_response"

func (client BrevoMailClient) GetAllList() (*brevoresponse.EmailListResponse, error) {
	var result brevoresponse.EmailListResponse
	err := client.get(client.urlBase, "/v3/contacts/lists", &result, getHeaders(client.brevoApiKey))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client BrevoMailClient) GetIdOfListByName(name string) (float64, error) {
	data, err := client.GetAllList()
	if err != nil {
		return 0, err
	}

	var result float64
	for _, v := range data.EmailListDataResponse {
		if v.Name == name {
			result = v.Id
			break
		}
	}
	return result, nil
}

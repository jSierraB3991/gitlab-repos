package brevoresponse

type EmailListDataResponse struct {
	Id   float64 `json:"id"`
	Name string  `json:"name"`
}

type EmailListResponse struct {
	Count                 float64                 `json:"count"`
	EmailListDataResponse []EmailListDataResponse `json:"lists"`
}

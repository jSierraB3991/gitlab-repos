package brevorequest

type CreateContactAtributtesRequest struct {
	LastName string `json:"APELLIDOS"`
	Name     string `json:"NOMBRE"`
	SMS      string `json:"SMS"`
	Whatsapp string `json:"WHATSAPP"`
}

type CreateContactRequest struct {
	Email         string                         `json:"email"`
	ListIds       []int64                        `json:"listIds"`
	UpdateEnabled bool                           `json:"updateEnabled"`
	Attributes    CreateContactAtributtesRequest `json:"attributes,omitempty"`
}

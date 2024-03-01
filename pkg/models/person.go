package models

type PessoaDados struct {
	CPF          string `json:"cpf"`
	IDPerson     int    `json:"idPerson"`
	Email        string `json:"email"`
	BirthCountry string `json:"birthCountry"`
	Zipcode      string `json:"zipcode"`
	City         string `json:"city"`
	State        string `json:"state"`
	StreetNumber string `json:"streetNumber"`
	Street       string `json:"street"`
	Name         string `json:"name"`
}

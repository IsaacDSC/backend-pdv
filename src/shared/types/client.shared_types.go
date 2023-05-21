package shared_types

type ClientCreateOnSet_DTO struct {
	ID           string     `json:"id"`
	EnterpriseId string     `json:"enterpriseId"`
	Name         string     `json:"name"`
	Telephone    string     `json:"telephone"`
	Email        string     `json:"email"`
	Address      AddressDTO `json:"address"`
	CreatedAt    string     `json:"createdAt"`
}

type AddressDTO struct {
	City         string `json:"city"`
	HomeNumber   string `json:"homeNumber"`
	Neighborhood string `json:"neighborhood"`
	Cep          string `json:"cep"`
	Address      string `json:"address"`
	Observation  string `json:"observation"`
}

package shared_types

type CreateEnterprise_DTO struct {
	CorporateName string `json:"corporateName"`
	FantasyName   string `json:"fantasyName"`
	Logo          string `json:"logo"`
	CPJ           string `json:"CPJ"`
	Token         string `json:"token"`
}

type OutputCreateEnterprise struct {
	EnterpriseId string `json:"enterpriseId"`
	Token        string `json:"appToken"`
}

package domain

import "github.com/savsgio/gotils/uuid"

type EnterpriseEntity struct {
	ID            string
	CorporateName string
	FantasyName   string
	Logo          string
	CPJ           string
	Status        string `default:"TRIAL"`
	Token         string
}

func (this_entity *EnterpriseEntity) Validate() (errors []string) {
	this_entity.ID = uuid.V4()
	if len(this_entity.Token) == 0 {
		errors = append(errors, "Is-Required-Token")
	}
	if len(this_entity.FantasyName) == 0 {
		errors = append(errors, "Is-Required-FantasyName")
	}
	return
}

package domain

type CategoryEntity struct {
	ID           string
	Name         string
	Description  string
	EnterpriseId string
}

func (this_entity *CategoryEntity) Validate() (errors []string) {
	if len(this_entity.EnterpriseId) == 0 {
		errors = append(errors, "EnterpriseId-Is-Required")
	}
	return
}

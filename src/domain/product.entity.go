package domain

type ProductEntity struct {
	ID           string
	Name         string
	Description  string
	Image        string
	Price        float32
	EnterpriseId string
	CategoryId   string
}

func (this_entity *ProductEntity) Validate() (errors []string) {
	if len(this_entity.EnterpriseId) == 0 {
		errors = append(errors, "EnterpriseId-Is-Required")
		return
	}
	if len(this_entity.CategoryId) == 0 {
		errors = append(errors, "CategoryId-Is-Required")
		return
	}
	if this_entity.Price < 0 {
		errors = append(errors, "Price-Is-Required-Type-Real")
		return
	}
	return
}

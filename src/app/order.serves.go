package app_services

import "backend-pdv/src/domain"

type OrderService struct {
	Entity       *domain.Order
	ProductsIDs  []string
	ClientId     string
	EnterpriseId string
	DeliveryId   string
}

func (this_service *OrderService) CreateOrder() (errors []string) {
	this_service.Entity.ValidateOrder()
	return
}

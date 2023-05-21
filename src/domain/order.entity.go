package domain

import (
	"log"

	"github.com/savsgio/gotils/uuid"
)

type Repository interface {
	createOrder(Order) error
	UpdateOrder(Order) error
}

type Order struct {
	ID            string
	Price         float32
	CustomerId    string
	AddressId     string
	MethodPayment string
	Status        string //Aceito Pedido - Preparando - Saiu Para entrega - Entregue
	Delivered     bool
	Payment       bool
	Observation   string
	CreatedAt     string
	UpdatedAt     string
	OrderItems    []OrderItems
	repository    Repository
}

type OrderItems struct {
	ProductId string
	Quantity  int
	Price     float32
}

func (this_order *Order) ValidateOrder() (errors []string) {
	this_order.Status = "PENDING"
	this_order.Delivered = false
	this_order.Payment = false
	if len(this_order.ID) == 0 {
		this_order.ID = string(uuid.V4())
	}
	if this_order.Price <= 0 {
		errors = append(errors, "Not-Found-Pricing")
	}
	if len(this_order.CustomerId) == 0 {
		errors = append(errors, "Not_Found-CustomerId")
	}
	if len(this_order.AddressId) == 0 {
		errors = append(errors, "Not-Found-AddressId")
	}
	if len(this_order.MethodPayment) == 0 {
		errors = append(errors, "Not-Found-MethodPayment")
	} // Money - Source
	if len(this_order.OrderItems) == 0 {
		errors = append(errors, "Not-Found-OrderItems")
	}
	return
}

func (this_order *Order) SaveOrder() (err error) {
	return this_order.repository.createOrder(*this_order)
}

func (this_order *Order) UpdateStatusOrder() (err error) {
	return this_order.repository.UpdateOrder(*this_order)
}

func (this_order *Order) CancelOrder() (err error) {
	return this_order.repository.UpdateOrder(*this_order)
}

func Test() {
	test := Order{}
	errors := test.ValidateOrder()
	if errors != nil {
		log.Fatalf("ERROR")
	}
	test.SaveOrder()
}

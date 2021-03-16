package interfaces

import "Ordering/core"

type IOrdersRepository interface {
	CreateOrder(order core.Order) bool
	GetAllOrders() []*core.Order
	GetOrderById(order int) *core.Order
	DeleteOrder(order core.Order) bool
	UpdateOrder(order core.Order) bool
}

package main

import (
	"Ordering/core"
	"Ordering/core/interfaces"
	"github.com/gin-gonic/gin"
	"strconv"
)

var jsonContentType = "application/json; charset=utf-8"
var orderRepository interfaces.IOrdersRepository

func RouteOrders(router *gin.Engine)  {
	router.GET("/orders", GetAllOrders)
	router.GET("/orders/:id", GetOrderById)
	router.POST("/orders", CreateOrder)
	router.DELETE("/orders/:id", DeleteOrder)
	router.PUT("/orders/:id", UpdateOrder)
}

func GetAllOrders(context *gin.Context)  {
	orders := orderRepository.GetAllOrders()
	context.JSON(200, orders)
}

func GetOrderById(context *gin.Context)  {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil || id < 1 {
		context.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	order := orderRepository.GetOrderById(id)
	context.JSON(200, order)
}

func CreateOrder(context *gin.Context)  {
	order := &core.Order{}
	err := context.BindJSON(order)
	if err != nil {
		context.Data(400, jsonContentType, []byte("Fill all fields"))
		return
	}
	if orderRepository.CreateOrder(*order) {
		context.Data(200, jsonContentType, []byte("Created order \n"))
	}
	context.Data(500, jsonContentType, []byte("Failed to create order"))
}

func DeleteOrder(context *gin.Context)  {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil || id < 1 {
		context.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	order := orderRepository.GetOrderById(id)
	if order == nil {
		context.Data(400, jsonContentType, []byte("No such order with id"))
		return
	}
	if orderRepository.DeleteOrder(*order) {
		context.Data(200, jsonContentType, []byte("Deleted order"))
		return
	}
	context.Data(500, jsonContentType, []byte("Failed to delete order"))
}

func UpdateOrder(context *gin.Context)  {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil || id < 1 {
		context.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	model := orderRepository.GetOrderById(id)
	order := &core.Order{}
	err = context.BindJSON(order)
	if err != nil {
		context.Data(400, jsonContentType, []byte("Fill all fields"))
		return
	}
	order.Id = id
	updateValues(model, order)
	if orderRepository.UpdateOrder(*order) {
		context.Data(200, jsonContentType, []byte("Updated order"))
		return
	}
	context.Data(500, jsonContentType, []byte("Failed to update order"))
}

func updateValues(order *core.Order, updateOrder *core.Order)  {
	if updateOrder.Sum > 0 {
		order.Sum = updateOrder.Sum
	}

	if len(updateOrder.Destination) > 0 {
		order.Destination = updateOrder.Destination
	}


}

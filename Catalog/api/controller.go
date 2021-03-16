package main

import (
	"Catalog/core"
	"Catalog/core/interfaces"
	"github.com/gin-gonic/gin"
	"strconv"
)
var jsonContentType = "application/json; charset=utf-8"
var sockRepository interfaces.ISocksRepository

func RouteSocks(router *gin.Engine)  {
	router.GET("/socks", GetAllSocks)
	router.GET("/socks/:id", GetSockById)
	router.POST("/socks", CreateSock)
	router.DELETE("/socks/:id", DeleteSock)
	router.PUT("/socks/:id", UpdateSock)
}

func GetAllSocks(context *gin.Context)  {
	socks := sockRepository.GetAllSocks()
	context.JSON(200, socks)
}

func GetSockById(context *gin.Context)  {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil || id < 1 {
		context.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	sock := sockRepository.GetSockById(id)
	context.JSON(200, sock)
}

func CreateSock(context *gin.Context)  {
	sock := &core.Sock{}
	err := context.BindJSON(sock)
	if err != nil {
		context.Data(400, jsonContentType, []byte("Fill all fields"))
		return
	}
	if sockRepository.CreateSock(*sock) {
		context.Data(200, jsonContentType, []byte("Created sock \n"))
	}
	context.Data(500, jsonContentType, []byte("Failed to create sock"))
}

func DeleteSock(context *gin.Context)  {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil || id < 1 {
		context.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	sock := sockRepository.GetSockById(id)
	if sock == nil {
		context.Data(400, jsonContentType, []byte("No such sock with id"))
		return
	}
	if sockRepository.DeleteSock(*sock) {
		context.Data(200, jsonContentType, []byte("Deleted sock"))
		return
	}
	context.Data(500, jsonContentType, []byte("Failed to delete sock"))
}

func UpdateSock(context *gin.Context)  {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil || id < 1 {
		context.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	model := sockRepository.GetSockById(id)
	sock := &core.Sock{}
	err = context.BindJSON(sock)
	if err != nil {
		context.Data(400, jsonContentType, []byte("Fill all fields"))
		return
	}
	sock.Id = id
	updateValues(model, sock)
	if sockRepository.UpdateSock(*sock) {
		context.Data(200, jsonContentType, []byte("Updated sock"))
		return
	}
	context.Data(500, jsonContentType, []byte("Failed to update sock"))
}

func updateValues(sock *core.Sock, updateSock *core.Sock)  {
	if len(updateSock.Name) > 0 {
		sock.Name = updateSock.Name
	}
	if updateSock.Price > 0 {
		sock.Price = updateSock.Price
	}
	if updateSock.Quantity > 0 {
		sock.Quantity = updateSock.Quantity
	}
}

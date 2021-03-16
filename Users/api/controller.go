package main

import (
	"Users/core"
	"Users/core/interfaces"
	"github.com/gin-gonic/gin"
	"strconv"
)

var jsonContentType = "application/json; charset=utf-8"
var userRepository interfaces.IUsersRepository

func RouteUsers(router *gin.Engine)  {
	router.GET("/users", GetAllUsers)
	router.GET("/users/:id", GetUserById)
	router.POST("/users", CreateUser)
	router.DELETE("/users/:id", DeleteUser)
	router.PUT("/users/:id", UpdateUser)
}

func GetAllUsers(context *gin.Context)  {
	users := userRepository.GetAllUsers()
	context.JSON(200, users)
}

func GetUserById(context *gin.Context){
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil || id < 1 {
		context.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	user := userRepository.GetUserById(id)
	context.JSON(200, user)
}

func CreateUser(context *gin.Context) {
	user := &core.User{}
	err := context.BindJSON(user)
	if err != nil {
		context.Data(400, jsonContentType, []byte("Fill all fields"))
		return
	}
	if userRepository.CreateUser(*user) {
		context.Data(200, jsonContentType, []byte("Created user \n"))
	}
	context.Data(500, jsonContentType, []byte("Failed to create user"))
}

func DeleteUser(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil || id < 1 {
		context.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	user := userRepository.GetUserById(id)
	if user == nil {
		context.Data(400, jsonContentType, []byte("No such user with id"))
		return
	}
	if userRepository.DeleteUser(*user) {
		context.Data(200, jsonContentType, []byte("Deleted user"))
		return
	}
	context.Data(500, jsonContentType, []byte("Failed to delete user"))
}

func UpdateUser(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil || id < 1 {
		context.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	model := userRepository.GetUserById(id)
	user := &core.User{}
	err = context.BindJSON(user)
	if err != nil {
		context.Data(400, jsonContentType, []byte("Fill all fields"))
		return
	}
	user.Id = id
	updateValues(model, user)
	if userRepository.UpdateUser(*user) {
		context.Data(200, jsonContentType, []byte("Updated user"))
		return
	}
	context.Data(500, jsonContentType, []byte("Failed to update user"))
}

func updateValues(user *core.User, updateUser *core.User)  {
	if len(updateUser.Name) > 0 {
		user.Name = updateUser.Name
	}
	if len(updateUser.Surname) > 0 {
		user.Surname = updateUser.Surname
	}
	if len(updateUser.Username) > 0 {
		user.Username = updateUser.Username
	}
}
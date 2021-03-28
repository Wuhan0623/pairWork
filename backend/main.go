package main

import (
	_ "backend/database"
	"backend/user"
	"github.com/gin-gonic/gin"
)


func main() {

	r := gin.Default()
	r.GET("/user/login", user.UserLogin)
	r.GET("/user/register", user.UserRegister)
	r.GET("/search/list", user.GetThesisList)
	//r.POST("/user", user.HandleList)
	r.GET("/")
	r.Run(":8080")
}


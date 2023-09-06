package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/Config"
	"go-web/Router"
)

func register(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}
func login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func main() {
	Config.Init()
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recover:%v\n", err)
		}
	}()

	//controller.StudentInfoLoad()
	e := gin.Default()
	e.GET("/register.html", register)
	e.GET("/", login)
	Router.InitRouter(e)
}

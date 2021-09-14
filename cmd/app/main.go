package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lazy-scaffold-api/internal/domain"
	"log"
)

// @title lazy-scaffold-api
// @version 1.0
// @in header
// @name Authorization
// @host localhost:8088
func main() {
	app := gin.New()

	//注册业务模块
	for _, option := range domain.Registry() {
		group := app.Group(option.Name)
		{
			for _, child := range option.ChildList {
				switch child.Method {
				case "GET":
					group.GET(child.Route, child.Handles...)
				case "POST":
					group.POST(child.Route, child.Handles...)
				case "PUT":
					group.PUT(child.Route, child.Handles...)
				case "DELETE":
					group.DELETE(child.Route, child.Handles...)
				}
			}
		}
	}

	// 启动服务
	if err := app.Run(fmt.Sprintf(":%v", 8080)); err != nil {
		log.Println("app run validate_err", err)
	}
}
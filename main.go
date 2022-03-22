package main

import (
	"fmt"
	"log"

	_ "github.com/NICEXAI/go-program-tuning"
	"github.com/gin-gonic/gin"

	// lazy replace:name>lazy-scaffold-api range:2
	"lazy-scaffold-api/internal/config"
	"lazy-scaffold-api/internal/domain"
)

// lazy replace:name>lazy-scaffold-api range:1
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
	if err := app.Run(fmt.Sprintf(":%v", config.Info.Server.Port)); err != nil {
		log.Println("app run validate_err", err)
	}
}

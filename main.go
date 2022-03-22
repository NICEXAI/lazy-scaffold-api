package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/NICEXAI/go-program-tuning"
	gracefulExit "github.com/NICEXAI/graceful-exit"
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

	server := &http.Server{
		Addr:         ":" + strconv.Itoa(config.Info.Server.Port),
		Handler:      app,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				os.Exit(1)
			}
		}
	}()

	// 服务优雅退出
	graceful := gracefulExit.NewGracefulExit()
	graceful.RegistryHandle("exit", func() {
		// 停止server服务
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatalf("server down error: %s", err.Error())
		}

		log.Println("service stop successfully")
	})

	graceful.Capture()
}

package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	gracefulExit "github.com/NICEXAI/graceful-exit"
	"github.com/gin-gonic/gin"

	// lazy replace:name>lazy-scaffold-api range:2
	"lazy-scaffold-api/internal/config"
	"lazy-scaffold-api/internal/domain"
)

func Run() {
	var app *gin.Engine

	if config.Config.Application.Mode == "debug" {
		app = gin.Default()
	} else {
		app = gin.New()
	}

	// 注册业务模块
	domainList := domain.Registry()
	for i := range domainList {
		option := domainList[i]
		group := app.Group(option.Name)
		{
			for j := range option.ChildList {
				child := option.ChildList[j]
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
		Addr:         ":" + strconv.Itoa(config.Config.Application.Port),
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

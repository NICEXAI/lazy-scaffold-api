package docs

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	// lazy replace:name>lazy-scaffold-api range:2
	_ "lazy-scaffold-api/docs"
	"lazy-scaffold-api/internal/domain/common"
)

// Setup 注册文档模块
func Setup() common.ModuleOption {
	return common.ModuleOption{
		Name: "docs",
		ChildList: []common.ModuleChild{
			{
				Route:   "/*any",
				Method:  "GET",
				Handles: []gin.HandlerFunc{ginSwagger.WrapHandler(swaggerFiles.Handler)},
			},
		},
	}
}

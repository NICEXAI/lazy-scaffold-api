package tool

import (
	"github.com/gin-gonic/gin"

	// lazy replace:name>lazy-scaffold-api range:1
	"lazy-scaffold-api/internal/domain/common"
)

// Setup 注册业务模块
func Setup() common.ModuleOption {
	return common.ModuleOption{
		Name: "tool",
		ChildList: []common.ModuleChild{
			{
				Route:   "/ping",
				Method:  "GET",
				Handles: []gin.HandlerFunc{Ping},
			},
			{
				Route:   "/real-ip",
				Method:  "GET",
				Handles: []gin.HandlerFunc{GetRealIP},
			},
		},
	}
}

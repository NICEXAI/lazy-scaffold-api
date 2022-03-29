package domain

import (

	// @Lazy var:name>lazy-scaffold-api scope:3
	"lazy-scaffold-api/internal/domain/common"
	"lazy-scaffold-api/internal/domain/docs"
	"lazy-scaffold-api/internal/domain/tool"
)

// Registry 挂载业务模块
func Registry() common.ModuleOptionList {
	return common.ModuleOptionList{
		docs.Setup(),
		tool.Setup(),
	}
}

package main

import (
	_ "github.com/NICEXAI/go-program-tuning"

	// @Lazy var:name>lazy-scaffold-api scope:1
	"lazy-scaffold-api/cmd"
)

// @Lazy var:name>lazy-scaffold-api scope:1
// @title lazy-scaffold-api
// @version 1.0
// @in header
// @name Authorization
// @host localhost:8088
func main() {
	cmd.Execute()
}

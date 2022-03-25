package main

import (
	_ "github.com/NICEXAI/go-program-tuning"

	// lazy replace:name>lazy-scaffold-api range:1
	"lazy-scaffold-api/cmd"
)

// lazy replace:name>lazy-scaffold-api range:1
// @title lazy-scaffold-api
// @version 1.0
// @in header
// @name Authorization
// @host localhost:8088
func main() {
	cmd.Execute()
}

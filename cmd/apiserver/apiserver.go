package main

import (
	"os"

	"github.com/luohu1/gin-apiserver/internal/apiserver/cmd"
)

func main() {
	command := cmd.NewCmdAPIServer()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

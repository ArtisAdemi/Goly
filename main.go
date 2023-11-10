package main

import (
	"github.com/artisademi/goly/model"
	"github.com/artisademi/goly/server"
)

func main() {
	model.Setup()

	server.SetupAndListen()
}

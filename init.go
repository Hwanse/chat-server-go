package main

import (
	"github.com/Hwanse/chat-server-go/repository"
	"github.com/Hwanse/chat-server-go/web"
)

func InitializeServer() {
	repository.Connect()
	web.InitRouter()
}

package main

import "github.com/Hwanse/chat-server-go/repository"

func InitializeServer() {
	repository.Connect()
}

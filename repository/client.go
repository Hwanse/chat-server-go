package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Client struct {
	*gorm.DB
}

var Connector *Client

func Connect() {
	host := "chat_server:chat_server@tcp(127.0.0.1:3306)/chat_server?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(host), &gorm.Config{})
	if err != nil {
		panic("fail connect to DB")
	}

	Connector = &Client{db}
}

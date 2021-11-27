module github.com/Hwanse/chat-server-go

go 1.16

replace github.com/Hwanse/chagot-server-go/repository => ./repository

require (
	github.com/jinzhu/now v1.1.3 // indirect
	gorm.io/driver/mysql v1.2.0
	gorm.io/gorm v1.22.3
)

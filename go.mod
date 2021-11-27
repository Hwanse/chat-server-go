module github.com/Hwanse/chat-server-go

go 1.16

//replace github.com/Hwanse/chagot-server-go/repository => ./repository
replace github.com/Hwanse/chagot-server-go/web => ./web

require (
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/jinzhu/now v1.1.3 // indirect
	gorm.io/driver/mysql v1.2.0
	gorm.io/gorm v1.22.3
)

package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InitRouter() {
	fmt.Println("initialize router handlers")
	router := mux.NewRouter()

	router.HandleFunc("/ws", websocketHandle)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe unexpected error occurred", err)
		return
	}
}

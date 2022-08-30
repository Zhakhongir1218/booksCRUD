package configuration

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func CreateServer(router *httprouter.Router) {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Server is running")
	log.Fatalln(server.Serve(listener))
}

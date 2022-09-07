package configuration

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func CreateServer(router *httprouter.Router) {
	_, address := LoadConfig(".", "address")
	_, networkType := LoadConfig(".", "network.type")
	listener, err := net.Listen(networkType, address)
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

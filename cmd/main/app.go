package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	router := httprouter.New()
	router.GET("/:id", IndexHandler)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatalln(server.Serve(listener))
}

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	w.Write([]byte(id))
}

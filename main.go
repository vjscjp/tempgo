package main

import (
	"fmt"
	"net/http"

	"github.com/CiscoCloud/shipped-utils/server"
	"github.com/gorilla/mux"
)

const (
	Port = "8888"
)

func main() {
	muxRouter := mux.NewRouter()

	server.InitRoutes(muxRouter)
	fmt.Println("Listening at port:", Port)
	http.ListenAndServe(":"+Port, muxRouter)
}

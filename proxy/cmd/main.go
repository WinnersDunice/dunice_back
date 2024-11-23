package main

import (
	"flag"
	"log"

	"github.com/WinnersDunice/dunice_back/proxy/api/handler"
	"github.com/WinnersDunice/dunice_back/proxy/pkg/service"
)

func main() {
	port := flag.String("port", "7222", "port of the server")
	flag.Parse()
	srv := new(service.Server)
	//create handler
	handlers := handler.NewHandler(srv)
	//run the application
	if err := srv.Run(*port, handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running http server: %s", err.Error())
	}
}

package main

import (
	"beard-catalog/routes"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 80, "The port to listen on")
	flag.Parse()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081"},
		AllowCredentials: true,
	})

	rr := routes.Register()

	handler := c.Handler(rr)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}

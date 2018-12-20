package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rezaindrag/restapi/router"
)

func main() {
	r := router.New()

	fmt.Println("Server running")
	log.Fatal(http.ListenAndServe(":8000", r))
}

package main

import (
	"api/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("API is running...")

	r := http.NewServeMux()

	fmt.Printf("Port %d,", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":%d",config.Port), r))

}

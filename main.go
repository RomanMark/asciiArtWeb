package main

import (
	"ascii-art-web/brakers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fileServer)
	http.HandleFunc("/ascii-art", brakers.FormHandler)
	fmt.Printf("Starting server at post: 8070\nhttp://localhost:8070/\n")
	log.Fatal(http.ListenAndServe(":8070", nil))
}

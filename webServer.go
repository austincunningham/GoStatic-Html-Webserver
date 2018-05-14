package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Hello world from golang\n")

	//outputs
	fmt.Printf("Server listing on http://localhost:8080\n")
	fmt.Printf("CTRL C to exit\n")

	// serves files from the public directory
	http.Handle("/", http.FileServer(http.Dir("./public")))
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	panic(err)
	// }

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, r.URL.Path[1:])
	// })
	log.Fatal(http.ListenAndServe(":8080", nil))
}

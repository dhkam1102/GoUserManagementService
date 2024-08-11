// cmd/main/main.go
package main

import (
	// "fmt"
	"log"
	"net/http"
	"user-management-service/cmd/api"
)

func main() {

	http.HandleFunc("/register", api.RegisterHandler)
	http.HandleFunc("/login", api.LoginHandler)
	http.HandleFunc("/update", api.UpdateUserHandler)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Hello, User Management Service!")
	// })

	log.Println("Server is running on port 8080")
	var err error = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

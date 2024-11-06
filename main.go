package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/train-do/Template-pada-golang-web-dan-Middleware/database"
	"github.com/train-do/Template-pada-golang-web-dan-Middleware/handler"
	"github.com/train-do/Template-pada-golang-web-dan-Middleware/middleware"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	userMux := http.NewServeMux()
	userMux.HandleFunc("/register", handler.Register(db))
	userMux.HandleFunc("/login", handler.Login(db))
	userMux.HandleFunc("/users", handler.GetUsers(db))

	todoMux := http.NewServeMux()
	todoMux.HandleFunc("/all", handler.GetTodo(db))
	todoMux.HandleFunc("/create", handler.CreateTodo(db))

	auth := middleware.Authentication(db, todoMux)
	serverMux := http.NewServeMux()
	serverMux.Handle("/", userMux)
	serverMux.Handle("/todo/", http.StripPrefix("/todo", auth))
	fmt.Println("server started on port 8080")
	http.ListenAndServe(":8080", serverMux)
}

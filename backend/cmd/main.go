package main

import (
	"fmt"
	"log"
	"net/http"

	"zed-demo/pkg/handlers"
	"zed-demo/pkg/storage"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=user password=password dbname=tododb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	storage.Migrate(db)

	r := mux.NewRouter()

	h := handlers.New(db)

	r.HandleFunc("/todos", h.CreateTodo).Methods("POST")
	r.HandleFunc("/todos", h.GetTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", h.GetTodo).Methods("GET")
	r.HandleFunc("/todos/{id}", h.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", h.DeleteTodo).Methods("DELETE")

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

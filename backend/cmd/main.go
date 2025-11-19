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
	dsn := "host=db user=user password=password dbname=tododb port=5432 sslmode=disable"
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

	// CORS middleware
	// CORS middleware function
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(r)))
}

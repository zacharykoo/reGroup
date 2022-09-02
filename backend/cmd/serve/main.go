package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zacharykoo/reGroup/backend/pkg/database"
	"github.com/zacharykoo/reGroup/backend/pkg/repository/lite"
	"github.com/zacharykoo/reGroup/backend/pkg/service"
)

func main() {
	fmt.Println("Hello world")

	db, err := database.ConnectSQLite()
	if err != nil {
		fmt.Printf("unable to connect to database: %v\n", err)
		return
	}

	err = database.MigrateTables(db)
	if err != nil {
		fmt.Printf("unable to migrate tables: %v", err)
		return
	}

	customerRepository := lite.GetUserRepository(db)
	customerService := service.GetUserService(&customerRepository)
	r := mux.NewRouter()

	r.Handle("/api/customer", customerService.Get()).Methods("GET")
	r.Handle("/api/customer", customerService.Create()).Methods("POST")
	r.Handle("/api/customer", customerService.Edit()).Methods("PUT")

	r.Use(corsMiddleware)

	fmt.Println("Serving on :8081...")
	http.ListenAndServe(":8081", r)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zacharykoo/reGroup/backend/pkg/database"
	"github.com/zacharykoo/reGroup/backend/pkg/repository/lite"
	"github.com/zacharykoo/reGroup/backend/pkg/service/services"
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

	userRepository := lite.GetUserRepository(db)
	userService := services.GetUserService(&userRepository)
	r := mux.NewRouter()

	r.Handle("/api/user", userService.Get()).Methods("GET")
	r.Handle("/api/user", userService.Create()).Methods("POST")
	r.Handle("/api/user", userService.Edit()).Methods("PUT")

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

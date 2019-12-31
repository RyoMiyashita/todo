package main

import (
	"miyazi.com/todo-go/loopback"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
)

const PORT = "9000"

func main() {
	router := makeRouter()
	router.Use(settingMiddleware)
	if err := http.ListenAndServe(":"+PORT, router); err != nil {
		log.Fatalf("err: %v", err)
	}
}

func makeRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/loopback/{message}", loopback.Loopback)
	router.HandleFunc("/loopback/", loopback.Loopback)
	return router
}

func settingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

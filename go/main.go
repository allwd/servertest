package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func test(w http.ResponseWriter, r *http.Request) {
	message := "{ 'status': 'Success.' }"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}


func middlewareAuthorize(next http.Handler) http.Handler {
	key := "Test"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != key {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("{ 'err': 'Unauthorized.' }"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	port := "3000"
	r := mux.NewRouter()

	r.HandleFunc("/test", test)
	r.Use(middlewareAuthorize)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	for k, v := range handlers {
		r.PathPrefix(k).Handler(v)
	}

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

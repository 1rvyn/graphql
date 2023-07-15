package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Login(router *mux.Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Successful login"))
	}
}

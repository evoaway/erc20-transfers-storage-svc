package requests

import (
	"github.com/go-chi/chi"
	"net/http"
)

func NewGetAddress(r *http.Request) string {
	address := chi.URLParam(r, "address")
	return address
}

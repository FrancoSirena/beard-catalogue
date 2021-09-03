package routes

import (
	"github.com/gorilla/mux"
)

func Register() *mux.Router {
	routes := mux.NewRouter().StrictSlash(true)
	routes.HandleFunc("/api/v1/brands/{brand_id}/", ResolveGetBrand).Methods("GET")
	routes.HandleFunc("/api/v1/brands/", ResolveGetBrands).Methods("GET")
	return routes
}

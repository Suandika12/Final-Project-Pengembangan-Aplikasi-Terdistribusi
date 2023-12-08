package routes

import (
	"github.com/Suandika12/go_Pemesanan/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterPemesanansRoutes = func(router *mux.Router) {
	router.HandleFunc("/pemesanan/",
		controllers.CreatePemesanan).Methods("POST")
	router.HandleFunc("/pemesanan/", controllers.GetPemesanan).Methods("GET")
	router.HandleFunc("/pemesanan/{pemesananId}",
		controllers.GetPemesananById).Methods("GET")
	router.HandleFunc("/pemesanan/{pemesananId}",
		controllers.UpdatePemesanan).Methods("PUT")
	router.HandleFunc("/pemesanan/{pemesananId}",
		controllers.DeletePemesanan).Methods("DELETE")
}

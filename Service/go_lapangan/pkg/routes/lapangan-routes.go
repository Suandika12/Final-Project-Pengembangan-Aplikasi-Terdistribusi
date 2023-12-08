package routes

import (
	"github.com/Suandika12/go_lapangan/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterLapangansRoutes = func(router *mux.Router) {
	router.HandleFunc("/lapangan/",
		controllers.CreateLapangan).Methods("POST")
	router.HandleFunc("/lapangan/", controllers.GetLapangan).Methods("GET")
	router.HandleFunc("/lapangan/{lapanganId}",
		controllers.GetLapanganById).Methods("GET")
	router.HandleFunc("/lapangan/{lapanganId}",
		controllers.UpdateLapangan).Methods("PUT")
	router.HandleFunc("/lapangan/{lapanganId}",
		controllers.DeleteLapangan).Methods("DELETE")
}

package routes

import (
	"github.com/Suandika12/go_bookingLap/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookingLapsRoutes = func(router *mux.Router) {
	router.HandleFunc("/pemesanan/",
		controllers.CreateBookingLap).Methods("POST")
	router.HandleFunc("/pemesanan/", controllers.GetBookingLap).Methods("GET")
	router.HandleFunc("/pemesanan/{pemesananId}",
		controllers.GetBookingLapById).Methods("GET")
	router.HandleFunc("/pemesanan/{pemesananId}",
		controllers.UpdateBookingLap).Methods("PUT")
	router.HandleFunc("/pemesanan/{pemesananId}",
		controllers.DeleteBookingLap).Methods("DELETE")
}

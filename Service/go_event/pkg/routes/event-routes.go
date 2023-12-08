package routes

import (
	"github.com/Suandika12/go_event/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterEventsRoutes = func(router *mux.Router) {
	router.HandleFunc("/event/",
		controllers.CreateEvent).Methods("POST")
	router.HandleFunc("/event/", controllers.GetEvent).Methods("GET")
	router.HandleFunc("/event/{eventId}",
		controllers.GetEventById).Methods("GET")
	router.HandleFunc("/event/{eventId}",
		controllers.UpdateEvent).Methods("PUT")
	router.HandleFunc("/event/{eventId}",
		controllers.DeleteEvent).Methods("DELETE")
}

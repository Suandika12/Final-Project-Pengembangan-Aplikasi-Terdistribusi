package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Suandika12/go_event/pkg/models"
	"github.com/Suandika12/go_event/pkg/utils"
	"github.com/gorilla/mux"
)

var NewEvent models.Event

func GetEvent(w http.ResponseWriter, r *http.Request) {
	newEvents := models.GetAllEvents()
	res, _ := json.Marshal(newEvents)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetEventById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["eventId"]
	Title, err := strconv.ParseInt(eventId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	eventDetails, _ := models.GetEventbyId(Title)
	res, _ := json.Marshal(eventDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	CreateEvent := &models.Event{}
	utils.ParseBody(r, CreateEvent)
	b := CreateEvent.CreateEvent()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["eventId"]
	Title, err := strconv.ParseInt(eventId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	event := models.DeleteEvent(Title)
	res, _ := json.Marshal(event)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var updateEvent = &models.Event{}
	utils.ParseBody(r, updateEvent)
	vars := mux.Vars(r)
	eventId := vars["eventId"]
	Title, err := strconv.ParseInt(eventId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	eventDetails, db := models.GetEventbyId(Title)
	if updateEvent.Title != "" {
		eventDetails.Title = updateEvent.Title
	}
	if updateEvent.Start_date != "" {
		eventDetails.Start_date = updateEvent.Start_date
	}
	if updateEvent.End_date != "" {
		eventDetails.End_date = updateEvent.End_date
	}
	db.Save(&eventDetails)
	res, _ := json.Marshal(eventDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

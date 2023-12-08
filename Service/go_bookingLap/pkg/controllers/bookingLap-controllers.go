package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Suandika12/go_bookingLap/pkg/models"
	"github.com/Suandika12/go_bookingLap/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBookingLap models.BookingLap

func GetBookingLap(w http.ResponseWriter, r *http.Request) {
	newBookingLaps := models.GetAllBookingLaps()
	res, _ := json.Marshal(newBookingLaps)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBookingLapById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookingLapId := vars["bookingLapId"]
	Field_name, err := strconv.ParseInt(bookingLapId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookingLapDetails, _ := models.GetBookingLapById(Field_name)
	res, _ := json.Marshal(bookingLapDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBookingLap(w http.ResponseWriter, r *http.Request) {
	CreateBookingLap := &models.BookingLap{}
	utils.ParseBody(r, CreateBookingLap)
	b := CreateBookingLap.CreateBookingLap()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBookingLap(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookingLapId := vars["bookingLapId"]
	Field_name, err := strconv.ParseInt(bookingLapId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookingLap := models.DeleteBookingLap(Field_name)
	res, _ := json.Marshal(bookingLap)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateBookingLap(w http.ResponseWriter, r *http.Request) {
	var updateBookingLap = &models.BookingLap{}
	utils.ParseBody(r, updateBookingLap)
	vars := mux.Vars(r)
	bookingLapId := vars["bookingLapId"]
	Field_name, err := strconv.ParseInt(bookingLapId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookingLapDetails, db := models.GetBookingLapById(Field_name)
	if updateBookingLap.Field_id != 0 {
		bookingLapDetails.Field_id = updateBookingLap.Field_id
	}
	if updateBookingLap.User_id != 0 {
		bookingLapDetails.User_id = updateBookingLap.User_id
	}
	if updateBookingLap.Description != "" {
		bookingLapDetails.Description = updateBookingLap.Description
	}
	if updateBookingLap.Bukti_pembayaran != "" {
		bookingLapDetails.Bukti_pembayaran = updateBookingLap.Bukti_pembayaran
	}
	if updateBookingLap.Payment_method != "" {
		bookingLapDetails.Payment_method = updateBookingLap.Payment_method
	}
	if updateBookingLap.Start_time != "" {
		bookingLapDetails.Start_time = updateBookingLap.Start_time
	}
	if updateBookingLap.End_time != "" {
		bookingLapDetails.End_time = updateBookingLap.End_time
	}
	db.Save(&bookingLapDetails)
	res, _ := json.Marshal(bookingLapDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

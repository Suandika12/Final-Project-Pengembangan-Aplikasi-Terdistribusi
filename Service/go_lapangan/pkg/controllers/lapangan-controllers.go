package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Suandika12/go_lapangan/pkg/models"
	"github.com/Suandika12/go_lapangan/pkg/utils"
	"github.com/gorilla/mux"
)

var NewLapangan models.Lapangan

func GetLapangan(w http.ResponseWriter, r *http.Request) {
	newLapangans := models.GetAllLapangans()
	res, _ := json.Marshal(newLapangans)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetLapanganById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lapanganId := vars["lapanganId"]
	Name, err := strconv.ParseInt(lapanganId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	lapanganDetails, _ := models.GetLapanganbyId(Name)
	res, _ := json.Marshal(lapanganDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateLapangan(w http.ResponseWriter, r *http.Request) {
	CreateLapangan := &models.Lapangan{}
	utils.ParseBody(r, CreateLapangan)
	b := CreateLapangan.CreateLapangan()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteLapangan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lapanganId := vars["lapanganId"]
	Name, err := strconv.ParseInt(lapanganId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	lapangan := models.DeleteLapangan(Name)
	res, _ := json.Marshal(lapangan)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateLapangan(w http.ResponseWriter, r *http.Request) {
	var updateLapangan = &models.Lapangan{}
	utils.ParseBody(r, updateLapangan)
	vars := mux.Vars(r)
	lapanganId := vars["lapanganId"]
	Name, err := strconv.ParseInt(lapanganId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	lapanganDetails, db := models.GetLapanganbyId(Name)
	if updateLapangan.Name != "" {
		lapanganDetails.Name = updateLapangan.Name
	}
	if updateLapangan.Description != "" {
		lapanganDetails.Description = updateLapangan.Description
	}
	if updateLapangan.Image != "" {
		lapanganDetails.Image = updateLapangan.Image
	}
	if updateLapangan.Status != "" {
		lapanganDetails.Status = updateLapangan.Status
	}
	db.Save(&lapanganDetails)
	res, _ := json.Marshal(lapanganDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

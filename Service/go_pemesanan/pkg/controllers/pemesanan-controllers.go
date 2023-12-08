package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Suandika12/go_Pemesanan/pkg/models"
	"github.com/Suandika12/go_Pemesanan/pkg/utils"
	"github.com/gorilla/mux"
)

var NewPemesanan models.Pemesanan

func GetPemesanan(w http.ResponseWriter, r *http.Request) {
	newPemesanans := models.GetAllPemesanans()
	res, _ := json.Marshal(newPemesanans)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetPemesananById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pemesananId := vars["pemesananId"]
	Code, err := strconv.ParseInt(pemesananId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	pemesananDetails, _ := models.GetPemesananbyId(Code)
	res, _ := json.Marshal(pemesananDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreatePemesanan(w http.ResponseWriter, r *http.Request) {
	CreatePemesanan := &models.Pemesanan{}
	utils.ParseBody(r, CreatePemesanan)
	b := CreatePemesanan.CreatePemesanan()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeletePemesanan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pemesananId := vars["pemesananId"]
	Code, err := strconv.ParseInt(pemesananId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	pemesanan := models.DeletePemesanan(Code)
	res, _ := json.Marshal(pemesanan)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdatePemesanan(w http.ResponseWriter, r *http.Request) {
	var updatePemesanan = &models.Pemesanan{}
	utils.ParseBody(r, updatePemesanan)
	vars := mux.Vars(r)
	pemesananId := vars["pemesananId"]
	Code, err := strconv.ParseInt(pemesananId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	pemesananDetails, db := models.GetPemesananbyId(Code)
	if updatePemesanan.Code != "" {
		pemesananDetails.Code = updatePemesanan.Code
	}
	if updatePemesanan.User_id != 0 {
		pemesananDetails.User_id = updatePemesanan.User_id
	}
	if updatePemesanan.Description != "" {
		pemesananDetails.Description = updatePemesanan.Description
	}
	if updatePemesanan.Total != "" {
		pemesananDetails.Total = updatePemesanan.Total
	}
	if updatePemesanan.Bukti_pembayaran != "" {
		pemesananDetails.Bukti_pembayaran = updatePemesanan.Bukti_pembayaran
	}
	if updatePemesanan.Payment_method != "" {
		pemesananDetails.Payment_method = updatePemesanan.Payment_method
	}
	if updatePemesanan.Status != "" {
		pemesananDetails.Status = updatePemesanan.Status
	}
	if updatePemesanan.Alamat != "" {
		pemesananDetails.Alamat = updatePemesanan.Alamat
	}
	db.Save(&pemesananDetails)
	res, _ := json.Marshal(pemesananDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

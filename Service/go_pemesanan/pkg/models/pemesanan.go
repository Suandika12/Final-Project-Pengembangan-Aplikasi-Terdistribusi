package models

import (
	"github.com/Suandika12/go_Pemesanan/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Pemesanan struct {
	gorm.Model
	User_id          uint   `gorm:""json:"user_id"`
	Code             string `gorm:""json:"code"`
	Total            string `json:"total"`
	Description      string `json:"description"`
	Bukti_pembayaran string `json:"bukti_pembayaran"`
	Payment_method   string `json:"payment_method"`
	Status           string `json:"status"`
	Alamat           string `json:"alamat"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Pemesanan{})
}
func (b *Pemesanan) CreatePemesanan() *Pemesanan {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllPemesanans() []Pemesanan {
	var Pemesanans []Pemesanan
	db.Find(&Pemesanans)
	return Pemesanans
}
func GetPemesananbyId(code int64) (*Pemesanan, *gorm.DB) {
	var getPemesanan Pemesanan
	db := db.Where("Code=?", code).Find(&getPemesanan)
	return &getPemesanan, db
}
func DeletePemesanan(code int64) Pemesanan {
	var pemesanan Pemesanan
	db.Where("Code=?", code).Delete(pemesanan)
	return pemesanan
}

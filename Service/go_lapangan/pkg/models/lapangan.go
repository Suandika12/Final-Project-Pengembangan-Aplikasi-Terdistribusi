package models

import (
	"github.com/Suandika12/go_lapangan/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Lapangan struct {
	gorm.Model

	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Status      string `json:"status"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Lapangan{})
}

func (b *Lapangan) CreateLapangan() *Lapangan {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllLapangans() []Lapangan {
	var Lapangans []Lapangan
	db.Find(&Lapangans)
	return Lapangans
}

func GetLapanganbyId(name int64) (*Lapangan, *gorm.DB) {
	var getLapangan Lapangan
	db := db.Where("Name=?", name).Find(&getLapangan)
	return &getLapangan, db
}

func DeleteLapangan(name int64) Lapangan {
	var lapangan Lapangan
	db.Where("Name=?", name).Delete(lapangan)
	return lapangan
}

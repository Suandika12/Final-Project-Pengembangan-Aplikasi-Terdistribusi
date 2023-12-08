package models

import (
	"github.com/Suandika12/go_bookingLap/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type BookingLap struct {
	gorm.Model
	User_id          uint   `gorm:""json:"user_id"`
	Field_id         uint   `gorm:""json:"field_id"`
	Field_name       uint   `gorm:""json:"field_name"`
	Description      string `json:"description"`
	Bukti_pembayaran string `json:"bukti_pembayaran"`
	Payment_method   string `json:"payment_method"`
	Start_time       string `json:"start"`
	End_time         string `json:"status"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&BookingLap{})
}
func (b *BookingLap) CreateBookingLap() *BookingLap {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllBookingLaps() []BookingLap {
	var BookingLaps []BookingLap
	db.Find(&BookingLaps)
	return BookingLaps
}
func GetBookingLapById(field_name int64) (*BookingLap, *gorm.DB) {
	var getBookingLap BookingLap
	db := db.Where("Field_name=?", field_name).Find(&getBookingLap)
	return &getBookingLap, db
}
func DeleteBookingLap(field_name int64) BookingLap {
	var BookingLap BookingLap
	db.Where("Field_name=?", field_name).Delete(BookingLap)
	return BookingLap
}

package models

import (
	"github.com/Suandika12/go_event/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Event struct {
	gorm.Model

	Title      string `json:"title"`
	Start_date string `json:"start_date"`
	End_date   string `json:"end_date"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Event{})
}

func (b *Event) CreateEvent() *Event {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllEvents() []Event {
	var Events []Event
	db.Find(&Events)
	return Events
}

func GetEventbyId(title int64) (*Event, *gorm.DB) {
	var getEvent Event
	db := db.Where("Title=?", title).Find(&getEvent)
	return &getEvent, db
}

func DeleteEvent(title int64) Event {
	var event Event
	db.Where("Title=?", title).Delete(event)
	return event
}

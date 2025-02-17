package models

import "time"

type Posts struct {
	Id          uint      `gorm:"primary_key;autoIncrement"`
	Title       string    `gorm:"type:varchar(200)"`
	Content     string    `gorm:"type:text"`
	Category    string    `gorm:"type:varchar(100)"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	Timestamp   time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
	Status      string    `gorm:"type:varchar(100)"`
}

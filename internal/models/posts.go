package models

import "time"

type Posts struct {
	ID          uint       `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Title       string     `json:"title" gorm:"type:varchar(200);column:title"`
	Content     string     `json:"content" gorm:"type:text;column:content"`
	Category    string     `json:"category" gorm:"type:varchar(100);column:category"`
	CreatedDate *time.Time `json:"created_date" gorm:"autoCreateTime;column:created_date"`
	Timestamp   *time.Time `json:"timestamp" gorm:"autoCreateTime;column:timestamp"`
	UpdatedDate *time.Time `json:"updated_date" gorm:"autoUpdateTime;column:updated_date"`
	Status      string     `json:"status" gorm:"type:varchar(100);column:status"`
}

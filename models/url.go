package models

import (
	"github.com/jinzhu/gorm"
)

type URL struct {
	gorm.Model
	LongURL  string `json:longurl`
	ShortURL string `json:longurl`
	Id       int
}

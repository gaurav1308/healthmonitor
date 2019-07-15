package model

import "github.com/jinzhu/gorm"

type(
UrlData struct {
gorm.Model
URLID       uint `db:"urlid"`
Attempts    int    `db:"attempts"`
Health string    `db:"health"`
Total_attempts int `db:total_attempts`
}

)
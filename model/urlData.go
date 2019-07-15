package model

import "github.com/jinzhu/gorm"

type(
UrlData struct {
gorm.Model
RID       uint `db:"rid"`
Attempts    int    `db:"attempts"`
Health string    `db:"health"`
Total_attempts int `db:total_attempts`
}

)
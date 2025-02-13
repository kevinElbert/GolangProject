package entity

import "gorm.io/gorm"

type PersonalData struct {
	gorm.Model
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

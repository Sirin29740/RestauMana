package database

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Cusname  string `json:"cusname"`
	Password string `json:"password"`
	Cusphone string `json:"cusphone"`
	Sex      string `json:"sex"`
}
type Staff struct {
	gorm.Model
	Staname  string `json:"staname"`
	Password string `json:"password"`
	Staphone string `json:"staphone"`
	Salary   string `json:"salary"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
	Role     string `json:"role"`
}

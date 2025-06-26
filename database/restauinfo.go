package database

import "gorm.io/gorm"

type Dish struct {
	gorm.Model
	Dishname string `json:"dishname"`
	Type     string `json:"type"`
	Price    int    `json:"price"`
	Picture  string `json:"picture"`
}
type Table struct {
	gorm.Model
	Statu  string `json:"status"`
	Maxnum int    `json:"maxnum"`
}

package database

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Dishorder []OrderItem `gorm:"foreignKey:OrderID" json:"dishorder"`
	Telephone string      `json:"telephone"`
	Tableid   int         `json:"tableid"`
	//Staffid    int         `json:"staffid"`
	Customerid int    `json:"customerid"`
	Totalprice int    `json:"totalprice"`
	Status     string `json:"status"`
}
type OrderItem struct {
	gorm.Model
	OrderID  uint   `json:"orderid"`
	DishID   uint   `json:"dishid"`
	Dishname string `json:"dishname"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

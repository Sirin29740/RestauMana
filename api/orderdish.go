package api

import (
	"RestauMana/database"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderRequest struct {
	Items []struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		Price    int    `json:"price"`
		Quantity int    `json:"quantity"`
	} `json:"items"`
}

func NewOrder(c *gin.Context) {
	db := database.Getdb()
	session := sessions.Default(c)
	tableid, _ := strconv.Atoi(c.Param("id"))
	var req OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "格式错误"})
		return
	}

	var total int
	var orderItems []database.OrderItem
	for _, item := range req.Items {
		total += item.Price * item.Quantity
		orderItems = append(orderItems, database.OrderItem{
			DishID:   item.ID,
			Dishname: item.Name,
			Quantity: item.Quantity,
			Price:    item.Price,
		})
	}
	useridRaw := session.Get("id")
	useridUint, ok := useridRaw.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户未登录或ID类型错误"})
		return
	}
	userid := int(useridUint)

	fmt.Println("userid:", userid)
	var user database.Customer
	db.Where("id = ?", userid).First(&user)
	order := database.Order{
		Dishorder:  orderItems,
		Totalprice: total,
		Tableid:    tableid,
		Customerid: userid,
		Telephone:  user.Cusphone,
		Status:     "进行中", // 添加订单状态字段
	}
	db.Create(&order)

	if err := db.Model(&database.Table{}).Where("id = ?", tableid).Update("statu", "已占用").Error; err != nil {
		c.JSON(500, gin.H{"error": "更新餐桌状态失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "订单创建成功",
		"id":      order.ID,
	})
}
func Listorder(c *gin.Context) {
	db := database.Getdb()

	var orders []database.Order
	err := db.Preload("Dishorder").Find(&orders).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询订单失败"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
func FinishOrder(c *gin.Context) {
	db := database.Getdb()
	orderid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "订单ID错误"})
		return
	}

	var order database.Order
	if err := db.Where("ID = ?", orderid).First(&order).Error; err != nil {
		c.JSON(404, gin.H{"error": "订单不存在"})
		return
	}

	// 更新订单状态为已完成
	if err := db.Model(&order).Update("status", "已完成").Error; err != nil {
		c.JSON(500, gin.H{"error": "更新订单状态失败"})
		return
	}

	// 更新对应餐桌状态为空闲
	if err := db.Model(&database.Table{}).Where("id = ?", order.Tableid).Update("statu", "空闲").Error; err != nil {
		c.JSON(500, gin.H{"error": "更新餐桌状态失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "订单已结束，餐桌已空闲"})
}
func ListOngoingOrders(c *gin.Context) {
	db := database.Getdb()

	var orders []database.Order
	// 关联查询订单菜品等，过滤状态为“进行中”的订单
	err := db.Preload("Dishorder").Where("status = ?", "进行中").Find(&orders).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询订单失败"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

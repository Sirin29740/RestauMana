package api

import (
	"RestauMana/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CusRegister(c *gin.Context) {
	db := database.Getdb()
	var cus database.Customer
	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(400, gin.H{"msg": "参数解析失败"})
		return
	}
	fmt.Println(cus.Password)
	if len(cus.Password) < 6 {
		c.JSON(422, "密码不得少于六位")
		return
	}
	if len(cus.Cusname) == 0 {
		c.JSON(422, "用户名不得为空")
		return
	}
	if len(cus.Cusphone) == 0 {
		c.JSON(422, "电话号码不得为空")
		return
	}
	db.Create(&cus)
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
		"data": gin.H{
			"id":       cus.ID,
			"name":     cus.Cusname,
			"phone":    cus.Cusphone,
			"password": cus.Password,
		},
	})
}

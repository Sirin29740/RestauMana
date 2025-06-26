package api

import (
	"RestauMana/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Newstaff(c *gin.Context) {
	db := database.Getdb()
	var sta database.Staff
	if err := c.ShouldBindJSON(&sta); err != nil {
		c.JSON(400, gin.H{"msg": "参数解析失败"})
		return
	}
	sta.Password = "123456"
	db.Create(&sta)
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
		"data": gin.H{
			"id":    sta.ID,
			"name":  sta.Staname,
			"phone": sta.Staphone,
		},
	})
}
func Liststaff(c *gin.Context) {
	db := database.Getdb()
	var stas []database.Staff
	result := db.Find(&stas)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}
	c.JSON(200, stas)
}
func Deletestaff(c *gin.Context) {
	id := c.Param("id")
	db := database.Getdb()
	var sta database.Staff
	db.Where("ID = ?", id).Delete(&sta)
	c.JSON(http.StatusOK, gin.H{
		"message": "delete ok",
	})
}
func Staffinfo(c *gin.Context) {
	id := c.Param("id")
	db := database.Getdb()
	var sta database.Staff
	db.Where("ID = ?", id).First(&sta)
	c.JSON(http.StatusOK, sta)
}

func Updatestaff(c *gin.Context) {
	id := c.Param("id")
	db := database.Getdb()
	var sta database.Staff
	db.Where("ID = ?", id).First(&sta)
	var input struct {
		Staname  string `json:"staname"`
		Staphone string `json:"staphone"`
		Salary   string `json:"salary"`
		Age      int    `json:"age"`
		Sex      string `json:"sex"`
		Role     string `json:"role"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定错误"})
		return
	}

	// 更新字段
	sta.Staname = input.Staname
	sta.Staphone = input.Staphone
	sta.Salary = input.Salary
	sta.Age = input.Age
	sta.Sex = input.Sex
	sta.Role = input.Role

	db.Save(&sta)
	c.JSON(http.StatusOK, gin.H{
		"message": "update ok",
	})
}

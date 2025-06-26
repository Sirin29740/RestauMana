package api

import (
	"RestauMana/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// 添加新餐桌
func Newtable(c *gin.Context) {
	db := database.Getdb()
	var table database.Table
	if err := c.ShouldBindJSON(&table); err != nil {
		c.JSON(400, gin.H{"msg": "参数解析失败"})
		return
	}
	if table.Statu == "" {
		table.Statu = "空闲"
	}
	db.Create(&table)
	c.JSON(http.StatusOK, gin.H{
		"message": "添加成功",
		"data":    table,
	})
}

// 查询所有餐桌
func Listtables(c *gin.Context) {
	db := database.Getdb()
	var tables []database.Table
	result := db.Find(&tables)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}
	c.JSON(http.StatusOK, tables)
}

// 删除餐桌
func Deletetable(c *gin.Context) {
	id := c.Param("id")
	db := database.Getdb()
	result := db.Delete(&database.Table{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// 查询某一张餐桌详情
func Tableinfo(c *gin.Context) {
	id := c.Param("id")
	db := database.Getdb()
	var table database.Table
	if err := db.First(&table, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "餐桌不存在"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		}
		return
	}
	c.JSON(http.StatusOK, table)
}

// 修改餐桌信息
func Updatetable(c *gin.Context) {
	id := c.Param("id")
	db := database.Getdb()
	var table database.Table
	if err := db.First(&table, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查找失败"})
		return
	}

	var input struct {
		Statu  string `json:"status"`
		Maxnum int    `json:"maxnum"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定错误"})
		return
	}

	table.Statu = input.Statu
	table.Maxnum = input.Maxnum

	if err := db.Save(&table).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

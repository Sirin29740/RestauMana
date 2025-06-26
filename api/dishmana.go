package api

import (
	"RestauMana/database"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
	"path/filepath"
)

func Newdishes(c *gin.Context) {
	db := database.Getdb()
	// 解析普通字段
	dishname := c.PostForm("dishname")
	dishType := c.PostForm("type")
	priceStr := c.PostForm("price")

	if dishname == "" || dishType == "" || priceStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少必填字段"})
		return
	}
	var di database.Dish
	result0 := db.Where("dishname = ?", dishname).First(&di)
	if result0.Error == nil {
		// 找到了，说明菜名已存在
		c.JSON(400, gin.H{"error": "菜名已存在"})
		return
	}
	if !errors.Is(result0.Error, gorm.ErrRecordNotFound) {
		// 查询时出现其他错误
		c.JSON(500, gin.H{"error": "查询数据库失败"})
		return
	}
	// 转换价格为int
	var price int
	_, err := fmt.Sscanf(priceStr, "%d", &price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "价格格式错误"})
		return
	}

	// 处理图片上传
	file, err := c.FormFile("picture")
	picturePath := "" // 图片路径，默认空
	if err == nil {
		// 图片上传了，保存文件到指定目录
		// 目录你需要保证存在，比如 ./uploads/
		ext := filepath.Ext(file.Filename)             // 取上传文件后缀名，比如 ".jpg" ".png"
		filename := fmt.Sprintf("%s%s", dishname, ext) // 拼成：菜名 + 后缀

		savePath := filepath.Join("uploads", filename)

		err = c.SaveUploadedFile(file, savePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "图片保存失败"})
			return
		}
		picturePath = "/" + savePath // 或根据你的静态资源路径调整
	}

	// 构造Dish结构体
	dish := database.Dish{
		Dishname: dishname,
		Type:     dishType,
		Price:    price,
		Picture:  picturePath,
	}

	// 保存到数据库
	result := db.Create(&dish)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存菜品失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "菜谱添加成功",
		"dish":    dish,
	})
}

func Listdish(c *gin.Context) {
	db := database.Getdb()
	var dishes []database.Dish
	result := db.Find(&dishes)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}
	c.JSON(200, dishes)
}
func Deletedish(c *gin.Context) {
	id := c.Param("id")
	db := database.Getdb()
	var dish database.Dish
	db.Where("ID = ?", id).Delete(&dish)
	c.JSON(http.StatusOK, gin.H{
		"message": "delete ok",
	})
}
func Dishinfo(c *gin.Context) {
	id := c.Param("id")
	db := database.Getdb()
	var dish database.Dish
	db.Where("ID = ?", id).First(&dish)
	c.JSON(http.StatusOK, dish)
}
func Updatedish(c *gin.Context) {
	id := c.Param("id")
	db := database.Getdb()

	// 查找现有菜品
	var dish database.Dish
	if err := db.First(&dish, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "菜品不存在"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库查询失败"})
		return
	}

	// 提取表单字段
	dishname := c.PostForm("dishname")
	dishType := c.PostForm("type")
	priceStr := c.PostForm("price")

	if dishname == "" || dishType == "" || priceStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "字段不能为空"})
		return
	}

	var price int
	if _, err := fmt.Sscanf(priceStr, "%d", &price); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "价格必须是整数"})
		return
	}

	// 检查菜名是否重复（排除当前ID）
	var exist database.Dish
	if err := db.Where("dishname = ? AND id != ?", dishname, id).First(&exist).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "该菜名已存在"})
		return
	}

	// 处理图片上传
	file, err := c.FormFile("picture")
	picturePath := "" // 图片路径，默认空
	if err == nil {
		// 如果上传了新图片，先删除旧图片
		if dish.Picture != "" {
			oldPath := dish.Picture
			if oldPath[0] == '/' {
				oldPath = oldPath[1:] // 去掉前导 /
			}
			if err := os.Remove(oldPath); err != nil && !os.IsNotExist(err) {
				// 不是“文件不存在”的错误才报错
				c.JSON(http.StatusInternalServerError, gin.H{"error": "旧图片删除失败"})
				return
			}
		}
		// 图片上传了，保存文件到指定目录
		// 目录你需要保证存在，比如 ./uploads/
		ext := filepath.Ext(file.Filename)             // 取上传文件后缀名，比如 ".jpg" ".png"
		filename := fmt.Sprintf("%s%s", dishname, ext) // 拼成：菜名 + 后缀

		savePath := filepath.Join("uploads", filename)

		err = c.SaveUploadedFile(file, savePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "图片保存失败"})
			return
		}
		picturePath = "/" + savePath // 或根据你的静态资源路径调整
	}
	// 更新字段
	dish.Dishname = dishname
	dish.Type = dishType
	dish.Price = price
	dish.Picture = picturePath
	if err := db.Save(&dish).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

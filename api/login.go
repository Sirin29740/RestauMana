package api

import (
	"RestauMana/database"
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type receive struct {
	Telephone string `json:"phone"`
	Password  string `json:"password"`
}

func Cuslogin(c *gin.Context) {
	db := database.Getdb()
	session := sessions.Default(c)
	var recus receive
	if err := c.ShouldBindJSON(&recus); err != nil {
		c.JSON(400, gin.H{"msg": "参数解析失败"})
		return
	}
	phone := recus.Telephone
	password := recus.Password
	var user database.Customer
	result := db.Where("cusphone = ?", phone).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "该号码未注册",
			})
		} else {
			log.Printf("查询用户失败 - 手机号: %s, 错误: %v", phone, result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "无法查询用户信息",
			})
		}
		return
	}
	if user.Password != password {
		c.JSON(401, gin.H{
			"error": "wrong passwrd114514",
		})
		return
	}
	userid := user.ID
	session.Set("id", userid)
	fmt.Printf("设置 session 时 ID 类型是：%T\n", userid)

	session.Set("role", "customer")
	session.Save()
	fmt.Println("Session ID:", session.Get("id")) // 打印看看
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
		"data": gin.H{
			"phone":    user.Cusphone,
			"password": user.Password,
		},
	})
}

func Stalogin(c *gin.Context) {
	db := database.Getdb()
	session := sessions.Default(c)

	// 从 POST 表单中获取数据
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	var user database.Staff
	result := db.Where("staphone = ?", phone).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "没有该员工",
			})
		} else {
			log.Printf("查询用户失败 - 手机号: %s, 错误: %v", phone, result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "无法查询用户信息",
			})
		}
		return
	}
	if user.Password != password {
		c.JSON(401, gin.H{
			"error": "wrong passwrd114514",
		})
		return
	}
	userid := user.ID
	session.Set("id", int(userid))
	session.Set("role", user.Role)
	session.Save()
	fmt.Println("Session ID:", session.Get("id")) // 打印看看

	if user.Role == "管理员" {
		c.Redirect(http.StatusSeeOther, "/admin/adpage")
	} else {
		c.Redirect(http.StatusSeeOther, "/staff/stapage")
	}

}
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear() // 清空所有 session
	session.Save()  // 保存修改
	c.JSON(http.StatusOK, gin.H{"message": "退出登录成功"})
}

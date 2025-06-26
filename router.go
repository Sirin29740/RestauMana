package main

import (
	"RestauMana/api"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret")) // 密钥
	router.Use(sessions.Sessions("mysession", store))
	router.Static("/uploads", "./uploads")
	router.LoadHTMLGlob("html/**/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	user := router.Group("/user")
	{
		user.GET("/register", func(c *gin.Context) {
			c.HTML(http.StatusOK, "register.html", gin.H{})
		})
		user.POST("/register", api.CusRegister)
		user.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", gin.H{})
		})
		user.POST("/login", api.Cuslogin)
		user.GET("/stalogin", func(c *gin.Context) {
			c.HTML(http.StatusOK, "stafflogin.html", gin.H{})
		})
		user.POST("/stalogin", api.Stalogin)
		user.GET("/logout", api.Logout)
	}
	admin := router.Group("/admin", api.AuthMiddleware(api.RoleAdmin))
	{
		admin.GET("/newstaff", func(c *gin.Context) {
			c.HTML(http.StatusOK, "newstaff.html", gin.H{})
		})
		admin.POST("/newstaff", api.Newstaff)
		admin.GET("/staffinfo/:id", api.Staffinfo)
		admin.GET("/manastaff", func(c *gin.Context) {
			c.HTML(http.StatusOK, "manastaff.html", gin.H{})
		})
		admin.GET("/liststaff", api.Liststaff)
		admin.DELETE("/deletestaff/:id", api.Deletestaff)
		admin.GET("/updatestaff/:id", func(c *gin.Context) {
			c.HTML(http.StatusOK, "updatestaff.html", gin.H{})
		})
		admin.POST("/updatestaff/:id", api.Updatestaff)
		admin.GET("/adpage", func(c *gin.Context) {
			c.HTML(http.StatusOK, "adpage.html", gin.H{})
		})
	}
	staff := router.Group("/staff", api.AuthMiddleware(api.RoleStaff, api.RoleAdmin))
	{
		staff.GET("/newdish", func(c *gin.Context) {
			c.HTML(http.StatusOK, "newdish.html", gin.H{})
		})
		staff.POST("/newdish", api.Newdishes)
		staff.GET("/manadishes", func(c *gin.Context) {
			c.HTML(http.StatusOK, "manadishes.html", gin.H{})
		})
		staff.DELETE("/deletedish/:id", api.Deletedish)
		staff.GET("/updatedish/:id", func(c *gin.Context) {
			c.HTML(http.StatusOK, "updatedish.html", gin.H{})
		})
		staff.POST("/updatedish/:id", api.Updatedish)
		staff.GET("/newtable", func(c *gin.Context) {
			c.HTML(http.StatusOK, "newtable.html", gin.H{})
		})
		staff.POST("/newtable", api.Newtable)
		staff.GET("/manatable", func(c *gin.Context) {
			c.HTML(http.StatusOK, "manatable.html", gin.H{})
		})
		staff.DELETE("/deletetable/:id", api.Deletetable)
		staff.GET("/updatetable/:id", func(c *gin.Context) {
			c.HTML(http.StatusOK, "updatetable.html", gin.H{})
		})
		staff.POST("/updatetable/:id", api.Updatetable)
		staff.GET("/manaorder", func(c *gin.Context) {
			c.HTML(http.StatusOK, "manaorder.html", gin.H{})
		})
		staff.GET("/listgoorder", api.ListOngoingOrders)
		staff.POST("/finishorder/:id", api.FinishOrder)
		staff.GET("/stapage", func(c *gin.Context) {
			c.HTML(http.StatusOK, "stapage.html", gin.H{})
		})
	}
	order := router.Group("/order", api.AuthMiddleware(api.RoleStaff, api.RoleAdmin, api.RoleCustomer))
	{
		order.GET("/listdishes", api.Listdish)
		order.GET("/dishinfo/:id", api.Dishinfo)
		order.GET("/tableinfo/:id", api.Tableinfo)
		order.GET("/listtables", api.Listtables)
		order.GET("/cuspage", func(c *gin.Context) {
			c.HTML(http.StatusOK, "cuspage.html", gin.H{})
		})
		order.GET("/choosetable", func(c *gin.Context) {
			c.HTML(http.StatusOK, "choosetable.html", gin.H{})
		})
		order.GET("/neworder/:id", func(c *gin.Context) {
			c.HTML(http.StatusOK, "neworder.html", gin.H{})
		})
		order.POST("/neworder/:id", api.NewOrder)
		order.GET("/vieworder", func(c *gin.Context) {
			c.HTML(http.StatusOK, "vieworder.html", gin.H{})
		})
		order.GET("/listorder", api.Listorder)
	}
	router.Run(":8080")
}

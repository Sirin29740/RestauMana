package api

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Role string

const (
	RoleAdmin    Role = "管理员"
	RoleStaff    Role = "普通员工"
	RoleCustomer Role = "customer"
	RoleGuest    Role = "guest" // 未登录
)

// AuthMiddleware 传入允许的角色，返回中间件
func AuthMiddleware(allowedRoles ...Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		roleVal := session.Get("role")
		var role string
		if roleVal == nil {
			role = string(RoleGuest)
		} else {
			role, _ = roleVal.(string)
		}

		allowed := false
		for _, r := range allowedRoles {
			if role == string(r) {
				allowed = true
				break
			}
		}

		if !allowed {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "权限不足或未登录",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

package handler

import (
	"api-gateway-service/generated/users"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userId string

		// URL parametrlardan user-id ni olish
		userId = c.Param("user-id")
		if userId == "" {
			// POST so'rovining bodysi orqali user-id ni olish
			var req struct {
				UserId string `json:"userId" binding:"required"`
			}
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
				c.Abort()
				return
			}
			userId = req.UserId
		}

		// Userni tekshirish
		resp, err := h.User.CheckUser(c, &users.UserRequest{UserId: userId})
		if err != nil {
			fmt.Println(errors.Is(err, sql.ErrNoRows), err)
			if !errors.Is(err, sql.ErrNoRows) || (resp != nil && !resp.Exists) {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				c.Abort()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if resp == nil || !resp.Exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		fmt.Println("user mavjud")

		// User mavjud bo'lsa, keyingi handlerga o'tkazish
		c.Next()
	}
}




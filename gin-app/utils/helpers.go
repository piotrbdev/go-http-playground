package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) (string, error) {
	val, exists := c.Get("userID")
	if !exists {
		return "", fmt.Errorf("user id not found in context")
	}

	idString, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("invalid user id type")
	}

	return idString, nil
}

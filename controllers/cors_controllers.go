package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"92HW/services"
)

func AddTrustedOrigin(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	var origin struct {
		Origin string `json:"origin"`
	}
	if err := c.ShouldBindJSON(&origin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := services.AddOrigin(userID, origin.Origin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Origin added"})
}

func RemoveTrustedOrigin(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	var origin struct {
		Origin string `json:"origin"`
	}
	if err := c.ShouldBindJSON(&origin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := services.RemoveOrigin(userID, origin.Origin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Origin removed"})
}

func GetTrustedOrigins(c *gin.Context) {
	userID := c.MustGet("userID").(string)

	origins, err := services.GetOrigins(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"origins": origins})
}

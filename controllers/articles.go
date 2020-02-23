package controllers

import "github.com/gin-gonic/gin"

type Articles struct {
}

func (a *Articles) FindAll(c *gin.Context) {
	c.JSON(200, gin.H{"hello": "World"})
}

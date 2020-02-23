package controllers

import (
	"go-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type Articles struct {
	DB *gorm.DB
}

type CreateArticlePayload struct {
	Title   string `binding:"required"`
	Excerpt string `binding:"required"`
	Body    string `binding:"required"`
}

func (a *Articles) FindAll(c *gin.Context) {
	c.JSON(200, gin.H{"hello": "World"})
}

func (a *Articles) FindOne(c *gin.Context) {
	c.JSON(200, gin.H{"hello": "World"})
}

func (a *Articles) Create(c *gin.Context) {
	var form CreateArticlePayload

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var article models.Article
	copier.Copy(&article, &form)
	a.DB.Create(&article)

	c.JSON(http.StatusCreated, gin.H{"article": article})
}

func (a *Articles) Update(c *gin.Context) {
	c.JSON(200, gin.H{"hello": "World"})
}

func (a *Articles) Delete(c *gin.Context) {
	c.JSON(200, gin.H{"hello": "World"})
}

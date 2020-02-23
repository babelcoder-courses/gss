package routes

import (
	"go-rest-api/config"
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine) {
	db := config.GetDB()
	v1 := r.Group("/api/v1")

	articlesCtrl := controllers.Articles{DB: db}
	v1.GET("/articles", articlesCtrl.FindAll)
	v1.POST("/articles", articlesCtrl.Create)
}

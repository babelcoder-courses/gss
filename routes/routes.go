package routes

import (
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	articlesCtrl := controllers.Articles{}
	v1.GET("/articles", articlesCtrl.FindAll)
}

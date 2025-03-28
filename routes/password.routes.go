package routes

import (
	"github.com/gin-gonic/gin"
	"www.github.com/NirajSalunke/controllers"
)

func LoadPasswordRoutes(r *gin.RouterGroup) {

	r.POST("/analyze-password", controllers.AnalyzePassword)
	r.POST("/suggest", controllers.Suggest)
}

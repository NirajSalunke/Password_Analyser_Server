package routes

import "github.com/gin-gonic/gin"

func LoadRoutes(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello, World!"})
	})
	PasswordGroup := r.Group("/password")
	LoadPasswordRoutes(PasswordGroup)

}

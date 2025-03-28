package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"www.github.com/NirajSalunke/config"
	"www.github.com/NirajSalunke/routes"
)

func init() {
	config.LoadEnv()
	config.SetupGemini()
}
func main() {

	r := gin.Default()
	routes.LoadRoutes(r)
	r.Run(":" + os.Getenv("PORT"))

}

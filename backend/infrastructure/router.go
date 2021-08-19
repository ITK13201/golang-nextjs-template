package infrastructure

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ITK13201/golang-nextjs-template/backend/config"
	"github.com/ITK13201/golang-nextjs-template/backend/interfaces/controllers"
)

func InitRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	ginCfg := cors.DefaultConfig()
	ginCfg.AllowOrigins = []string{"*"}
	router.Use(cors.New(ginCfg))

	userController := controllers.NewUserController(NewSqlHandler(cfg))

	router.POST("/api/users/", func(c *gin.Context) { userController.Create(c) })
	router.GET("/api/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/api/users/:id", func(c *gin.Context) { userController.Show(c) })

	return router
}

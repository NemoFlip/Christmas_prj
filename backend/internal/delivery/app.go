package delivery

import (
	"Christmas_prj/backend/internal/delivery/handlers"
	"Christmas_prj/backend/internal/delivery/router"
	"Christmas_prj/backend/pkg/log"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Christmas prj
// @description Simple project with recommendation system
// @host localhost:8080
// @BasePath /
func StartServer(logger *log.Logger, giftServer *handlers.GiftServer) {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.InitRouting(r, giftServer)

	if err := r.Run(":8080"); err != nil {
		logger.ErrorLogger.Fatal().Msgf("unable to run server on port(8080): %s", err)
	}
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.InitRouting(r, giftServer)

	if err := r.Run(":8080"); err != nil {
		logger.ErrorLogger.Fatal().Msgf("unable to run server on port(8080): %s", err)
	}
}

package delivery

import (
	"Christmas_prj/backend/pkg/log"
	"github.com/gin-gonic/gin"
)

func StartServer(logger *log.Logger) {
	r := gin.Default()

	if err := r.Run(":8080"); err != nil {
		logger.ErrorLogger.Fatal().Msgf("unable to run server on port(8080): %s", err)
	}
}

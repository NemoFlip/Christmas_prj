package router

import (
	"Christmas_prj/backend/internal/delivery/handlers"
	"Christmas_prj/backend/pkg/log"
	"github.com/gin-gonic/gin"
)

func InitRouting(r *gin.Engine, logger *log.Logger, apiServer *handlers.APIServer) {
	r.POST("/api/users/data", apiServer.PostUserInfo)
}

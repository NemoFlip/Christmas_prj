package router

import (
	"Christmas_prj/backend/internal/delivery/handlers"
	"github.com/gin-gonic/gin"
)

func InitRouting(r *gin.Engine, apiServer *handlers.GiftServer) {
	r.POST("/api/users/data", apiServer.PostGiftInfo)
}

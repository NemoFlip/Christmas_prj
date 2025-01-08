package handlers

import (
	"Christmas_prj/backend/internal/entity"
	"Christmas_prj/backend/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type APIServer struct {
	logger *log.Logger
}

func NewAPIServer(logger *log.Logger) *APIServer {
	return &APIServer{logger: logger}
}

// @Summary User request
// @Description post request from user for present recommendation
// @Tags API
// @Accept json
// @Param user body entity.User true "user's preferences"
// @Success 200 {object} string "user's preferences has been accepted"
// @Failure 400 {object} string "invalid user's data"
// @Router /api/users/data [post]
func (s *APIServer) PostUserInfo(ctx *gin.Context) {
	var user entity.User
	if err := ctx.BindJSON(&user); err != nil {
		s.logger.ErrorLogger.Error().Msgf("unable to bind user: %s", err)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

}

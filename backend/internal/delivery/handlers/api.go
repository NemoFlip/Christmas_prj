package handlers

import (
	"Christmas_prj/backend/internal/entity"
	"Christmas_prj/backend/pkg/log"
	"bytes"
	"encoding/json"
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
// @Produce json
// @Param user body entity.User true "user's preferences"
// @Success 200 {object} map[string]string "user's preferences has been accepted"
// @Failure 400 {object} string "invalid user's data"
// @Router /api/users/data [post]
func (s *APIServer) PostUserInfo(ctx *gin.Context) {
	var user entity.User
	if err := ctx.BindJSON(&user); err != nil {
		s.logger.ErrorLogger.Error().Msgf("unable to bind user: %s", err)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	userJSON, err := json.Marshal(user)
	if err != nil {
		s.logger.ErrorLogger.Error().Msgf("unable to convert user struct to json: %s", err)
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	r := bytes.NewBuffer(userJSON)
	mlURL := "https://ml-model-recommendation"
	resp, err := http.Post(mlURL, "application/json", r) //TODO: replace url
	if err != nil {
		s.logger.ErrorLogger.Error().Msgf("unable to send request to ml: %s", err)
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		s.logger.ErrorLogger.Error().Msgf("ML model responded with status: %d", resp.StatusCode)
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": resp.Status,
	})
}

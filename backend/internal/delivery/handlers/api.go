package handlers

import (
	"Christmas_prj/backend/internal/gateways"
	"Christmas_prj/backend/internal/payload"
	"Christmas_prj/backend/pkg/log"
	"Christmas_prj/backend/pkg/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GiftServer struct {
	logger    *log.Logger
	mlGateway *gateways.MLGateway
}

func NewGiftServer(logger *log.Logger, mlGateway *gateways.MLGateway) *GiftServer {
	return &GiftServer{logger: logger, mlGateway: mlGateway}
}

// @Summary User request
// @Description post request from user for present recommendation
// @Tags API
// @Accept json
// @Produce json
// @Param user body payload.GiftRequest true "user's preferences"
// @Success 200 {object} payload.GiftResponse "Recommended gift's details"
// @Failure 400 {object} string "invalid user's data"
// @Failure 500 {object} string "internal server error"
// @Router /api/users/data [post]
func (as *GiftServer) PostGiftInfo(ctx *gin.Context) {
	var giftRequest payload.GiftRequest
	if err := utility.ParseJson(ctx, &giftRequest); err != nil {
		as.logger.ErrorLogger.Error().Msg(err.Error())
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	as.logger.InfoLogger.Info().Msgf("gift Request: %s", giftRequest.Description)
	giftRequest.TopN = 10
	fmt.Println(giftRequest)

	giftResponse, err := as.mlGateway.GetGiftRecommendation(giftRequest)
	if err != nil {
		as.logger.ErrorLogger.Error().Msg(err.Error())
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	as.logger.InfoLogger.Info().Msgf("Response in post: %+v", giftResponse)
	ctx.JSON(http.StatusOK, giftResponse)
}

package main

import (
	_ "Christmas_prj/backend/docs"
	"Christmas_prj/backend/internal/delivery"
	"Christmas_prj/backend/internal/delivery/handlers"
	"Christmas_prj/backend/internal/gateways"
	"Christmas_prj/backend/pkg/log"
)

func main() {
	logger := log.InitLogger()

	mlURL := "http://ml_model:8000/recommend"
	mlGateway := gateways.NewMLGateway(mlURL)
	giftServer := handlers.NewGiftServer(logger, mlGateway)

	delivery.StartServer(logger, giftServer)
}

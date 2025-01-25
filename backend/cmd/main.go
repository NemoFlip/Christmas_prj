package main

import (
	_ "Christmas_prj/backend/docs"
	"Christmas_prj/backend/internal/delivery"
	"Christmas_prj/backend/internal/delivery/handlers"
	"Christmas_prj/backend/pkg/log"
)

func main() {
	logger := log.InitLogger()

	apiServer := handlers.NewAPIServer(logger)
	delivery.StartServer(logger, apiServer)
}

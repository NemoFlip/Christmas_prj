package main

import (
	"Christmas_prj/backend/internal/database"
	"Christmas_prj/backend/internal/delivery"
	"Christmas_prj/backend/internal/delivery/handlers"
	"Christmas_prj/backend/pkg"
	"Christmas_prj/backend/pkg/log"
	_ "Christmas_prj/docs"
	"os"
)

const (
	DBName     = "POSTGRES_DB"
	DBUser     = "POSTGRES_USER"
	DBPassword = "POSTGRES_PASSWORD"
)

func main() {
	logger := log.InitLogger()
	postgresDBname, exists := os.LookupEnv(DBName)
	if !exists {
		logger.ErrorLogger.Fatal().Msgf("unable to find %s env variable", DBName)
	}
	postgresUser, exists := os.LookupEnv(DBUser)
	if !exists {
		logger.ErrorLogger.Fatal().Msgf("unable to find %s env variable", DBUser)
	}
	postgresPassword, exists := os.LookupEnv(DBPassword)
	if !exists {
		logger.ErrorLogger.Fatal().Msgf("unable to find %s env variable", DBPassword)
	}

	db, err := pkg.ConnectToPostgres(
		"products_postgres",
		postgresUser,
		postgresPassword,
		"5432",
		postgresDBname,
		"disable",
	)
	if err != nil {
		logger.ErrorLogger.Fatal().Msg(err.Error())
	}
	_ = database.NewProductsStorage(db)
	apiServer := handlers.NewAPIServer(logger)
	delivery.StartServer(logger, apiServer)
}

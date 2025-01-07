package pkg

import (
	"database/sql"
	"fmt"
)

func ConnectToPostgres(host, user, password, port, dbname, sslmode string) (*sql.DB, error) {
	sourceString := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s",
		host, user, password, port, dbname, sslmode)
	db, err := sql.Open("postgres", sourceString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to db(%s): %w", dbname, err)
	}
	return db, nil
}

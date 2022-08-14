package db

import (
	"database/sql"
	"fmt"
	"go-grpc/server/utils/helpers"
)

/*PostgresConnect - Connect function which is used in the main package for database connection */
func PostgresConnect() (*sql.DB, error) {

	/* Loading TOML file */
	config, err := helpers.LoadEnvFile()
	if err != nil {
		return nil, err
	}

	host := config.Get("postgres.host").(string)
	port := config.Get("postgres.port").(int64)
	user := config.Get("postgres.user").(string)
	password := config.Get("postgres.password").(string)
	dbName := config.Get("postgres.dbName").(string)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

package extract

import (
	"database/sql"
	"go-snowflake-app/internal/config"
	"log"

	sf "github.com/snowflakedb/gosnowflake"
)

const SNOWFLAKE = "snowflake"

// OpenDbConnection function opens a connection to Snowflake based on the config.SnowflakeConfig config.
func OpenDbConnection(config config.SnowflakeConfig) *sql.DB {
	cfg := &sf.Config{
		Account:   config.Account,
		User:      config.Username,
		Password:  config.Password,
		Database:  config.Database,
		Schema:    config.Schema,
		Warehouse: config.Warehouse,
	}

	dsnString, dsnErr := sf.DSN(cfg)
	if dsnErr != nil {
		log.Println("DSN String err: ", dsnErr)
		return nil
	}

	db, dbErr := sql.Open(SNOWFLAKE, dsnString)
	if dbErr != nil {
		log.Println("DB Conn err: ", dbErr)
		return nil
	}

	log.Println("Connection success.")
	return db
}

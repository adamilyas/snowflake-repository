package extract

import (
	"database/sql"
	"go-snowflake-app/internal/config"
	"log"
)

const (
	Fields = "S_STORE_ID, S_STORE_NAME, S_COMPANY_NAME"
	QUERY  = "SELECT " + Fields + " FROM STORE LIMIT 10"
)

// StoreDAO encapsulates all access to the Store table on Snowflake
type StoreDAO struct {
	db *sql.DB
}

// Store struct is a representation of data from Snowflake.
type Store struct {
	StoreID     sql.NullString
	StoreName   sql.NullString
	CompanyName sql.NullString
}

func NewStoreDAO(config config.SnowflakeConfig) *StoreDAO {
	return &StoreDAO{db: OpenDbConnection(config)}
}

func (s *StoreDAO) FindAll() ([]Store, error) {
	var results []Store

	log.Println("Query: ", QUERY)
	rows, err := s.db.Query(QUERY)
	if err != nil {
		log.Println("DB Query err: ", err)
		return nil, err
	}

	for rows.Next() {
		var result Store

		// The mapping occurs via the position of the field
		err := rows.Scan(
			&result.StoreID,
			&result.StoreName,
			&result.CompanyName,
		)

		if err != nil {
			log.Println("DB Scan err: ", err)
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

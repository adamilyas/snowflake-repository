package main

import (
	"go-snowflake-app/internal/config"
	"go-snowflake-app/internal/extract"
	"log"
)

func main() {

	dev := config.DevConfig{}
	devConfig := dev.LoadSnowflakeConfig()
	storeDAO := extract.NewStoreDAO(devConfig)

	results, err := storeDAO.FindAll()

	if err != nil {
		log.Fatalln(err)
	}

	for _, result := range results {
		log.Println(result)
	}
}

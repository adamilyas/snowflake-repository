package config

// SnowflakeConfig is a struct that stores the config that is needed to access Snowflake on Azure
type SnowflakeConfig struct {
	Database  string
	Warehouse string
	Username  string
	Password  string
	Account   string
	Schema    string
}

// SnowflakeConfigHandler ...
type SnowflakeConfigHandler interface {
	LoadSnowflakeConfig() SnowflakeConfig
}

// DevConfig implements SnowflakeConfigHandler
type DevConfig struct{}

// LoadSnowflakeConfig ...
func (d *DevConfig) LoadSnowflakeConfig() SnowflakeConfig {
	// Read
	return SnowflakeConfig{
		Database:  "SNOWFLAKE_SAMPLE_DATA",
		Warehouse: "COMPUTE_WH",
		Username:  "", // TODO: Change me
		Password:  "", // TODO: Change me
		Account:   "", // TODO: Change me
		Schema:    "TPCDS_SF100TCL",
	}
}

package configuration

type AppConfig struct {
	Database DatabaseConfig `json:"database"`
	Hash     HashKeyConfig  `json:"hash_key"`
	NewRelic NewRelicConfig `json:"new_relic"`
	Redis    RedisConfig    `json:"redis"`
}

type DatabaseConfig struct {
	Driver       string `json:"driver"`
	Host         string `json:"host"`
	DatabaseName string `json:"database_name"`
	Password     string `json:"password"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
}

type HashKeyConfig struct {
	Password string `json:"password"`
}

type NewRelicConfig struct {
	UserKey    string `json:"user_key"`
	LicenseKey string `json:"license_key"`
}

type RedisConfig struct {
	Address  string `json:"address"`
	Password string `json:"password"`
}

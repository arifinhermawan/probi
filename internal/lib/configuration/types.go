package configuration

type AppConfig struct {
	// secret config
	Database DatabaseConfig `json:"database"`
	Hash     HashKeyConfig  `json:"hash_key"`
	NewRelic NewRelicConfig `json:"new_relic"`
	Redis    RedisConfig    `json:"redis"`

	// non-secret config
	Timeout TimeoutConfig `yaml:"timeout"`
	TTL     TTLConfig     `yaml:"ttl"`
}

/*
*
// SECRET CONFIGS
*/
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
	JWT      string `json:"jwt"`
}

type NewRelicConfig struct {
	UserKey    string `json:"user_key"`
	LicenseKey string `json:"license_key"`
}

type RedisConfig struct {
	Address  string `json:"address"`
	Password string `json:"password"`
}

/**
// NON-SECRET CONFIGS
*/

type TimeoutConfig struct {
	FiveSeconds int `yaml:"five_seconds"`
}

type TTLConfig struct {
	FifteenMinutes int `yaml:"fifteen_minutes"`
	OneDay         int `yaml:"one_day"`
}

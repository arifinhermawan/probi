package configuration

type AppConfig struct {
	// secret config
	Database DatabaseConfig `json:"database"`
	Hash     HashKeyConfig  `json:"hash_key"`
	NewRelic NewRelicConfig `json:"new_relic"`
	NSQ      NSQConfig      `json:"nsq"`
	Redis    RedisConfig    `json:"redis"`

	// non-secret config
	Channel         Channel               `yaml:"channel"`
	Cron            CronConfig            `yaml:"cron"`
	PublishReminder PublishReminderConfig `yaml:"publish_reminder"`
	Timeout         TimeoutConfig         `yaml:"timeout"`
	TTL             TTLConfig             `yaml:"ttl"`
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

type NSQConfig struct {
	NSQD    string   `json:"nsqd"`
	Lookupd []string `json:"lookupd"`
}

type RedisConfig struct {
	Address  string `json:"address"`
	Password string `json:"password"`
}

/*
*
// NON-SECRET CONFIGS
*/

type Channel struct {
	Reminder ChannelConfig `yaml:"reminder"`
}

type ChannelConfig struct {
	Topic       string `yaml:"topic"`
	Channel     string `yaml:"channel"`
	MaxAttempt  int    `yaml:"max_attempt"`
	MaxInFlight int    `yaml:"max_in_flight"`
}

type CronConfig struct {
	ProcessDueReminder string `yaml:"process_due_reminder"`
}

type PublishReminderConfig struct {
	BatchSize int `yaml:"batch_size"`
}

type TimeoutConfig struct {
	FiveSeconds int `yaml:"five_seconds"`
}

type TTLConfig struct {
	FiveMinutes    int `yaml:"five_minutes"`
	FifteenMinutes int `yaml:"fifteen_minutes"`
	OneDay         int `yaml:"one_day"`
}

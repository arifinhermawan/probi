package configuration

type AppConfig struct {
	// secret config
	Database DatabaseConfig `json:"database"`
	Hash     HashKeyConfig  `json:"hash_key"`
	NewRelic NewRelicConfig `json:"new_relic"`
	Redis    RedisConfig    `json:"redis"`
	RMQ      RMQConfig      `json:"rabbit_mq"`

	// non-secret config
	Channel         ChannelConfig         `yaml:"channel"`
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

type RMQConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

type NewRelicConfig struct {
	UserKey    string `json:"user_key"`
	LicenseKey string `json:"license_key"`
}

type RedisConfig struct {
	Address  string `json:"address"`
	Password string `json:"password"`
}

/*
*
// NON-SECRET CONFIGS
*/

type ChannelConfig struct {
	Exchange Exchange `yaml:"exchange"`
	Queue    Queue    `yaml:"queue"`
}

type CronConfig struct {
	ProcessDueReminder string `yaml:"process_due_reminder"`
}

type Exchange struct {
	Reminder ExchangeConfig `yaml:"reminder"`
}

type ExchangeConfig struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

type Queue struct {
	ReminderSendEmail QueueConfig `yaml:"reminder_send_email"`
}

type QueueConfig struct {
	Name       string `yaml:"name"`
	RoutingKey string `yaml:"routing_key"`
	Exchange   string `yaml:"exchange"`
	Consumer   string `yaml:"consumer"`
}

type ReminderConfig struct {
	Exchange    string                    `yaml:"exchange"`
	Queue       ReminderQueueConfig       `yaml:"queue"`
	RoutingKeys ReminderRoutingKeysConfig `yaml:"routing_key"`
	MaxRetry    int                       `yaml:"max_retry"`
}

type ReminderQueueConfig struct {
	Email string `yaml:"email"`
}

type ReminderRoutingKeysConfig struct {
	Email string `yaml:"email"`
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

package config

// Config structure required
type Config struct {
	TelegramToken   string
	TelegramMessage string
	TelegramChatID  string
}

// NewConfig to be returned
func NewConfig() *Config {
	return &Config{}
}

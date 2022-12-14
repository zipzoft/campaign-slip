package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port            int    `mapstructure:"APP_PORT"`
	DBName          string `mapstructure:"DB_NAME"`
	DBUri           string `mapstructure:"DB_URI"`
	Env             string `mapstructure:"ENV"`
	CMSUrl          string `mapstructure:"CMS_URL"`
	AMMBOUrl        string `mapstructure:"AMMBO_URL"`
	WalletAPI       string `mapstructure:"WALLET_API"`
	WalletSettingID string `mapstructure:"WALLET_SETTING_ID"`
	CUSTOMER        string `mapstructure:"CUSTOMER"`
	DEP             string `mapstructure:"DEP"`
	GoogleRecaptcha struct {
		Enabled bool   `mapstructure:"GOOGLE_RECAPTCHA_ENABLED"`
		Secret  string `mapstructure:"GOOGLE_RECAPTCHA_SECRET"`
	}
	IncreaseTime int `mapstructure:"INCREASE_TIME"`
}
type DatabaseConfiguration struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

var _config *Config

func init() {
	if _config == nil {
		_config = New()
	}
}

func New() *Config {
	var cfg Config

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No env file", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Println("Error type unmarshal", err)
	}
	// Map google recaptcha environment variables
	cfg.GoogleRecaptcha.Enabled = viper.GetBool("GOOGLE_RECAPTCHA_ENABLED")
	cfg.GoogleRecaptcha.Secret = viper.GetString("GOOGLE_RECAPTCHA_SECRET")

	return &cfg
}

func GetConfig() *Config {
	return _config
}

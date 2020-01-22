package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// Config all global object that holds all application level variables
var Config appConfig

type appConfig struct {
	// the share DB ORM object
	DB *gorm.DB

	// the error thrown be GORM when using DB ORM object
	DBErr error

	// the server port. Defaults to 8080
	ServerPort int `mapstructure:"server_port"`

	// the data source name (DNS) for connnecting to the database. required
	DNS string `mapstructure:"dns"`

	// the API key needed to authorize to API. required
	ApiKey string `mapstructure:"api_key"`

	// certificate file for HTTPS
	CertFile string `mapstructure:"cert_file"`

	// Private key file HTTPS
	KeyFile string `mapstructure:"key_file"`
}

// LoadConfig loads config from files
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("example")
	v.SetConfigType("yaml")
	// v.SetEnvPrefix("server")
	v.AddConfigPath(".")
	v.AutomaticEnv()

	// Config.DNS = v.Get("DNS").(string)
	Config.ApiKey = v.Get("API_KEY").(string)
	v.SetDefault("server_port", 1234)

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read the configuation file : %s", err)
	}

	return v.Unmarshal(&Config)
}

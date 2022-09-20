package configuration

import (
	"fmt"
	"github.com/spf13/viper"
)

var Config appConfig

type appConfig struct {
	ConfigVar string
}

func LoadConfig(configPath string, propertyName string) (error, string) {
	v := viper.New()
	v.SetConfigName("properties")
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	v.AddConfigPath("properties.yaml")

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err), ""
	}
	return v.Unmarshal(&Config), v.GetString(propertyName)
}

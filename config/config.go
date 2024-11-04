package config

import (
	"os"
	"regexp"

	"github.com/ritik-patel05/image-player/internal/constants"
	"github.com/spf13/viper"
)

var config *Store

func init() {
	config = newConfig()
}

func newConfig() *Store {
	serviceName := regexp.MustCompile(`^(.*` + constants.SERVICE_NAME + `)`)
	workingDirectory, _ := os.Getwd()
	projectRoot := string(serviceName.Find([]byte(workingDirectory)))
	config, _ := loadEnvironmentConfig(projectRoot + "/assets/")
	return config
}

func GetConfig() *Store {
	return config
}

func loadEnvironmentConfig(configFilePath string) (*Store, error) {
	viper.AddConfigPath(configFilePath)
	activeEnv, ok := os.LookupEnv(constants.ACTIVE_ENV)
	if !ok {
		activeEnv = constants.STAGING
	}

	switch activeEnv {
	case constants.PRODUCTION:
		viper.SetConfigName(constants.PROD_CONFIG_FILE)
	case constants.DEV:
		viper.SetConfigName(constants.DEV_CONFIG_FILE)
	default:
		viper.SetConfigName(constants.STAG_CONFIG_FILE)
	}

	config := &Store{
		ActiveEnv: activeEnv,
	}

	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)

	return config, err
}

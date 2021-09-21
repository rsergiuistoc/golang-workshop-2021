package internal

import "github.com/spf13/viper"

// Configuration holds the configuration properties.
type Configuration struct {

	// Database config properties
	DBHost 				string 	`mapstructure:"SERVICE_DATABASE_HOST"`
	DBPort 				int		`mapstructure:"SERVICE_DATABASE_PORT"`
	DBName 				string	`mapstructure:"SERVICE_DATABASE_NAME"`
	DBUsername 			string	`mapstructure:"SERVICE_DATABASE_USER"`
	DBPassword 			string	`mapstructure:"SERVICE_DATABASE_PASSWORD"`
}

// InitConfiguration initialize Configuration
func InitConfiguration(filename, path string) *Configuration{

	viper.SetConfigName(filename)
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var config Configuration
	err = viper.Unmarshal(&config)
	if err != nil{
		panic(err)
	}

	return &config
}
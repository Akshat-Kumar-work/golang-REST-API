package config

import (
	"log"
	"os"
	"strings"

	"github.com/Akshat-Kumar-work/golang-rest-api/pkg/logger"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Using mapstructure (via Viper) to unmarshal YAML config into a typed Go struct.
// Viper reads config from file (and optionally environment variables), then mapstructure
// maps the keys to the corresponding struct fields using `mapstructure` tags.

type HTTPServer struct {
	Address string `mapstructure:"address"`
}

type Config struct {
	Env              string `mapstructure:"env"`
	StoragePath      string `mapstructure:"storage_path"`
	HTTPServer       `mapstructure:"http_server"`
	logger.LogConfig `mapstructure:"log_config"`
}

func LoadConfig() *Config {

	// it first loaded .env file and check for variables from that if not found fall-back to config file
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env file not found, using OS environment variables only")
	}

	//get env value from .env if available if not available use local
	env := os.Getenv("ENV")
	if env == "" {
		viper.SetConfigName("local") //config file name without extension
	}
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	viper.AutomaticEnv()                                   //it automatically loads env var, whenever we say viper.GetString("app.Env") looks in env var if not found look into config file
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) //for config file we can use app.HTTPServer.Address but not in env var, so it replace . with _

	//read config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg
}

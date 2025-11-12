package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string
}

type Config struct {
	Env         string `yaml:"env"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:http_server`
}

func MustLoad() *Config {

	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "Path to config file for the configuration") // this is for to check if the flags are passed like go run main.go 	--config-path ;
		flag.Parse()
		// Define flags first (flag.String), then call flag.Parse() to read and assign their values from command-line args.
		// Reference this chat: https://chatgpt.com/c/691388eb-10b4-8324-ae6b-c34e90d6c87f

		configPath = *flags

		if configPath == "" {
			log.Fatal("Unable to get config file path")
		}
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("Config file doesn't exist on this path %s", configPath)
	}
	var config Config
	err := cleanenv.ReadConfig(configPath, &config)

	if err != nil {
		log.Fatalf("Can't read config file :: %s", err)
	}
	return &config
}

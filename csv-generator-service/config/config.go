package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Facebook Facebook `json:"facebook,required"`
}

type Facebook struct {
	Db db `json:"db,required"`
}

type db struct {
	Url string `json:"url,required"`
}

func GetServiceConfig() Config {
	log.Println("loading configuration...")
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("error getting working directory: %v", err)
	}
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(fmt.Sprintf("%s/config", workDir))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error getting service configuration: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("error unmarshalling service configuration: %v", err)
	}

	log.Println("creating output folder...")
	if err := os.MkdirAll("./output", os.ModePerm); err != nil {
		log.Fatalf("error creating folder: %v", err)
	}

	return config
}

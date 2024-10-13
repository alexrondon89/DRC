package config

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/viper"
)

type Config struct {
	Facebook Facebook `json:"facebook,required"`
}

type Facebook struct {
	BaseUrl             string `json:"baseUrl,required"`
	AccessToken         string `json:"accessToken,required"`
	MaxPagesForGroups   uint   `json:"maxPagesForGroups,required"`
	MaxPagesForPosts    uint   `json:"maxPagesForPosts,required"`
	MaxPagesForComments uint   `json:"maxPagesForComments,required"`
	Db                  db     `json:"db,required"`
}

type db struct {
	Url string `json:"url,required"`
}

func GetServiceConfig() Config {
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
	log.Println("loading configuration...")
	return config
}

func ExecDbMigrator() {
	log.Println("executing flyway scripts...")
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("error getting working directory: %v", err)
	}
	cmd := exec.Command("flyway", "migrate", fmt.Sprintf("-configFiles=%s/flyway/flyway.conf", workDir))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("error migrating database: %s ,%v", output, err)
	}
	log.Println(string(output))
}

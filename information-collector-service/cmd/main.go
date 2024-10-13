package main

import (
	"log"
	"net/http"

	"github.com/alexrondon89/DRC/information-collector-service/config"
	"github.com/alexrondon89/DRC/information-collector-service/internal/bll"
	"github.com/alexrondon89/DRC/information-collector-service/internal/bll/dal/facebook"
)

func main() {
	log.Println("starting information-collector-service...")
	srvConfig := config.GetServiceConfig()
	config.ExecDbMigrator()
	httpCli := &http.Client{}
	//faceCollector := facebook.NewFacebookCollector(httpCli, srvConfig.Facebook)
	faceCollector := facebook.NewMockFacebookCollector(httpCli, srvConfig.Facebook)
	faceBll := bll.NewProcessor(faceCollector, srvConfig.Facebook)
	err := faceBll.GetFacebookInformation()
	if err != nil {
		log.Println("service information-collector-service failed. [error] ", err)
	}
	faceCollector.Close()
	log.Println("service information-collector-service finished successfully...")
}

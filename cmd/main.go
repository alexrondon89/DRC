package main

import (
	"fmt"
	"github.com/alexrondon89/DRC/config"
	"github.com/alexrondon89/DRC/internal/bll"
	"github.com/alexrondon89/DRC/internal/bll/dal/facebook"
	"net/http"
)

func main() {
	fmt.Println("starting DRC service to collect data...")
	srvConfig := config.GetServiceConfig()
	config.ExecDbMigrator()
	httpCli := &http.Client{}
	//faceCollector := facebook.NewFacebookCollector(httpCli, srvConfig.Facebook)
	faceCollector := facebook.NewMockFacebookCollector(httpCli, srvConfig.Facebook)
	faceBll := bll.NewProcessor(faceCollector, srvConfig.Facebook)
	err := faceBll.GetFacebookInformation()
	if err != nil {
		fmt.Println("service DRC failed. [error] ", err)
	}
	faceCollector.Close()
	fmt.Println("service DRC finished successfully...")
}

package main

import (
	"log"
	"os"

	"github.com/alexrondon89/DRC/csv-generator-service/config"
	"github.com/alexrondon89/DRC/csv-generator-service/internal"
)

func main() {
	log.Println("starting csv-generator-service to build csv...")
	srvConfig := config.GetServiceConfig()
	dbCli := internal.NewDbClient(srvConfig)
	option := os.Getenv("TYPE")
	switch option {
	case "full_information":
		log.Println("full_information not implemented yet")

	case "groups_information":
		err := dbCli.GetGroups()
		if err != nil {
			log.Fatalf("groups_information execution failed, [error] %v", err)
		}
		log.Println("groups_information executed successfully...")

	case "posts_information":
		err := dbCli.GetPosts()
		if err != nil {
			log.Fatalf("posts_information execution failed, [error] %v", err)
		}
		log.Println("posts_information executed successfully...")

	case "comments_information":
		err := dbCli.GetComments()
		if err != nil {
			log.Fatalf("comments_information execution failed, [error] %v", err)
		}
		log.Println("comments_information executed successfully...")

	default:
		log.Fatalf("%s is not a valid option", option)
	}
}

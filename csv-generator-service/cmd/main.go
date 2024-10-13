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
	case "groups_posts_comments":
		err := dbCli.GetGroupsPostsAndComments()
		if err != nil {
			log.Fatalf("groups_posts_comments command execution failed, [error] %v", err)
		}
		log.Println("groups_posts_comments command executed successfully...")

	case "groups":
		err := dbCli.GetGroups()
		if err != nil {
			log.Fatalf("groups command execution failed, [error] %v", err)
		}
		log.Println("groups command executed successfully...")

	case "posts":
		err := dbCli.GetPosts()
		if err != nil {
			log.Fatalf("posts command execution failed, [error] %v", err)
		}
		log.Println("posts command executed successfully...")

	case "comments":
		err := dbCli.GetComments()
		if err != nil {
			log.Fatalf("comments command execution failed, [error] %v", err)
		}
		log.Println("comments command executed successfully...")

	default:
		log.Fatalf("%s is not a valid option", option)
	}
}

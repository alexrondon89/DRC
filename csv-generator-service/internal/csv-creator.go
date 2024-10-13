package internal

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"time"

	"github.com/alexrondon89/DRC/csv-generator-service/config"
	"github.com/jackc/pgx/v4"
)

type DbClient struct {
	dbCli *pgx.Conn
}

func NewDbClient(config config.Config) DbClient {
	conn, err := pgx.Connect(context.Background(), config.Facebook.Db.Url)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	return DbClient{
		dbCli: conn,
	}
}

func (cli DbClient) Close() {
	err := cli.dbCli.Close(context.Background())
	if err != nil {
		log.Fatalf("unable to close connection with database: %v", err)
	}
}

func (cli DbClient) GetGroups() error {
	query := `
		SELECT 
			id,
			name,
			description,
			privacy,
			origin,
			updated_time
		FROM 
			meta.groups
		ORDER BY 
			updated_time desc
	`
	rows, err := cli.dbCli.Query(context.Background(), query)
	if err != nil {
		return err
	}

	file, err := os.Create("groups.csv")
	writer := csv.NewWriter(file)
	err = writer.Write([]string{"id", "name", "description", "privacy", "origin", "updated_time"})
	if err != nil {
		return err
	}
	for rows.Next() {
		var id, name, description, privacy, origin string
		var updatedTime time.Time
		if err := rows.Scan(&id, &name, &description, &privacy, &origin, &updatedTime); err != nil {
			return err
		}
		updatedTimeStr := updatedTime.Format(time.RFC3339)
		writer.Write([]string{id, name, description, privacy, origin, updatedTimeStr})
	}
	if err := rows.Err(); err != nil {
		return err
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
	}
	log.Println("file groups.csv created successfully....")
	return nil
}

func (cli DbClient) GetPosts() error {
	query := `
		SELECT 
			id,
			group_id,
			message,
			user_id,
			origin,
			created_at
		FROM 
			meta.posts
		ORDER BY 
			created_at desc
	`
	rows, err := cli.dbCli.Query(context.Background(), query)
	if err != nil {
		return err
	}

	file, err := os.Create("posts.csv")
	writer := csv.NewWriter(file)
	err = writer.Write([]string{"id", "group_id", "message", "user_id", "origin", "created_at"})
	if err != nil {
		return err
	}
	for rows.Next() {
		var id, groupId, message, userId, origin string
		var createdAt time.Time
		if err := rows.Scan(&id, &groupId, &message, &userId, &origin, &createdAt); err != nil {
			return err
		}
		createdAtStr := createdAt.Format(time.RFC3339)
		writer.Write([]string{id, groupId, message, userId, origin, createdAtStr})
	}
	if err := rows.Err(); err != nil {
		return err
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
	}
	log.Println("file posts.csv created successfully....")
	return nil
}

func (cli DbClient) GetComments() error {
	query := `
		SELECT 
			id,
			post_id,
			message,
			user_id,
			origin,
			created_at
		FROM 
			meta.comments
		ORDER BY 
			created_at desc
	`
	rows, err := cli.dbCli.Query(context.Background(), query)
	if err != nil {
		return err
	}

	file, err := os.Create("comments.csv")
	writer := csv.NewWriter(file)
	err = writer.Write([]string{"id", "post_id", "message", "user_id", "origin", "created_at"})
	if err != nil {
		return err
	}
	for rows.Next() {
		var id, postId, message, userId, origin string
		var createdAt time.Time
		if err := rows.Scan(&id, &postId, &message, &userId, &origin, &createdAt); err != nil {
			return err
		}
		createdAtStr := createdAt.Format(time.RFC3339)
		writer.Write([]string{id, postId, message, userId, origin, createdAtStr})
	}
	if err := rows.Err(); err != nil {
		return err
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
	}
	log.Println("file posts.csv created successfully....")
	return nil
}

func (cli DbClient) GetFullInfo() {

}

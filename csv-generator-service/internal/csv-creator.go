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

func (cli DbClient) GetGroupsPostsAndComments() error {
	query := `
		SELECT 
			g.id AS group_id,
			g.name AS group_name,
			g.description AS group_description,
			g.privacy AS group_privacy,
			g.origin AS group_origin,
			g.updated_time AS group_updated_time,
			p.id AS post_id,
			p.message AS post_message,
			p.user_id AS post_user_id,
			p.origin AS post_origin,
			p.created_at AS post_created_at,
			c.id AS comment_id,
			c.message AS comment_message,
			c.user_id AS comment_user_id,
			c.origin AS comment_origin,
			c.created_at AS comment_created_at
		FROM 
			meta.groups g
		LEFT JOIN 
			meta.posts p ON g.id = p.group_id
		LEFT JOIN 
			meta.comments c ON p.id = c.post_id
	`
	rows, err := cli.dbCli.Query(context.Background(), query)
	if err != nil {
		return err
	}

	file, err := os.Create("groups_posts_comments.csv")
	writer := csv.NewWriter(file)
	err = writer.Write([]string{
		"group_id", "group_name", "group_description", "group_privacy", "group_origin", "group_updated_time",
		"post_id", "post_message", "post_user_id", "post_origin", "post_created_at",
		"comment_id", "comment_message", "comment_user_id", "comment_origin", "comment_created_at"})
	if err != nil {
		return err
	}

	for rows.Next() {
		var groupId, groupName, groupDescription, groupPrivacy, groupOrigin string
		var postId, postMessage, postUserId, postOrigin string
		var commentId, commentMessage, commentUserId, commentOrigin string
		var groupUpdatedTime, postCreatedAt, commentCreatedAt time.Time
		if err := rows.Scan(
			&groupId, &groupName, &groupDescription, &groupPrivacy, &groupOrigin, &groupUpdatedTime,
			&postId, &postMessage, &postUserId, &postOrigin, &postCreatedAt,
			&commentId, &commentMessage, &commentUserId, &commentOrigin, &commentCreatedAt,
		); err != nil {
			return err
		}
		groupUpdatedTimeStr := groupUpdatedTime.Format(time.RFC3339)
		postCreatedAtStr := postCreatedAt.Format(time.RFC3339)
		commentCreatedAtStr := commentCreatedAt.Format(time.RFC3339)
		writer.Write([]string{
			groupId, groupName, groupDescription, groupPrivacy, groupOrigin, groupUpdatedTimeStr,
			postId, postMessage, postUserId, postOrigin, postCreatedAtStr,
			commentId, commentMessage, commentUserId, commentOrigin, commentCreatedAtStr,
		})
	}
	if err := rows.Err(); err != nil {
		return err
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
	}
	log.Println("file groups_posts_comments.csv created successfully....")
	return nil
}

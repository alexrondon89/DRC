package facebook

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/alexrondon89/DRC/information-collector-service/config"
	"github.com/alexrondon89/DRC/information-collector-service/internal/bll/dal/facebook/models"
	"github.com/alexrondon89/DRC/information-collector-service/pkg"
	"github.com/jackc/pgx/v4"
)

type FacebookCollector struct {
	http   *http.Client
	config config.Facebook
	dbCli  *pgx.Conn
}

func NewFacebookCollector(http *http.Client, config config.Facebook) FacebookCollector {
	log.Println("creating facebook collector client")
	conn, err := pgx.Connect(context.Background(), config.Db.Url)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	return FacebookCollector{
		http:   http,
		config: config,
		dbCli:  conn,
	}
}

func (faceColl FacebookCollector) GetUserInfo() (models.User, error) {
	url := fmt.Sprintf("%s/me?access_token=%s", faceColl.config.BaseUrl, faceColl.config.AccessToken)
	log.Println("requesting user info with url ", url)
	var userInfo models.User
	err := pkg.ExecHttp(faceColl.http, nil, "GET", url, &userInfo)
	if err != nil {
		return userInfo, err
	}
	return userInfo, nil
}

func (faceColl FacebookCollector) GetUserGroups(userId string, url string) (models.Groups, error) {
	if url == "" {
		url = fmt.Sprintf("%s/%s/groups?access_token=%s", faceColl.config.BaseUrl, userId, faceColl.config.AccessToken)
	}
	log.Println("requesting user groups with url ", url)
	var groupsInfo models.Groups
	err := pkg.ExecHttp(faceColl.http, nil, "GET", url, &groupsInfo)
	if err != nil {
		return groupsInfo, err
	}
	return groupsInfo, nil
}

func (faceColl FacebookCollector) GetGroupPosts(groupId string, url string) (models.Posts, error) {
	if url == "" {
		url = fmt.Sprintf("%s/%s/feed/?access_token=%s", faceColl.config.BaseUrl, groupId, faceColl.config.AccessToken)
	}
	log.Println("requesting group's posts with url ", url)
	var postsInfo models.Posts
	err := pkg.ExecHttp(faceColl.http, nil, "GET", url, &postsInfo)
	if err != nil {
		return postsInfo, err
	}
	return postsInfo, nil
}

func (faceColl FacebookCollector) GetPostComments(postId string, url string) (models.Comments, error) {
	if url == "" {
		url = fmt.Sprintf("%s/%s/comments?access_token=%s", faceColl.config.BaseUrl, postId, faceColl.config.AccessToken)
	}
	log.Println("requesting post's comments with url ", url)
	var commentsInfo models.Comments
	err := pkg.ExecHttp(faceColl.http, nil, "GET", url, &commentsInfo)
	if err != nil {
		return commentsInfo, err
	}
	return commentsInfo, nil
}

func (faceColl FacebookCollector) SaveUserGroups(groups []models.Group) error {
	tx, err := faceColl.dbCli.Begin(context.Background())
	if err != nil {
		log.Printf("failed to begin transaction: %v", err)
		return err
	}

	defer tx.Rollback(context.Background())

	query := `INSERT INTO meta.groups (
		id,
		name,
		description,
		privacy,
		origin,
		updated_time
	) VALUES (
		$1, $2, $3, $4, $5, $6
	) ON CONFLICT (id) DO NOTHING`

	for _, group := range groups {
		_, err := tx.Exec(
			context.Background(),
			query,
			group.ID,
			group.Name,
			group.Description,
			group.Privacy,
			"facebook",
			group.UpdatedTime,
		)
		if err != nil {
			log.Printf("failed to execute insert: %v", err)
			return err
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		log.Printf("failed to commit transaction: %v", err)
		return err
	}

	return nil
}

func (faceColl FacebookCollector) SaveGroupPosts(posts []models.Post, groupId string) error {
	tx, err := faceColl.dbCli.Begin(context.Background())
	if err != nil {
		log.Printf("failed to begin transaction: %v", err)
		return err
	}

	defer tx.Rollback(context.Background())

	query := `INSERT INTO meta.posts (
		id,
		group_id,
		message,
		user_id,
		origin,
		created_at
	) VALUES (
		$1, $2, $3, $4, $5, $6
	) ON CONFLICT (id) DO NOTHING`

	for _, post := range posts {
		_, err := tx.Exec(
			context.Background(),
			query,
			post.ID,
			groupId,
			post.Message,
			post.From.ID,
			"facebook",
			post.CreatedAt,
		)
		if err != nil {
			log.Printf("failed to execute insert: %v", err)
			return err
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		log.Printf("failed to commit transaction: %v", err)
		return err
	}

	return nil
}

func (faceColl FacebookCollector) SavePostComments(comments []models.Comment, postId string) error {
	tx, err := faceColl.dbCli.Begin(context.Background())
	if err != nil {
		log.Printf("failed to begin transaction: %v", err)
		return err
	}

	defer tx.Rollback(context.Background())

	query := `INSERT INTO meta.comments (
		id,
		post_id,
		message,
		user_id,
		origin,
		created_at
	) VALUES (
		$1, $2, $3, $4, $5, $6
	) ON CONFLICT (id) DO NOTHING`

	for _, comment := range comments {
		_, err := tx.Exec(
			context.Background(),
			query,
			comment.ID,
			postId,
			comment.Message,
			comment.From.ID,
			"facebook",
			comment.CreatedAt,
		)
		if err != nil {
			log.Printf("failed to execute insert: %v", err)
			return err
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		log.Printf("failed to commit transaction: %v", err)
		return err
	}

	return nil
}

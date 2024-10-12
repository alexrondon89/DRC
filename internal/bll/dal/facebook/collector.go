package facebook

import (
	"context"
	"fmt"
	"github.com/alexrondon89/DRC/config"
	"github.com/alexrondon89/DRC/internal/bll/dal/facebook/models"
	"github.com/alexrondon89/DRC/pkg"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
)

type FacebookCollector struct {
	http   *http.Client
	config config.Facebook
	dbCli  *pgx.Conn
}

func NewFacebookCollector(http *http.Client, config config.Facebook) FacebookCollector {
	fmt.Println("creating facebook collector client")
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
	fmt.Println("requesting user info with url ", url)
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
	fmt.Println("requesting user groups with url ", url)
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
	fmt.Println("requesting group's posts with url ", url)
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
	fmt.Println("requesting post's comments with url ", url)
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
		return fmt.Errorf("failed to begin transaction: %v", err)
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
			return fmt.Errorf("failed to execute insert: %v", err)
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (faceColl FacebookCollector) SaveGroupPosts(posts []models.Post, groupId string) error {
	tx, err := faceColl.dbCli.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
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
			return fmt.Errorf("failed to execute insert: %v", err)
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (faceColl FacebookCollector) SavePostComments(comments []models.Comment, postId string) error {
	tx, err := faceColl.dbCli.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
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
			return fmt.Errorf("failed to execute insert: %v", err)
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

package facebook

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"math/big"
	"net/http"

	"github.com/alexrondon89/DRC/information-collector-service/config"
	"github.com/alexrondon89/DRC/information-collector-service/internal/bll/dal/facebook/models"
)

type MockFacebookCollector struct {
	dbCli *pgx.Conn
}

func NewMockFacebookCollector(http *http.Client, config config.Facebook) MockFacebookCollector {
	conn, err := pgx.Connect(context.Background(), config.Db.Url)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	return MockFacebookCollector{
		dbCli: conn,
	}
}

func (mockColl MockFacebookCollector) Close() {
	err := mockColl.dbCli.Close(context.Background())
	if err != nil {
		log.Fatalf("unable to close connection with database: %v", err)
	}
}

func (mockColl MockFacebookCollector) GetUserInfo() (models.User, error) {
	return models.User{
		Name: "Pedro luis",
		ID:   "19592130",
	}, nil
}

func (mockColl MockFacebookCollector) GetUserGroups(userId string, url string) (models.Groups, error) {
	firstId, _ := rand.Int(rand.Reader, big.NewInt(1000))
	secondId, _ := rand.Int(rand.Reader, big.NewInt(1000))
	return models.Groups{
		Data: []models.Group{
			{
				ID:          fmt.Sprintf("%d", firstId),
				Name:        "delta group",
				Description: "this is the delta group",
				Privacy:     "CLOSED",
				UpdatedTime: "2024-09-29T14:53:00+0000",
			},
			{
				ID:          fmt.Sprintf("%d", secondId),
				Name:        "alfa group",
				Description: "this is the alfa group",
				Privacy:     "PUBLIC",
				UpdatedTime: "2024-09-22T12:33:00+0000",
			},
		},
		Paging: models.Paging{
			Previous: "https://graph.facebook.com/v2.11/123456789/groups?since=2024-09-29T14%3A53%3A00%2B0000",
			Next:     "https://graph.facebook.com/v2.11/123456789/groups?until=2024-10-01T10%3A22%3A00%2B0000",
		},
	}, nil
}

func (mockColl MockFacebookCollector) GetGroupPosts(groupId string, url string) (models.Posts, error) {
	firstId, _ := rand.Int(rand.Reader, big.NewInt(1000))
	secondId, _ := rand.Int(rand.Reader, big.NewInt(1000))
	return models.Posts{
		Data: []models.Post{
			{
				ID:        fmt.Sprintf("%d", firstId),
				Message:   "This is the first post!",
				CreatedAt: "2024-09-29T15:45:00+0000",
				From: models.User{
					ID:   "19592130",
					Name: "Carlos Andres",
				},
			},
			{
				ID:        fmt.Sprintf("%d", secondId),
				Message:   "This is the second post!",
				CreatedAt: "2024-09-30T10:49:00+0000",
				From: models.User{
					ID:   "19592130",
					Name: "Jesse Perez",
				},
			},
		},
		Paging: models.Paging{
			Previous: "https://graph.facebook.com/v2.11/123456789/posts?since=2024-09-29T12%3A00%3A00%2B0000",
			Next:     "https://graph.facebook.com/v2.11/123456789/posts?until=2024-09-30T15%3A45%3A00%2B0000",
		},
	}, nil
}

func (mockColl MockFacebookCollector) GetPostComments(postId string, url string) (models.Comments, error) {
	firstId, _ := rand.Int(rand.Reader, big.NewInt(1000))
	secondId, _ := rand.Int(rand.Reader, big.NewInt(1000))

	return models.Comments{
		Data: []models.Comment{
			{
				ID:        fmt.Sprintf("%d", firstId),
				Message:   "This is the first comment!",
				CreatedAt: "2024-10-02T13:45:00+0000",
				From: models.User{
					ID:   "987654321",
					Name: "John Doe",
				},
			},
			{
				ID:        fmt.Sprintf("%d", secondId),
				Message:   "This is the second comment!",
				CreatedAt: "2024-10-02T14:30:00+0000",
				From: models.User{
					ID:   "543216789",
					Name: "Jane Smith",
				},
			},
		},
		Paging: models.Paging{
			Previous: "https://graph.facebook.com/v2.11/123456789/posts?since=2024-09-29T12%3A00%3A00%2B0000",
			Next:     "https://graph.facebook.com/v2.11/123456789/posts?until=2024-09-30T15%3A45%3A00%2B0000",
		},
	}, nil
}

func (mockColl MockFacebookCollector) SaveUserGroups(groups []models.Group) error {
	tx, err := mockColl.dbCli.Begin(context.Background())
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

func (mockColl MockFacebookCollector) SaveGroupPosts(posts []models.Post, groupId string) error {
	tx, err := mockColl.dbCli.Begin(context.Background())
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

func (mockColl MockFacebookCollector) SavePostComments(comments []models.Comment, postId string) error {
	tx, err := mockColl.dbCli.Begin(context.Background())
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

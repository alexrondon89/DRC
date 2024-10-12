package mock

import "github.com/alexrondon89/DRC/internal/bll/dal/facebook/models"

func GroupResponse() models.Groups {
	return models.Groups{
		Data: []models.Group{
			{
				ID:          "1",
				Name:        "First Group",
				Description: "this is the first group",
				Privacy:     "CLOSED",
				UpdatedTime: "2024-09-29T14:53:00+0000",
			},
			{
				ID:          "2",
				Name:        "Second Group",
				Description: "this is the second group",
				Privacy:     "OPEN",
				UpdatedTime: "2024-10-29T14:53:00+0000",
			},
		},
		Paging: models.Paging{
			Previous: "https://graph.facebook.com/v2.11/123456789/groups?since=2024-09-29T14%3A53%3A00%2B0000",
			Next:     "https://graph.facebook.com/v2.11/123456789/groups?until=2024-10-01T10%3A22%3A00%2B0000",
		},
	}
}

func GroupPostResponse() models.Posts {
	return models.Posts{
		Data: []models.Post{
			{
				ID:        "123456789_987654321",
				Message:   "Welcome to the first group!",
				CreatedAt: "2024-09-29T15:45:00+0000",
				From: models.User{
					ID:   "987654321",
					Name: "John Doe",
				},
			},
			{
				ID:        "123456789_654987321",
				Message:   "Welcome to the second group!",
				CreatedAt: "2024-09-30T10:49:00+0000",
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
	}
}

func CommentResponse() models.Comments {
	return models.Comments{
		Data: []models.Comment{
			{
				ID:        "cmt_001",
				Message:   "This is the first comment!",
				CreatedAt: "2024-10-02T13:45:00+0000",
				From: models.User{
					ID:   "987654321",
					Name: "John Doe",
				},
			},
			{
				ID:        "cmt_002",
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
	}
}

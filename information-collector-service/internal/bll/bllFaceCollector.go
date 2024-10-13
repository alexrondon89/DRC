package bll

import (
	"log"

	"github.com/alexrondon89/DRC/information-collector-service/config"
	"github.com/alexrondon89/DRC/information-collector-service/internal/bll/dal/facebook/models"
)

type FacebookInterface interface {
	GetUserInfo() (models.User, error)
	GetUserGroups(userId string, url string) (models.Groups, error)
	GetGroupPosts(groupId string, url string) (models.Posts, error)
	GetPostComments(postId string, url string) (models.Comments, error)
	SaveUserGroups(groups []models.Group) error
	SaveGroupPosts(posts []models.Post, groupId string) error
	SavePostComments(comments []models.Comment, postId string) error
}

type pageCounter struct {
	groupPagesCount   uint
	postPagesCount    uint
	commentPagesCount uint
}

type Processor struct {
	Facebook FacebookInterface
	counter  pageCounter
	Config   config.Facebook
}

func NewProcessor(facebook FacebookInterface, config config.Facebook) *Processor {
	return &Processor{
		Facebook: facebook,
		Config:   config,
		counter: pageCounter{
			groupPagesCount:   1,
			postPagesCount:    1,
			commentPagesCount: 1,
		},
	}
}

func (pro *Processor) GetFacebookInformation() error {
	log.Println("getting facebook information...")
	respUserInfo, err := pro.Facebook.GetUserInfo()
	if err != nil {
		return err
	}

	log.Println("user root: ", respUserInfo)
	err = pro.getFacebookGroupsInformation(respUserInfo.ID, models.Paging{})
	if err != nil {
		return err
	}
	return nil
}

func (pro *Processor) getFacebookGroupsInformation(userId string, paging models.Paging) error {
	if pro.counter.groupPagesCount <= pro.Config.MaxPagesForGroups {
		log.Println("group information, page ", pro.counter.groupPagesCount)
		pro.counter.groupPagesCount += 1

		respUserGroups, err := pro.Facebook.GetUserGroups(userId, paging.Next)
		if err != nil {
			return err
		}

		err = pro.Facebook.SaveUserGroups(respUserGroups.Data)
		if err != nil {
			return err
		}

		for _, group := range respUserGroups.Data {
			err := pro.getFacebookPostsInformation(group.ID, models.Paging{})
			if err != nil {
				return err
			}
		}

		if respUserGroups.Paging.Next != "" {
			err = pro.getFacebookGroupsInformation(userId, respUserGroups.Paging)
			if err != nil {
				return err
			}
		}
	}

	pro.counter.groupPagesCount = 1
	return nil
}

func (pro *Processor) getFacebookPostsInformation(groupId string, paging models.Paging) error {
	if pro.counter.postPagesCount <= pro.Config.MaxPagesForPosts {
		log.Println("post information, page ", pro.counter.postPagesCount)
		pro.counter.postPagesCount += 1

		respGroupPosts, err := pro.Facebook.GetGroupPosts(groupId, paging.Next)
		if err != nil {
			return err
		}

		err = pro.Facebook.SaveGroupPosts(respGroupPosts.Data, groupId)
		if err != nil {
			return err
		}

		for _, post := range respGroupPosts.Data {
			err := pro.getFacebookCommentsInformation(post.ID, models.Paging{})
			if err != nil {
				return err
			}
		}

		if respGroupPosts.Paging.Next != "" {
			err = pro.getFacebookPostsInformation(groupId, respGroupPosts.Paging)
			if err != nil {
				return err
			}
		}

	}

	pro.counter.postPagesCount = 1
	return nil
}

func (pro *Processor) getFacebookCommentsInformation(postId string, paging models.Paging) error {
	if pro.counter.commentPagesCount <= pro.Config.MaxPagesForComments {
		log.Println("comment information, page ", pro.counter.commentPagesCount)
		pro.counter.commentPagesCount += 1

		respPostComments, err := pro.Facebook.GetPostComments(postId, paging.Next)
		if err != nil {
			return err
		}

		err = pro.Facebook.SavePostComments(respPostComments.Data, postId)
		if err != nil {
			return err
		}

		if respPostComments.Paging.Next != "" {
			err = pro.getFacebookCommentsInformation(postId, respPostComments.Paging)
			if err != nil {
				return err
			}
		}
	}

	pro.counter.commentPagesCount = 1
	return nil
}

package repositories

import (
	"go-restapi-boiltertemplate/app/models"
	"errors"
)

type PostRepository interface {
	FindAll() (*[]models.Post, error)
	FindById(postId int) (*models.Post, error)
	Create(post models.Post) (*models.Post, error)
}

type PostMemoryRepository struct {
	posts []models.Post
}

func NewPostMemoryRepository(posts []models.Post) *PostMemoryRepository {
	return &PostMemoryRepository{posts: posts}
}

func (r *PostMemoryRepository) FindAll() (*[]models.Post, error) {
	return &r.posts, nil
}

func (r *PostMemoryRepository) FindById(postId int) (*models.Post, error) {
	for _, v := range r.posts {
		if int(v.ID) == postId {
			return &v, nil
		}
	}
	return nil, errors.New("投稿が見つかりません")
}

func (r *PostMemoryRepository) Create(post models.Post) (*models.Post, error) {
	// 新しい投稿を追加
	r.posts = append(r.posts, post)
	return &post, nil
}

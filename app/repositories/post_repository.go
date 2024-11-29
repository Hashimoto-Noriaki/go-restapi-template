package repositories

import (
	"go-restapi-boiltertemplate/app/models"
	"errors"
)

type PostRepository interface {
	FindAll() (*[]models.Post, error)
	FindByID(postID uint) (*models.Post, error)
	Create(post models.Post) (*models.Post, error)
	Update(updatePost models.Post)(*models.Post,error)
	Delete(postID uint) error
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

func (r *PostMemoryRepository) FindByID(postID uint) (*models.Post, error) {
	for _, v := range r.posts {
		if uint(v.ID) == postID {
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

func (r *PostMemoryRepository) Update(updatePost models.Post)(*models.Post,error){
	for i,v := range r.posts {
		if v.ID == updatePost.ID {
			r.posts[i] = updatePost
			return &r.posts[i],nil
		}
	}
	return nil, errors.New("Unexpected error")
}

func (r *PostMemoryRepository) Delete(postID uint) error {
	for i,v := range r.posts {
		if v.ID == postID {
			r.posts = append(r.posts[:i], r.posts[i+1:]...)
			return nil
		}
	}
	return errors.New("Post not found")
}
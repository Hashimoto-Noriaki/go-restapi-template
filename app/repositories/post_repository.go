package repositories

import (
	"errors"
	"go-restapi-boiltertemplate/app/models";
)

type PostRepository interface {
	//repositoryが満たすべきメソッドを定義
	FindAll() (*[]models.Post,error)
	FindById(postId int) (*models.Post, error)
}

type PostMemoryRepository struct {
	//メモリ上のデータソースとして
	posts []models.Post
}

func NewPostMemoryRepository(posts []models.Post) PostRepository {
	return &PostMemoryRepository{posts: posts}
}

func (r *PostMemoryRepository) FindAll() (*[]models.Post,error){
	return &r.posts,nil
}

func (r *PostMemoryRepository) FindById(postId int) (*models.Post, error) {
	for _, v := range r.posts {
		if int(v.ID) == postId { // v.IDをintにキャストして比較
			return &v, nil
		}
	}
	return nil, errors.New("投稿がありません")
}
package repositories

import "go-restapi-boiltertemplate/app/models";

type PostRepository interface {
	//repositoryが満たすべきメソッドを定義
	FindAll() (*[]models.Post,error)
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
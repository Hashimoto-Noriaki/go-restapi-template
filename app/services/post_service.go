package services

import (
	"go-restapi-boiltertemplate/app/models"
	"go-restapi-boiltertemplate/app/repositories"
)

type PostService interface {
	FindAll() (*[]models.Post, error)
	FindById(postId int) (*models.Post, error)
}

type postService struct {
	repository repositories.PostRepository
}

func NewPostService(repository repositories.PostRepository) PostService {
	return &postService{repository: repository}
}

func (s *postService) FindAll() (*[]models.Post, error) {
	return s.repository.FindAll()
}

func (s *postService) FindById(postId int) (*models.Post, error) {
    return s.repository.FindById(postId)
}
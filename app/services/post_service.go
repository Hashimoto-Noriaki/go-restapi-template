package services

import (
	"go-restapi-boiltertemplate/app/models"
	"go-restapi-boiltertemplate/app/repositories"
	"go-restapi-boiltertemplate/app/dto"
)

type PostService interface {
	FindAll() (*[]models.Post, error)
	FindById(postId uint) (*models.Post, error)
	Create(createPostInput dto.CreatePostInput) (*models.Post, error)
	Update(postId uint, updatePostInput dto.UpdatePostInput) (*models.Post, error)
	Delete(postId uint) error
}

type postService struct {
	repository repositories.PostRepository
}

// NewPostServiceはPostServiceインターフェースを返す関数
func NewPostService(repository repositories.PostRepository) PostService {
	return &postService{repository: repository}
}

func (s *postService) FindAll() (*[]models.Post, error) {
	return s.repository.FindAll()
}

func (s *postService) FindById(postId uint) (*models.Post, error) {
	return s.repository.FindById(postId)
}

func (s *postService) Create(createPostInput dto.CreatePostInput) (*models.Post, error) {
	newPost := models.Post{
		Text:   createPostInput.Text,
		UserID: createPostInput.UserID,
	}
	return s.repository.Create(newPost)
}

// Updateメソッドの修正
func (s *postService) Update(postId uint, updatePostInput dto.UpdatePostInput) (*models.Post, error) {
	// postId を使って投稿を取得
	targetPost, err := s.FindById(postId)
	if err != nil {
		return nil, err
	}

	// Text の更新
	if updatePostInput.Text != "" {
		targetPost.Text = updatePostInput.Text
	}

	// UserID の更新（通常は変更しないが、もしリクエストに含まれていれば変更）
	if updatePostInput.UserID != 0 {
		targetPost.UserID = updatePostInput.UserID
	}

	// 更新された投稿を保存
	updatedPost, err := s.repository.Update(*targetPost) // ポインタをデリファレンスして値を渡す
	if err != nil {
		return nil, err
	}

	return updatedPost, nil
}

// Deleteメソッドの追加
func (s *postService) Delete(postId uint) error {
	return s.repository.Delete(postId) // 投稿を削除する
}

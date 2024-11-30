package repositories

import (
	"errors"
	"go-restapi-boiltertemplate/app/models"
	"gorm.io/gorm"
)

// PostRepositoryインターフェースの定義
type PostRepository interface {
	FindAll() (*[]models.Post, error)
	FindByID(postID uint) (*models.Post, error)
	Create(newPost models.Post) (*models.Post, error)
	Update(updatePost models.Post) (*models.Post, error)
	Delete(postID uint) error
}

// PostMemoryRepository構造体の定義
type PostMemoryRepository struct {
	db *gorm.DB
}

// NewPostRepository: リポジトリのコンストラクタ
func NewPostRepository(db *gorm.DB) PostRepository {
	return &PostMemoryRepository{db: db}
}

// FindAll: すべての投稿を取得
func (r *PostMemoryRepository) FindAll() (*[]models.Post, error) {
	var posts []models.Post
	result := r.db.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return &posts, nil
}

// FindByID: 指定されたIDの投稿を取得
func (r *PostMemoryRepository) FindByID(postID uint) (*models.Post, error) {
	var post models.Post
	result := r.db.First(&post, postID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("投稿が見つかりません")
		}
		return nil, result.Error
	}
	return &post, nil
}

// Create: 新しい投稿を作成
func (r *PostMemoryRepository) Create(newPost models.Post) (*models.Post, error) {
	result := r.db.Create(&newPost)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newPost, nil
}

// Update: 投稿を更新
func (r *PostMemoryRepository) Update(updatePost models.Post) (*models.Post, error) {
	result := r.db.Save(&updatePost)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updatePost, nil
}

// Delete: 指定されたIDの投稿を削除
func (r *PostMemoryRepository) Delete(postID uint) error {
	result := r.db.Delete(&models.Post{}, postID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("投稿が見つかりません")
		}
		return result.Error
	}
	return nil
}

package dto

type PostCreateDTO struct {
    Text   string `json:"text" binding:"required,min=1,max=255"`  // テキストは必須で1〜255文字
    UserID uint   `json:"user_id" binding:"required"`             // ユーザーIDは必須
}
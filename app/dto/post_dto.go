package dto

type CreatePostInput struct {
    Text   string `json:"text" binding:"required,min=1,max=255"`  // テキストは必須で1〜255文字
    UserID uint   `json:"user_id" binding:"required"`             // ユーザーIDは必須
}

type UpdatePostInput struct {
    Text   string `json:"text" binding:"omitempty,min=1,max=255"`  // テキストは任意で、更新時のみ1〜255文字
    UserID uint   `json:"user_id" binding:"omitempty"`             // ユーザーIDは任意
}
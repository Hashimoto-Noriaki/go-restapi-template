package dto

type UserCreateDTO struct {
    Name            string `json:"name" binding:"required,min=2,max=255"`// 名前は必須で、2〜255文字
    Email           string `json:"email" binding:"required,email"` // メールアドレスは必須で、有効なメール形式
    Password        string `json:"password" binding:"required,min=6,max=255"` // パスワードは必須で、6〜255文字
	PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password"` // パスワード確認が必要で、Passwordと一致する
}
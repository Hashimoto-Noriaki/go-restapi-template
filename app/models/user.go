package models

import (
    "time"
    "gorm.io/gorm"
)

//ユーザーテーブルの構造を定義するモデル
type User struct {
    ID              uint           `gorm:"primaryKey;autoIncrement"` // 主キー
    Name            string         `gorm:"type:varchar(255);not null"` // 名前
    Email           string         `gorm:"type:varchar(255);unique;not null"` // メールアドレス (一意制約付き)
    EmailVerifiedAt *time.Time     `gorm:"type:timestamp"`           // メールアドレス確認日時 (NULL許可)
    Password        string         `gorm:"type:varchar(255);not null"` // パスワード
    RememberToken   string         `gorm:"type:varchar(100)"`        // トークントークン
    CreatedAt       time.Time      // 作成日時
    UpdatedAt       time.Time      // 更新日時
    DeletedAt       gorm.DeletedAt `gorm:"index"`                    // ソフトデリート
}

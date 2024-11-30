package models

import (
    "time"
)

type Post struct {
    ID        uint           `gorm:"primaryKey;autoIncrement"` // 主キー
    Text      string         `gorm:"type:varchar(255);not null"` // テキスト内容
    UserID    uint           `gorm:"not null`           // ユーザーID（外部キー）
    CreatedAt time.Time      // 作成日時
    UpdatedAt time.Time      // 更新日時
    DeletedAt *time.Time     `gorm:"index"`                    // ソフトデリート用
}

# golang docker REST API のボイラーテンプレート

### 技術構成
- Go
- Gin
- GORM
- MySQL
- Adminer
- JWT
- golangci-lint
- GitHubActions
- Docker

### Docker
```
docker compose up -d
```
localhost:8080
- Adminer

### サーバー起動
```
go run main.go
```

### APIの動作確認
- Postmanを使用

- すべての投稿を取得 (GET /posts)
```
% curl http://localhost:8081/posts
```
- 特定の投稿を取得 (GET /posts/:id)
```
% curl http://localhost:8081/posts/1
```
- 投稿を作成 (POST /posts)
```
 % curl -X POST http://localhost:8081/posts \
-H "Content-Type: application/json" \
-d '{
  "Text": "3番目の投稿",
  "UserID": 3
}'
```

### GORM特徴
ロールバック、シーディング、カラム削除機能がない

### 依存関係を整理
```
go mod tidy
```

###  バージョン
```
$ go version
1.23.1
```

### マイグレーション実行
```
go run app/db/migrations/migration.go
```

### 静的解析ツール
- golangci-lint
- GitHubActions

main.goのrunをコメントアウト(基本linterを走らせる時はrunをコメントアウトする)
```
golangci-lint run
```

- golangciとGitHubActionsの資料

https://qiita.com/Hashimoto-Noriaki/items/e31dc42f2fe13f67c58c

### 必要なGoパッケージをインストール
```
go mod init todo-app
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u github.com/dgrijalva/jwt-go
```

### サーバー起動中にターミナルを閉じた時の対処
```
$ lsof -i :8080
COMMAND    PID             USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
com.docke 3444 hashimotonoriaki  171u  IPv6 0x8acba539313f8b73      0t0  TCP *:http-alt (LISTEN)
```
```
kill -9 3444 
```
3444はPID

### Linux/macOSで、opensslコマンドを使用してランダムなキーを生成
secret keyを生成
```
openssl rand -base64 32
```

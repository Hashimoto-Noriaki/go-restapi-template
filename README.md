# golang docker REST API のボイラーテンプレート

### 技術構成
- Go
- Gin
- GORM
- MySQL
- Adminer
- JWT

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
- golangci-int
- GitHubActions

### 必要なGoパッケージをインストール

```
go mod init todo-app
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u github.com/dgrijalva/jwt-go
```

### Linux/macOSで、opensslコマンドを使用してランダムなキーを生成
secret keyを生成
```
openssl rand -base64 32
```

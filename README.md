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

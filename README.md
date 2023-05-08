# Echo REST API

## use

golang(echo)
postgres
docker compose

## create go project
go mod init go-rest-api
go get gorm.io/gorm ...
<!-- or -->
go mod tidy


## migrate
マイグレーションの実行
migrate/migrate.goのmain関数を実行させるため下記を実行する
GO_ENVのみdb.goのNewDBでgodotenv.Load()を実行する前に処理をしているためここで指定
GO_ENV=dev go run migrate/migrate.go

## ハンズオンの環境構築
```zsh
cd none-attack // またはweak-secret
go run main.go
```

## エンドポイントへのアクセス方法
`curl http://localhost:5555/token`

## JWTをAuthorizationヘッダに入れてHTTPリクエストを送信する
`curl http://localhost:5555/admin -H "Authorization: Bearer <JWT>"`
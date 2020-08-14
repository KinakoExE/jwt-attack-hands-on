## ハンズオンの環境構築
```zsh
cd none-attack // またはweak-secret
go run main.go
```

## エンドポイントへのアクセス方法
- `curl http://localhost:5555/token`
- なお、どのハンズオンも`/`、`/token`, `/admin`の3つのエンドポイントしかありません


## JWTをAuthorizationヘッダに入れてHTTPリクエストを送信する
`curl http://localhost:5555/admin -H "Authorization: Bearer <JWT>"`

## その他
- none-attack内のjwt-goはこのハンズオンのために一部ライブラリ内のコードを変更しています
- .envファイルにはJWTの署名に用いる鍵を設定しています。weak-secretの.envを先に見てしまうとネタバレになってしまうので、ハンズオンが終わった後に開くことをお勧めします
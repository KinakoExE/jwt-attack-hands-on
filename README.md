## ハンズオンの環境構築

### local
```zsh
cd none-attack // またはbrute-force-secret
go get // 必要なライブラリをすべて入れる
go run main.go
```

John the Ripperの環境構築
```
git clone https://github.com/magnumripper/JohnTheRipper // John the Ripperの公式リポジトリからclone
cd JohnTheRipper/src
./configure
make -s clean && make -sj4 // ビルド

// JWTのsecretをブルートフォースしたい場合
cd JohnTheRipper/run
vim jwt.txt // JWTをテキストファイルに書く
./john jwt.txt // jwt.txtの中身はさっきのJWT
```

### docker
```zsh
# build and run for none-attack hands-on
make docker-none-attack

# build and run for none-attack hands-on
# if you want to use JohnTheRipper, you execute `docker exec -it <process id> sh` and use `./run/john jwt.txt`
make docker-brute-force-secret
```


## ハンズオンのゴール
- none-attack、brute-force-secret共に`"user":"admin"`として`/admin`にアクセスすることがゴールとなります

## エンドポイントへのアクセス方法
- `curl http://localhost:5555/token`
- なお、どのハンズオンも`/`,`/token`, `/admin`の3つのエンドポイントしかありません


## JWTをAuthorizationヘッダに入れてHTTPリクエストを送信する
`curl http://localhost:5555/admin -H "Authorization: Bearer <JWT>"`

## その他
- none-attack内のjwt-goはこのハンズオンのために一部ライブラリ内のコードを変更しています
- .envファイルにはJWTの署名に用いる鍵を設定しています。brute-force-secretの.envを先に見てしまうとネタバレになってしまうので、ハンズオンが終わった後に開くことをお勧めします

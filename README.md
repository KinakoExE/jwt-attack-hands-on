# Notification
**Docker上で動くようになりました🐳**
(Special Thanks👏 -> @kyoto7250)

---


## ハンズオンの環境構築

### local
```zsh
# またはbrute-force-secret
cd none-attack 

# 必要なライブラリをすべて入れる
go get 
go run main.go
```

John the Ripperの環境構築
```zsh
# John the Ripperの公式リポジトリからclone
git clone https://github.com/magnumripper/JohnTheRipper
cd JohnTheRipper/src
./configure
# ビルド
make -s clean && make -sj4

# JWTのsecretをブルートフォースしたい場合
cd JohnTheRipper/run

# JWTをテキストファイルに書く(Dockerの場合はviを利用してください)
vim jwt.txt

# jwt.txtの中身は総当たりしてsecretを調べたいJWT
./john jwt.txt
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
- ~~none-attack内のjwt-goはこのハンズオンのために一部ライブラリ内のコードを変更しています~~
  - **2022年現在、ライブラリの関係でsignature部分を削除しなくてもalgorithmをnoneに書き換えて動くようにしています**
- .envファイルにはJWTの署名に用いる鍵を設定しています。brute-force-secretの.envを先に見てしまうとネタバレになってしまうので、ハンズオンが終わった後に開くことをお勧めします

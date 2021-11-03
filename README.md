dockerで立ち上た時はそれぞれ下記が必要。

Go側
```
make bash-api

go run main.go
```

## Migrationコマンド

使用しているGooseのDoc
URL: https://bitbucket.org/liamstask/goose/src/master/

```
// 最後まで終わるまで待つようにする。
go install bitbucket.org/liamstask/goose/cmd/goose

goose create AddSomeColumns sql
goose: created db/migrations/20130106093224_AddSomeColumns.sql
```

Vue側
```
make bash-app

npm run dev
```

## JWTで使用する秘密鍵の生成方法のメモ
### pemキーの作成
```
openssl pkcs8 -topk8 -inform PEM -outform PEM -nocrypt -in private_key.pem -out pkcs8.key
```

### public keyのpem化
```
openssl rsa -pubout -in key.pem 
```
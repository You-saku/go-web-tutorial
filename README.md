# go-practice
 - Goのテスト環境(docker)
 - 開発しやすいようにAirを採用する
 - ライブラリはここで試そう

## 構成
 - golang 1.20
 - mysql 5.7

## セットアップ
```
$git clone git@github.com:You-saku/go-practice.git
$cd go-sample-app
$make setup
$make seeding
$make start
```

## セットアップ(n > 2)
```
$make up
```

## コンテナを止める
```
$make down
```

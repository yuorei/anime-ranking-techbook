# アニメランキング (技術書典14)

[Zli TechBook Vol.4](https://techbookfest.org/product/anxJ8MV0gX3F2nDTXeu601?productVariantID=hBG7it4HWiGhXZYUjsum3r) 第1章 Go と GraphQL によるバックエンド開発の参考リポジトリとなります。

基本的には[こちら](https://github.com/yuorei/anime-ranking)と同じです。

## 実行方法

envのコピー

```
$ cp .env.sample .env
```

mysqlをdockerで起動

```
docker run -d \
  --name anime_ranking_db \
  -e MYSQL_ROOT_PASSWORD=password \
  -e MYSQL_DATABASE=ANIME_RANKING \
  -p 3306:3306 \
  mysql:8.0 --default-time-zone='+00:00'
```

Goを実行

```
go run main.go
```

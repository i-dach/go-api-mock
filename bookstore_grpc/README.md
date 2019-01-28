# What's bookstore_mock ?
- このリポジトリは「Gin」フレームワークを利用したAPIの作成学習を目的とした本の管理システムのモックです
- 基本的なCURDができるようにはなっています

# Structure
```
├── Dockerfile				// API実行用 Docker image
├── api
│   ├── main.go   			// API Server
│   └── scope
│       ├── book.go			// URI「/book」以下の処理
│       └── router.go		// HTTP Request Method毎に呼び出すメソッドを振り分ける
├── docker-compose.yml
└── mysql
    ├── Dockerfile			// mysql image
    ├── db-data
    ├── initdb.d
    │   ├── 1_init.sh
    │   ├── 1_schema.sql_	// mysql schema
    │   └── 2_testdata.sql_	// 初期データ
    └── my.cnf
```

# Requirement
## 初回起動

```
-- 起動
$ docker-compose build
$ docker-compose up -d

-- 初期データ投入
$ docker exec -it bookstore_mock_db_1 bash
$ mysql -h 127.0.0.1 -u api -D api -papi < /docker-entrypoint-initdb.d/1_schema.sql_ 
```

## CURD
- 一覧取得

	```
	api_1  | [GIN-debug] GET    /book/list                --> _/api/scope.getBookList (3 handlers)
	```
- 単体取得（タイトルlike検索）

	```
	api_1  | [GIN-debug] GET    /book/search/:title       --> _/api/scope.getBookByTitle (3 handlers)
	```

- label更新（labelはhttp request bodyにて指定）
	```
	api_1  | [GIN-debug] PATCH  /book/:id                 --> _/api/scope.updateBookInfo (3 handlers)
	```

- 本の追加（title, labelをhttp request bodyにて指定）

	```
	api_1  | [GIN-debug] POST   /book/insert              --> _/api/scope.addBookInfo (3 handlers)
	```

- 本の削除（id指定）

	```
	api_1  | [GIN-debug] DELETE /book/:id                 --> _/api/scope.delBookInfo (3 handlers)
	```

## 動作検証

```
-- 検索
$ curl -X GET http://localhost:9000/book/list
$ curl -X GET http://localhost:9000/book/search/solr

-- 更新
$ curl -X PATCH -H 'content-type: application/json' -d '{"title":"RESTful web service","label":"architecture"}' http://localhost:9000/book/st.sol.18

-- 追加
$ curl -X POST -d "title=RESTful web service&label=architecture" http://localhost:9000/book/insert

-- 削除
curl -X DELETE  http://localhost:9000/book/ar.RE.1
```
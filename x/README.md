# What's this?
- Golang API for Web service.

# Target
- Used by micro service architecture.

# How to play?
1. install golang
1. setting $GOPATH
1. run process `docker build & run`
	```
	docker build -t api_tmp .
	docker run -d -p 8080:8080 api_tmp:latest
	```

1. you'll see "Hello World!" when you access to localhost:8080 it printing.

# How to development?

- If you need path "/bookmanage" URI, when you will create "bookmanege" directory into src directory.
- First, you should be created "hoge.doc".
- You must write "bookmanage" specification documentation to "hoge.doc".
- Break down the text written in "hoge.doc" and make it a list of function have return values.
- Sort functions in order of ease of creation.
- Next, Copy one of the created functions to "hoge_test.go".
- After that, develop according to TDD.

# Directory structure

```
.
├── Dockerfile				// Buildしたmain.exc ファイルしか
├── README.md
├── setup.sh				// $GOPATH以外で開発する時にでも$GOPATHからこのプロジェクトを参照されるようにするマジックアイテム
└── src
    ├── main.go				// serverを立ち上げるgo file
    ├── hoge				// URIベースで作成したディレクトリ
    │   ├── hoge.doc		// テストコードの元となる受け入れ条件を記載したドキュメント
    │   ├── hoge.go			// URIベースのコード
    │   └── hoge_test.go	// テストコード
    ├── ....
    └── foo
```

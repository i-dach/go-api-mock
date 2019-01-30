# What's template?
- target: micro service scope (not ECS's service)
- application specification  -> app/README.md
- middlewere specification  -> middle/README.md

# Directory structure

```
	.
	├── README.md 					// リポジトリはTask(Pod)単位で作成。スコープはサービス 
	├── setup.sh					// aws configure, ecs-cli, docker-composeとか開発に必要な環境のセットアップ
	├── Makefile					// ECS,ECRへのデプロイ、イメージ参照先の（ローカル - ECR）の切り替え、初期cluster作成など
	├── docker-comopose.yaml		// このサービスの依存関係を記す
	├── app							// Applicationのディレクトリ。サーバーサイドやクライアントで動くものはこっち
	│   ├── README.md 				// 仕様書。どんなURIがあってどんな機能なのか、呼び出し方、引数、メソッドを記す
	│   ├── Dockerfile				// アプリケーション用Docker image
	│   └── rest 					// RESTfull アーキテクチャで組んだAPI用のディレクトリ。 gRPCならgRPC, graphQLならgrqphQL
	│       ├── main.go 			// routerに引き渡すファイル。serverの受け口であり、golang実行用サーバーを立てる子。
	│       └── scope			
	│          └── router.go 		// 各エンドポイントにルーティングする子
	│          └── hoge.go 			// 各エンドポイントの処理用interface
	│          └── hoge_event.go	// 各エンドポイントの実際の処理
	│          └── hoge_test.go		// 各エンドポイントのテストコード
	├── middle
	│   ├── README.md 				// ミドルウェア用の仕様書。どんな要件を満たすのか、等を定義
	│   └── nginx
	│       └── Dockerfile			// 各ミドルウェア用Docker image。専用のカスタマイズがある場合のみ作成
	└── mock						// mock用のデータ置き場。api用のデータは基本JSONで記述
```
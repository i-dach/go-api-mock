# How to lib?
- logの出力
- zerologを利用したログを出力するテンプレート


# Log Level

| level | target |
|:---:|:---:|
| panic | 意図しない挙動で意図的に処理を処理を落とさないと他にも影響が出る場合 (logic) |
| fatal | 意図しない挙動で意図的に処理を処理を落とさないと他にも影響が出る場合 (server) |
| error | ユーザーが利用する情報のCURDや表示、操作が出来ない場合 |
| warn | (開発/検証中のみの利用) 処理の本筋には影響しないが、エラーになったもの (e.g. validateなど) |
| info | プロセスの開始や終了、CURDの情報などを残すためのもの |
| debug | (開発/検証中のみの利用)　プロセスの細かい挙動を表示 |

# Log output template

- logの出力は標準出力で行われる
- 標準出力内容は全レベルで統一

```
{
	"time"      : "Thu, 04 Oct 2014 23:59:45 GMT",							// RFC1123(JST)
	"level"     : "err",													// Log level
	"title"     : "[Exception] kubectl process can't action",				// Log title
	"status"    : "stop",													// Process status
	"exception" : "Unable to connect to the server: EOF",					// ERR detail
	"msg"	    : "perhaps... CPU upper over limit exceeded",				// Developer Message
	"retry"	    : "",														// If func have retry structure, Retry count print
	"file"	    : "setup.go",												// Doing func used file
	"func"	    : "update(namespase string, value string)",					// Doing func fullname
	"args"	    : ["kube-system","helm/elastic-search/values.yaml"],		// Doing args
	"document"  : "https://qiita.com/secondly/items/9c0ecd9c3ce583f6e0d3",	// Helpper link 
}
```
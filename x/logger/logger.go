package logger

import (
	"fmt"
	"log"
	"os"
)

// Mode = log print level(default=production mode, Mode debug=> debug mode
var Mode = os.Getenv("MODE")

/*
	log level
		EMERGE: このfunctionが最低限担保しなければならないことが守れなかった時 (cloudwatch alertが反応)
		WARN:	それ以外でのエラー
		INFO:	実行情報
		DEBUG:	試験的に表示しておきたい時（Lambda側の環境変数でmode=debugを指定する必要あり）
*/
type Option struct {
	id     string
	level  string
	format string
	v      interface{}
}

func doPrinter(id, level, s string) int {
	str := id + "  <" + level + "> " + s

	switch level {
	case "EMEGE", "WARN":
		log.Fatalln(str)
		return 1
	case "DEBUG":
		if Mode != "debug" && level == "DEBUG" {
			return 1
		}
	}

	log.Println(str)
	return 0
}

// Printer = 構造体で指定した形でログを出力する
func Printer(op *Option) int {
	if op.format != "" {
		s := fmt.Sprintf(op.format, op.v)
		return doPrinter(op.id, op.level, s)
	}

	s := fmt.Sprint(op.v)
	return doPrinter(op.id, op.level, s)
}

// Printf = フォーマットありで順番固定で引数を渡す
func Printf(id, level string, format string, v ...interface{}) {
	Printer(&Option{
		id:     id,
		level:  level,
		format: format,
		v:      v,
	})
}

// Print = フォーマットなしで順番固定で引数を渡す
func Print(id, level string, v ...interface{}) {
	Printer(&Option{
		id:    id,
		level: level,
		v:     v,
	})
}

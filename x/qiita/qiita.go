package qiita

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

var api = "https://qiita.com/api/v2/items"

type Qiita interface {
	Resp(c *gin.Context, msg ...interface{})
}

type RespFunc func(c *gin.Context, msg ...interface{})

func (res RespFunc) Resp(c *gin.Context, msg ...interface{}) {
	res(c, msg)
}

func Success() Qiita {
	return RespFunc(func(c *gin.Context, msg ...interface{}) {
		c.JSON(200, gin.H{
			"todo": msg,
		})
	})
}

func Error() Qiita {
	return RespFunc(func(c *gin.Context, msg ...interface{}) {
		c.JSON(500, gin.H{
			"todo": msg,
		})
	})
}

func GetItem(api string, d *interface{}) error {
	u, err := url.Parse(api)
	if err != nil {
		return err
	}

	fmt.Println(u.String())

	resp, err := http.Get(u.String())
	defer resp.Body.Close()

	fmt.Println(resp)

	if err = json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return err
	}
	fmt.Println(d)

	return nil
}

// Trend = qiita trend get
func Trend(c *gin.Context) {
	fmt.Println("Trend")

	f := Error()
	var msg string
	var data interface{}

	if err := GetItem(api, &data); err != nil {
		msg = fmt.Sprint(err)
		f.Resp(c, msg)
		return
	}

	// var jsonBody map[string]interface{}
	// if err = json.Unmarshal(body[:len(body)], &jsonBody); err != nil && err != io.EOF {
	// 	msg = fmt.Sprint(err)
	// 	f.Resp(c, msg)
	// }

	//	msg = fmt.Sprint("aaa")

	f = Success()
	f.Resp(c, data)

	// return nil
}

// TagTrend = qiita trend tag get
// func TagTrendHandler() {
// 	fmt.Println("TagTrend")
// 	TagTrend
// }

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

type Body struct {
	result interface{} `json:"result"`
}

type Item struct {
	title string `json:"title"`
}

func GetItem(api string) (*[]Item, error) {
	u, err := url.Parse(api)
	if err != nil {
		return nil, err
	}

	fmt.Println(u.String())

	resp, err := http.Get(u.String())
	defer resp.Body.Close()

	var row []map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&row); err != nil {
		return nil, err
	}

	var data []Item
	for _, r := range row {
		for k, v := range r {
			if str, ok := v.(string); ok && k == "title" {
				data = append(data, Item{str})
			}
		}
	}

	return &data, nil
}

// Trend = qiita trend get
func Trend(c *gin.Context) {
	fmt.Println("Trend")

	f := Error()
	var msg string
	data, err := GetItem(api)

	if err != nil {
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

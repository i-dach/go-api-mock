package router

import (
	q "go-api-mock/x/qiita"

	"github.com/gin-gonic/gin"
)

// GinRouter is router that API method router.
func Router(r *gin.Engine) error {

	// helth check
	r.GET("/helth", func(c *gin.Context) {
		// apm.TraceSeg(c, "/helth")

		c.JSON(200, gin.H{
			"message": "helth check ok",
		})
	})

	// Simple group
	qiita := r.Group("/qiita")
	{
		qiita.GET("/trend", q.Trend)
		// q.GET("/:tag", q.TagTrend)
	}

	return nil
}

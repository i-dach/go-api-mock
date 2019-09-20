package router

import (
	"go-api-mock/x/qiita"

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
	q := r.Group("/qiita/trend")
	{
		q.GET("/", qiita.Trend)
		q.GET("/:tag", qiita.TagTrend)
	}

	return nil
}

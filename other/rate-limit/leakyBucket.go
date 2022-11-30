package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

// 定义全局限流器对象
var rateLimit ratelimit.Limiter

// 在 gin.HandlerFunc 加入限流逻辑
func leakyBucket() gin.HandlerFunc {
	prev := time.Now()
	return func(c *gin.Context) {
		now := rateLimit.Take()
		fmt.Println(now.Sub(prev)) // 为了打印时间间隔
		prev = now                 // 记录上一次的时间，没有这个打印的会有问题
	}
}

func main() {
	rateLimit = ratelimit.New(10)
	r := gin.Default()
	r.GET("/ping", leakyBucket(), func(c *gin.Context) {
		c.JSON(200, true)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

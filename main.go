package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {

	// docker run -d --name myredis -p 6379:6379 redis --requirepass "mypassword"

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "mypassword", // no password set
		DB:       0,            // use default DB
	})
	pong, err := rdb.Ping(context.Background()).Result()
	if err == nil {
		log.Println("redis 回應成功，", pong)
	} else {
		log.Fatal("redis 無法連線，錯誤為", err)
	}

	server := gin.Default()
	server.GET("/", rateLimit)
	server.Run(":8000")
}

func rateLimit(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "rate limit",
		"ip":      c.ClientIP(),
		"times":   10,
	})
}

package main

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Response struct {
	Ip      string
	Request string
}

func main() {
	server := gin.Default()
	server.GET("/", rateLimit)
	server.Run(":8080")
}

func rateLimit(c *gin.Context) {
	var limit int64 = 60
	r := initRedis()
	response := &Response{Ip: c.ClientIP()}
	now := time.Now()
	begin := now.Add(time.Minute * (-1))

	removeExpired(r, response.Ip, begin)
	counts := getRecordCounts(r, response.Ip)
	createRecord(r, counts, limit, response.Ip, now, begin)
	SetResponse(counts, limit, response)
	c.JSON(200, response)
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "mypassword",
		DB:       0,
	})
}

func removeExpired(r *redis.Client, ip string, begin time.Time) {
	scoreBegin := strconv.FormatInt(begin.Unix(), 10)
	_, err := r.ZRemRangeByScore(context.Background(), ip, "0", scoreBegin).Result()

	if err != nil {
		panic(err)
	}
}

func getRecordCounts(r *redis.Client, ip string) int64 {
	counts, err := r.ZCard(context.Background(), ip).Result()
	if err != nil {
		panic(err)
	}

	return counts
}

func createRecord(r *redis.Client, counts int64, limit int64, ip string, now time.Time, begin time.Time) {
	if counts >= limit {
		return
	}

	err := r.ZAdd(context.Background(), ip, &redis.Z{Score: float64(now.Unix()), Member: begin}).Err()
	if err != nil {
		panic(err)
	}
}

func SetResponse(counts int64, limit int64, response *Response) {
	if counts < limit {
		response.Request = strconv.FormatInt(counts+1, 10)
	} else {
		response.Request = "Error"
	}
}

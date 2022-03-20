package main

import (
	"context"
	"io/ioutil"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var size = []int{
	10000,
	50000,
	100000,
	200000,
	300000,
	400000,
	500000,
}

var beforePrefix = "=============Before=============\n"
var afterPrefix = "==============After=============\n"

func main() {
	ctx := context.Background()
	wg := &sync.WaitGroup{}
	rdb, err := InitRedisClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, s := range size {

		rdb.FlushDB(ctx)

		before, err := rdb.Info(ctx, "memory").Result()
		if err != nil {
			log.Fatal(err.Error())
		}

		for i := 0; i < s; i++ {
			wg.Add(1)
			go func(data int) {
				rdb.Set(ctx, uuid.NewString(), 0, 0)
				wg.Done()
			}(i)
		}
		wg.Wait()

		after, err := rdb.Info(ctx, "memory").Result()
		if err != nil {
			log.Fatal(err.Error())
		}

		ioutil.WriteFile(
			"./q2/result_"+strconv.Itoa(s),
			[]byte(beforePrefix+"\n"+before+"\n\n"+afterPrefix+"\n"+after),
			0644,
		)
		rdb.FlushDB(ctx)
		time.Sleep(5 * time.Second)
	}

	defer rdb.Close()
}

func InitRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	return client, client.Ping(context.TODO()).Err()
}

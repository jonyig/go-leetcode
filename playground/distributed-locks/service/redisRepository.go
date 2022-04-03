package service

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const lockKey = "key1"
func Lock(redis *redis.Client) {
	var ctx = context.Background()

	for {
		nx := redis.SetNX(ctx, lockKey, GetCurrentGoroutineId(), 5*time.Second)
		result, err := nx.Result()

		if err == nil && result {
			log.Print("lock success ", GetCurrentGoroutineId())
			return
		}
	}
}
func Unlock(redis *redis.Client) {
	var ctx = context.Background()
	nx := redis.Get(ctx, lockKey)
	value, err := nx.Result()
	if err != nil || value == "" {
		return
	}

	if value != strconv.Itoa(GetCurrentGoroutineId()) {
		return
	}
	time.Sleep(2 * time.Second)
	if value, err := redis.Get(ctx, lockKey).Result(); err == nil {
		log.Printf("%d ==> %s", GetCurrentGoroutineId(), value)
	}
	// 此版會有解到別人 key 的問題
	result := redis.Del(ctx, lockKey)
	unlockSuccess, err := result.Result()
	if err == nil && unlockSuccess > 0 {
		log.Println("unlock success! ", GetCurrentGoroutineId())
	} else {
		log.Println("unlock failed!", err)
	}
}

func GetCurrentGoroutineId() int {

	buf := make([]byte, 128)

	buf = buf[:runtime.Stack(buf, false)]

	stackInfo := string(buf)

	goIdStr := strings.TrimSpace(strings.Split(strings.Split(stackInfo, "[running]")[0], "goroutine")[1])

	goId, err := strconv.Atoi(goIdStr)

	if err != nil {

		fmt.Println("err=", err)

		return 0

	}

	return goId

}

func UnlockUseLua(client *redis.Client) {
	time.Sleep(2 * time.Second)

	script := redis.NewScript(`
    if redis.call('get', KEYS[1]) == ARGV[1]
    	then
      		return redis.call('del', KEYS[1]) 
		else
        	return 0
     end
  	`)
	var ctx = context.Background()

	resp := script.Run(ctx, client, []string{lockKey}, GetCurrentGoroutineId())
	result, err := resp.Result()
	log.Print(result)
	log.Print(err)
	if err != nil || result == 0 {
		fmt.Println("unlock failed:", err)
	}
}

package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"testing"
)

func TestHash01(t *testing.T) {
	// 创建 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // Redis 密码
		DB:       0,                // 默认数据库
	})

	// 确保连接有效
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("无法连接到 Redis: %v", err)
	}
	fmt.Printf("连接成功: %s\n", pong)

	// 定义有序集合的键
	hashKey := "hashKey"

	// 清空有序集合（避免重复数据影响测试）
	rdb.Del(ctx, hashKey)

	for i := 0; i < 5; i++ {
		rdb.HMSet(ctx, hashKey, i)
	}

	// 关闭 Redis 连接
	err = rdb.Close()
	if err != nil {
		log.Fatalf("关闭 Redis 连接失败: %v", err)
	}
}

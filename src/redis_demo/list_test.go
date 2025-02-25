package redis_demo

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
	"testing"
)

func TestList01(t *testing.T) {
	// 创建 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // Redis 密码
		DB:       0,                // 默认数据库
	})

	// 确保连接有效
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal().Msgf("无法连接到 Redis: %v", err)
		log.Fatalf()
	}
	fmt.Printf("连接成功: %s\n", pong)

	// 定义有序集合的键
	listKey := "listKey"

	// 清空有序集合（避免重复数据影响测试）
	rdb.Del(ctx, listKey)

	for i := 0; i < 5; i++ {
		rdb.LPush(ctx, listKey, fmt.Sprintf("%d", i+1))
	}

	// 关闭 Redis 连接
	err = rdb.Close()
	if err != nil {
		log.Fatalf("关闭 Redis 连接失败: %v", err)
	}
}

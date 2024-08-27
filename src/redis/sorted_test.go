package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"testing"
)

// redis key:value结构   value:String、hash、list、set、sort set

// 创建全局的上下文
var ctx = context.Background()

func TestSorted01(t *testing.T) {
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
	zsetKey := "myzset"

	// 清空有序集合（避免重复数据影响测试）
	rdb.Del(ctx, zsetKey)

	// 向有序集合添加元素
	members := []*redis.Z{
		{Score: 10.5, Member: "member1"},
		{Score: 5.3, Member: "member2"},
		{Score: 20.1, Member: "member3"},
		{Score: 15.8, Member: "member4"},
	}

	// 批量添加元素
	for _, member := range members {
		err := rdb.ZAdd(ctx, zsetKey, member).Err()
		if err != nil {
			log.Fatalf("添加有序集合元素失败: %v", err)
		}
	}

	// 获取按分数排序的所有元素
	zsetElements, err := rdb.ZRangeWithScores(ctx, zsetKey, 0, -1).Result()
	if err != nil {
		log.Fatalf("获取有序集合元素失败: %v", err)
	}

	// 打印排序后的元素
	fmt.Println("有序集合按分数排序后的元素：")
	for _, z := range zsetElements {
		fmt.Printf("Member: %s, Score: %.2f\n", z.Member, z.Score)
	}

	// 获取分数从大到小排序的元素
	zsetElementsRev, err := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, -1).Result()
	if err != nil {
		log.Fatalf("获取有序集合元素失败: %v", err)
	}

	// 打印反向排序后的元素
	fmt.Println("有序集合按分数从大到小排序后的元素：")
	for _, z := range zsetElementsRev {
		fmt.Printf("Member: %s, Score: %.2f\n", z.Member, z.Score)
	}

	// 关闭 Redis 连接
	err = rdb.Close()
	if err != nil {
		log.Fatalf("关闭 Redis 连接失败: %v", err)
	}
}

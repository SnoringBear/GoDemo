package json_demo

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

// Item 结构体定义
type Item struct {
	ID          int                    `json:"id"`
	OtherFields map[string]interface{} `json:"-"`
}

func TestJson01(t *testing.T) {
	jsonArrayString := `
	[
		{
			"id": 1,
			"name": "Product A",
			"price": 99.99,
			"tags": ["electronics", "audio"]
		},
		{
			"id": 2,
			"service_name": "Subscription B",
			"monthly_fee": 15.50
		},
		{
			"id": 3,
			"title": "Article C",
			"author": "John Doe",
			"published": true
		}
	]`

	var stringSlice []map[string]interface{}
	err := json.Unmarshal([]byte(jsonArrayString), &stringSlice)
	if err != nil {
		log.Fatalf("反序列化失败: %v", err)
	}

	fmt.Println("解析后的 Go 字符串切片:")
	for _, item := range stringSlice {
		fmt.Println(item)
		i := item["id"]
		marshal, err := json.Marshal(item)
		if err != nil {
			continue
		}
		fmt.Printf("i = %v,marshal = %v\n", i, string(marshal))
	}
}

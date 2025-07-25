package agollo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

// ApolloConfig 存储 Apollo Open API 的连接信息
type ApolloConfig struct {
	PortalURL      string
	Token          string
	AppID          string
	Env            string
	ClusterName    string
	Namespace      string
	Operator       string
	RequestTimeout time.Duration
}

func TestAdd01(t *testing.T) {
	apolloClient := NewApolloConfig()

	// 1. 准备需要批量添加的配置
	configsToAdd := map[string]string{
		"go.batch.key1":            "go_value_1",
		"go.batch.key2":            "go_value_2",
		"go.batch.feature.enabled": "true",
	}

	fmt.Println("开始批量添加配置...")

	var hasError bool
	// 2. 循环调用新增接口
	for key, value := range configsToAdd {
		if err := apolloClient.CreateConfigItem(key, value); err != nil {
			fmt.Printf("添加配置出错: %v\n", err)
			hasError = true
		}
	}

	if hasError {
		fmt.Println("\n批量添加完成，但存在部分错误")
	} else {
		fmt.Println("\n所有配置添加成功")
	}
}

func NewApolloConfig() *ApolloConfig {
	return &ApolloConfig{
		PortalURL:      "http://localhost:8070",
		Token:          "948c81e4cfd942a0012cb80a9a817ea996fb73421faf0de129ee996ae7b2d5d9",
		AppID:          "SampleApp",
		Env:            "LOCAL",
		ClusterName:    "default",
		Namespace:      "application",
		Operator:       "apollo",
		RequestTimeout: 5 * time.Second,
	}
}

// CreateConfigItem 调用 Apollo API 创建单个配置项
func (c *ApolloConfig) CreateConfigItem(key, value string) error {
	apiURL := fmt.Sprintf("%s/openapi/v1/envs/%s/apps/%s/clusters/%s/namespaces/%s/items",
		c.PortalURL, c.Env, c.AppID, c.ClusterName, c.Namespace)

	payload := map[string]string{
		"key":                 key,
		"value":               value,
		"dataChangeCreatedBy": c.Operator,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("请求体序列化失败: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("创建HTTP请求失败: %w", err)
	}

	req.Header.Set("Authorization", c.Token)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client := &http.Client{Timeout: c.RequestTimeout}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("请求发送失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Printf("配置添加成功: %s = %s\n", key, value)
		return nil
	}

	respBody, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("添加配置项'%s'失败. 状态码: %s, 响应体: %s", key, resp.Status, string(respBody))
}

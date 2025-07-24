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

	fmt.Println("Starting to add configs in batch...")

	var hasError bool
	// 2. 循环调用新增接口
	for key, value := range configsToAdd {
		if err := apolloClient.CreateConfigItem(key, value); err != nil {
			fmt.Printf("Error: %v\n", err)
			hasError = true
		}
	}

	if hasError {
		fmt.Println("\nFinished batch adding process with one or more errors.")
		// 根据业务需求决定是否在部分失败时继续发布
		// return
	} else {
		fmt.Println("\nAll configs added successfully.")
	}

}

func NewApolloConfig() *ApolloConfig {
	return &ApolloConfig{
		PortalURL:      "http://localhost:8070",                                            // 替换为您的 Apollo Portal 地址
		Token:          "948c81e4cfd942a0012cb80a9a817ea996fb73421faf0de129ee996ae7b2d5d9", // 替换为您的 Token
		AppID:          "SampleApp",                                                        // 替换您的 App ID
		Env:            "LOCAL",
		ClusterName:    "default",
		Namespace:      "application",
		Operator:       "apollo", // 操作人标识
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
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create http request: %w", err)
	}

	req.Header.Set("Authorization", c.Token)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client := &http.Client{Timeout: c.RequestTimeout}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Printf("Successfully added config: %s = %s\n", key, value)
		return nil
	}

	// 读取响应体以获取更多错误信息
	respBody, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("failed to add config item '%s'. Status: %s, Body: %s", key, resp.Status, string(respBody))
}

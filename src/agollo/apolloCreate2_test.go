package agollo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type Namespace struct {
	Name                string `json:"name"`
	AppID               string `json:"appId"`
	Format              string `json:"format"`
	IsPublic            bool   `json:"isPublic"`
	Comment             string `json:"comment,omitempty"`
	DataChangeCreatedBy string `json:"dataChangeCreatedBy"`
}

func TestCreate02(t *testing.T) {
	portalAddress := "http://localhost:8070"
	token := "948c81e4cfd942a0012cb80a9a817ea996fb73421faf0de129ee996ae7b2d5d9"
	creator := "apollo"

	privateNamespace := Namespace{
		Name:                "my-private-namespace",
		AppID:               "SampleApp",
		Format:              "properties",
		IsPublic:            false, // 是否是私有命名空间
		Comment:             "This is a private namespace for my application.",
		DataChangeCreatedBy: creator,
	}

	err := CreateApolloNamespace(portalAddress, token, privateNamespace)
	if err != nil {
		fmt.Println("Error creating private namespace:", err)
	}
}

// CreateApolloNamespace 创建Apollo命名空间
func CreateApolloNamespace(portalAddress, token string, namespace Namespace) error {
	payload, err := json.Marshal(namespace)
	if err != nil {
		return fmt.Errorf("failed to marshal namespace payload: %w", err)
	}

	url := fmt.Sprintf("%s/openapi/v1/apps/%s/appnamespaces", portalAddress, namespace.AppID)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create new HTTP request: %w", err)
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client := &http.Client{}
	// http请求
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create namespace, status code: %d, response: %s", resp.StatusCode, string(body))
	}

	return nil
}

// createOrUpdateConfigItem 调用 Apollo API 在指定 Namespace 中创建或更新一个配置项
func createOrUpdateConfigItem(namespaceName, content string) error {
	// API 端点: /openapi/v1/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/items
	url := fmt.Sprintf("%s/openapi/v1/envs/%s/apps/%s/clusters/%s/namespaces/%s/items",
		apolloPortalAddress, env, appId, clusterName, namespaceName)

	reqBody := CreateOrUpdateItemRequest{
		Key:                 "config", // Key 固定为 "config"
		Value:               content,  // Value 是文件内容
		Comment:             fmt.Sprintf("Updated by Go script on %s", time.Now().Format("2006-01-02")),
		DataChangeCreatedBy: operator,
	}

	jsonBytes, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("序列化配置项请求失败: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return fmt.Errorf("创建 HTTP 请求失败: %w", err)
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("发送 HTTP 请求失败: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API 返回错误状态码 %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

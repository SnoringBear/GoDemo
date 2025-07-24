package agollo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

// Apollo Portal 的地址
const apolloPortalAddress = "http://localhost:8070"

// 要操作的 Apollo App ID
const appId = "SampleApp"

// 要操作的环境 (DEV, FAT, UAT, PRO...)
const env = "LOCAL"

// 要操作的集群名称 (通常是 default)
const clusterName = "default"

// Apollo 开放平台授权的 Token
const token = "948c81e4cfd942a0012cb80a9a817ea996fb73421faf0de129ee996ae7b2d5d9"

// 操作人（显示在 Apollo 操作历史上）
const operator = "apollo"

// 存放配置文件的本地文件夹路径
const configFolderPath = "../Tables"

// --- 配置结束 ---

// 创建或更新配置项 API 的请求体结构
type CreateOrUpdateItemRequest struct {
	Key                 string `json:"key"`
	Value               string `json:"value"`
	Comment             string `json:"comment"`
	DataChangeCreatedBy string `json:"dataChangeCreatedBy"`
}

type Item struct {
	ID          int                    `json:"id"`
	OtherFields map[string]interface{} `json:"-"`
}

func TestCreate01(t *testing.T) {
	portalAddress := "http://localhost:8070"

	// 1. 读取文件夹中的所有文件
	files, err := os.ReadDir(configFolderPath)
	if err != nil {
		log.Fatalf("无法读取文件夹 '%s': %v", configFolderPath, err)
	}

	log.Printf("开始处理文件夹 '%s' 下的配置文件...", configFolderPath)

	for _, file := range files {
		// 忽略子目录
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		// 将文件名作为 namespace 名称
		// 通常 namespace 不包含文件扩展名，这里我们去掉它
		namespaceName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		namespaceName = strings.TrimPrefix(namespaceName, "Table_")

		log.Printf("--- 正在处理文件: %s, 准备创建 Namespace: %s ---", fileName, namespaceName)

		privateNamespace := Namespace{
			Name:                namespaceName,
			AppID:               "SampleApp",
			Format:              "properties",
			IsPublic:            false,
			Comment:             "This is a private namespace for my application.",
			DataChangeCreatedBy: operator,
		}

		// 2. 为每个文件创建一个私有 Namespace
		err := CreateApolloNamespace(portalAddress, token, privateNamespace)
		if err != nil {
			// 如果 Namespace 已存在，API会返回4xx错误，这里我们选择继续而不是中止
			log.Printf("创建 Namespace '%s' 失败或已存在: %v", namespaceName, err)
		} else {
			log.Printf("成功创建私有 Namespace: %s", namespaceName)
		}

		// 3. 读取文件内容
		filePath := filepath.Join(configFolderPath, fileName)
		contentBytes, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("读取文件 '%s' 内容失败: %v", filePath, err)
			continue // 继续处理下一个文件
		}
		fileContent := string(contentBytes)

		var stringSlice []map[string]interface{}
		err = json.Unmarshal([]byte(fileContent), &stringSlice)
		if err != nil {
			log.Printf("反序列化失败: %v", err)
			continue
		}

		for _, item := range stringSlice {
			i := item["ID"]
			marshal, err := json.Marshal(item)
			if err != nil {
				continue
			}
			key, ok := i.(float64)
			if !ok {
				continue
			}
			// 4. 将文件内容作为 key 为 "config" 的配置项添加
			err = createOrUpdateConfigItem2(namespaceName, string(marshal), strconv.FormatFloat(key, 'f', -1, 64))
			if err != nil {
				log.Printf("为 Namespace '%s' 添加配置项失败: %v", namespaceName, err)
			} else {
				log.Printf("成功为 Namespace '%s' 添加 key='config' 的配置项", namespaceName)
			}
		}

	}

	log.Println("--- 所有文件处理完毕 ---")
}

// UnmarshalJSON 的实现
func (i *Item) UnmarshalJSON(data []byte) error {
	var tempMap map[string]interface{}
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if idVal, ok := tempMap["id"]; ok {
		if idFloat, ok := idVal.(float64); ok {
			i.ID = int(idFloat)
		}
	}

	delete(tempMap, "id")
	i.OtherFields = tempMap
	return nil
}

func UnmarshalJsonArray(data []byte) {

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

	// 使用 PUT 方法可以实现"创建或更新"的效果
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

// createOrUpdateConfigItem 调用 Apollo API 在指定 Namespace 中创建或更新一个配置项
func createOrUpdateConfigItem2(namespaceName, content string, key string) error {
	// API 端点: /openapi/v1/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/items
	url := fmt.Sprintf("%s/openapi/v1/envs/%s/apps/%s/clusters/%s/namespaces/%s/items",
		apolloPortalAddress, env, appId, clusterName, namespaceName)

	reqBody := CreateOrUpdateItemRequest{
		Key:                 key,     // Key 固定为 "config"
		Value:               content, // Value 是文件内容
		Comment:             fmt.Sprintf("Updated by Go script on %s", time.Now().Format("2006-01-02")),
		DataChangeCreatedBy: operator,
	}

	jsonBytes, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("序列化配置项请求失败: %w", err)
	}

	// 使用 PUT 方法可以实现"创建或更新"的效果
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

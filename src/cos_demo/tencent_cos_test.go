package cos_demo

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// Config 存储所有必要的配置信息
type Config struct {
	SecretID     string        // 腾讯云 API SecretId
	SecretKey    string        // 腾讯云 API SecretKey
	BucketURL    string        // COS 存储桶 URL, 格式: https://<BucketName-APPID>.cos.<Region>.myqcloud.com
	LocalPath    string        // 本地下载目录
	PollInterval time.Duration // 轮询间隔
}

// 全局配置变量
var appConfig = Config{
	// ============================ 请在这里修改您的配置 ============================
	SecretID:     "AKIDBRWP4Zr0BPZZVK3pTzY3oQp0zkHbmqiN",                            // 替换为您的 SecretId
	SecretKey:    "Vegi4IRVP6au2zjFYdZKsIWeaUM2Vkmb",                                // 替换为您的 SecretKey
	BucketURL:    "https://wechat-mini-game-1259432156.cos.ap-chengdu.myqcloud.com", // 替换为您的存储桶 URL
	LocalPath:    "./cos_downloads",                                                 // 本地存储目录，可按需修改
	PollInterval: 10 * time.Second,                                                  // 每 10 秒检查一次
	// ===========================================================================
}

// processedFiles 用于跟踪已处理的文件及其 ETag
// key: COS上的文件路径 (object key)
// value: 文件的 ETag
var processedFiles = make(map[string]string)

func TestCOS01(t *testing.T) {
	// 1. 创建本地下载目录（如果不存在）
	if _, err := os.Stat(appConfig.LocalPath); os.IsNotExist(err) {
		if err := os.MkdirAll(appConfig.LocalPath, os.ModePerm); err != nil {
			log.Fatalf("创建本地目录失败: %v", err)
		}
	}

	// 2. 初始化 COS 客户端
	u, err := url.Parse(appConfig.BucketURL)
	if err != nil {
		log.Fatalf("解析 Bucket URL 失败: %v", err)
	}
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  appConfig.SecretID,
			SecretKey: appConfig.SecretKey,
		},
	})

	// 3. 首次启动时先执行一次检查
	log.Println("执行首次文件状态检查...")
	checkAndDownloadFiles(client)

	// 4. 启动定时器，按指定间隔轮询
	ticker := time.NewTicker(appConfig.PollInterval)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("----------------------------------")
		log.Println("开始新一轮的变更检查...")
		checkAndDownloadFiles(client)
	}
}

// checkAndDownloadFiles 检查存储桶中的文件并下载变更
func checkAndDownloadFiles(client *cos.Client) {
	// 获取存储桶中的对象列表
	opt := &cos.BucketGetOptions{
		Prefix:  "", // 可指定前缀来监听特定目录
		MaxKeys: 1000,
	}
	resp, _, err := client.Bucket.Get(context.Background(), opt)
	if err != nil {
		log.Printf("错误: 获取对象列表失败: %v", err)
		return
	}

	foundChanges := false
	for _, object := range resp.Contents {
		// 如果是目录，则跳过 (COS中的目录本质上是以'/'结尾的空对象)
		if strings.HasSuffix(object.Key, "/") {
			continue
		}

		// ETag 通常被引号包围，例如 "d41d8cd98f00b204e9800998ecf8427e"，需要去除
		currentETag := strings.Trim(object.ETag, `"`)

		// 检查文件是否是新的或已更新
		if previousETag, ok := processedFiles[object.Key]; !ok || previousETag != currentETag {
			foundChanges = true
			log.Printf("检测到变更: 文件名='%s', 新ETag='%s'", object.Key, currentETag)

			// 下载文件
			err := downloadFile(client, object.Key, appConfig.LocalPath)
			if err != nil {
				log.Printf("错误: 下载文件 '%s' 失败: %v", object.Key, err)
			} else {
				// 下载成功后，更新状态
				processedFiles[object.Key] = currentETag
				log.Printf("成功: 文件 '%s' 已下载并更新状态。", object.Key)
			}
		}
	}

	if !foundChanges {
		log.Println("未检测到任何文件变更。")
	}
}

// downloadFile 从 COS 下载单个文件到本地
func downloadFile(client *cos.Client, key, localDir string) error {
	// 构造本地文件的完整路径
	localFilePath := filepath.Join(localDir, key)

	// 确保本地子目录存在
	if err := os.MkdirAll(filepath.Dir(localFilePath), os.ModePerm); err != nil {
		return fmt.Errorf("创建本地子目录失败: %w", err)
	}

	// 从COS获取对象
	resp, err := client.Object.Get(context.Background(), key, nil)
	if err != nil {
		return fmt.Errorf("请求COS对象失败: %w", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("获取文件时返回非200状态码: %d", resp.StatusCode)
	}

	// 创建本地文件用于写入。os.Create 会自动创建新文件或清空已存在的文件，实现覆盖效果。
	localFile, err := os.Create(localFilePath)
	if err != nil {
		return fmt.Errorf("创建本地文件失败: %w", err)
	}
	defer localFile.Close()

	// 将COS响应体的内容拷贝到本地文件
	_, err = io.Copy(localFile, resp.Body)
	if err != nil {
		return fmt.Errorf("写入本地文件失败: %w", err)
	}

	return nil
}

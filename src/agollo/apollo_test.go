package agollo

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	"testing"
)

func Test01(t *testing.T) {
	c := &config.AppConfig{
		AppID:          "SampleApp",
		Cluster:        "default",
		IP:             "http://localhost:8080",
		NamespaceName:  "application",
		IsBackupConfig: true,
	}

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})

	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("Apollo 客户端初始化成功")
	// 后续操作...
	// 获取默认 namespace (application) 的配置
	value := client.GetStringValue("timeout", "default_value")
	fmt.Printf("获取到配置 timeout 的值为: %s\n", value)
}

func Test02(t *testing.T) {
	c := &config.AppConfig{
		AppID:          "SampleApp",
		Cluster:        "default",
		IP:             "http://localhost:8080",
		NamespaceName:  "application",
		IsBackupConfig: true,
		Secret:         "", // 如果您的 Namespace 设置了密钥
	}

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})

	if err != nil {
		fmt.Println("err:", err)
		return
	}
	// Create an instance of your custom listener
	listener := &CustomChangeListener{}

	// Add the change listener
	client.AddChangeListener(listener)
	// 为了演示，让主程序保持运行
	for {

	}
}

type CustomChangeListener struct{}

func (c *CustomChangeListener) OnNewestChange(event *storage.FullChangeEvent) {
	fmt.Println("OnNewestChange:", event)
	for key, change := range event.Changes {
		fmt.Printf(" NewValue: %v, ChangeType: %s\n",
			key, change)
	}
}

// OnChange will be called when a configuration change occurs.
func (c *CustomChangeListener) OnChange(changeEvent *storage.ChangeEvent) {
	fmt.Printf("Configuration changed for namespace: %s\n", changeEvent.Namespace)
	for key, change := range changeEvent.Changes {
		fmt.Printf("  Key: %v, OldValue: %v, NewValue: %v, ChangeType: %v\n",
			key, change.OldValue, change.NewValue, change.ChangeType)
	}
}

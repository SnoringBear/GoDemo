package label_demo

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"testing"
)

type Person struct {
	Name    string `mapstructure:"name"`
	Age     int    `mapstructure:"age"`
	Address string `mapstructure:"address"`
}

// 在 Go 中，mapstructure 是一个非常有用的库，通常用于将结构体从映射（如 map）转换为 Go 结构体。这通常用于 JSON 解析或其他数据格式的映射，尤其是在处理配置文件或者 API 响应时。
//
// mapstructure 库支持在映射过程中使用标签来控制字段的解析，允许你指定如何从映射的键名与结构体字段匹配。
// github.com/spf13/viper 解析配置的第三方库,引用了github.com/mitchellh/mapstructure库
func TestMapStructure01(t *testing.T) {
	// 解析的字段必须一致
	data := map[string]interface{}{
		"name":     "John Doe",
		"age":      30,
		"address2": "123 Main St",
	}

	// 创建一个空的结构体实例
	var person Person

	// 使用mapstructure库解码map数据
	err := mapstructure.Decode(data, &person)
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}
	log.Info().Msgf("Decoded Person: %+v \n", person)
}

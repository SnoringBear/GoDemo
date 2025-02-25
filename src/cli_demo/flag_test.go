package cli_demo

import (
	"flag"
	"github.com/rs/zerolog/log"
	"testing"
)

var config = flag.String("config", "application.yml", "config file")
var game = flag.String("game", "game.json", "game config file")

func TestFlagTest01(t *testing.T) {
	// cli 格式为 -x1=a -x2=b
	// 获取的变量必须定义为全局变量,而不是方法里面的局部变量，否则会报错
	// 与cobra相比,区别:cobra支持子命令、或其他更复杂的命令结构
	flag.Parse()
	log.Info().Msgf("config value: %v", *config)
	log.Info().Msgf("game value: %v", *game)
}

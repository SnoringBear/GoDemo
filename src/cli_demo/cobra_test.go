package cli_demo

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
	"testing"
)

var rootCmd = &cobra.Command{
	Use:   "connector",                    //  指定了这个命令的名称是connector。这意味着，如果你的程序编译后命名为myapp，那么在命令行中运行myapp connector就会触发这个命令
	Short: "connector 管理连接，session以及路由请求", // 短描述
	Long:  `connector 管理连接，session以及路由请求`, // 长描述
	//  这是一个函数，当connector命令被执行时会被调用。这里的函数体为空，意味着执行connector命令时不会有任何操作。
	//  在实际应用中，你可以在这个函数体内添加执行特定任务的代码。
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msgf("cmd:%v, args:%v", cmd, args)
	},
	// 这是一个在Run函数执行之后立即执行的函数。它同样接收一个cobra.Command实例和一个字符串切片作为参数。
	// 这里的函数体也是空的，但在实际应用中，你可以在这个函数体内添加一些需要在主命令执行完成后立即进行的清理工作或后续操作
	PostRun: func(cmd *cobra.Command, args []string) {
		log.Info().Msgf("cmd2:%v, args2:%v", cmd, args)
	},
}
var (
	configFile    string
	gameConfigDir string
	serverId      string
)

func init() {
	rootCmd.Flags().StringVar(&configFile, "config", "application.yml", "app config yml file")
	rootCmd.Flags().StringVar(&gameConfigDir, "gameDir", "../config", "game config dir")
	rootCmd.Flags().StringVar(&serverId, "serverId", "connector001", "app server id， required")
	//_ = rootCmd.MarkFlagRequired("serverId")
}

func TestCobraTest01(t *testing.T) {
	// 使用方式  Use:命令+ config=x1 gameDir=x2 serverId=x3
	//1.加载配置
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	log.Info().Msgf("config:%v", configFile)
	log.Info().Msgf("gameDir:%v", gameConfigDir)
	log.Info().Msgf("serverId:%v", serverId)
}

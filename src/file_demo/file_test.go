package file_demo

import (
	"github.com/rs/zerolog/log"
	"os"
	"testing"
)

func TestFileExist01(t *testing.T) {
	dir, _ := os.Getwd()
	log.Info().Msgf("current dir:%s", dir) // reslut:D:\\GoDemo\\src\\file_demo
	fileInfo, err := os.Stat(dir + "/file.md")
	if err != nil {
		log.Error().Msgf("err:%v", err)
		return
	}
	log.Info().Msgf("fileInfo:%v", fileInfo)
}

func TestFile02(t *testing.T) {
	file, err := os.Open("file.md")
	if err != nil {
		log.Error().Msgf("err:%v", err)
	}
	defer file.Close()
	log.Info().Msgf("file:%v", file)
}

package time_demo

import (
	"github.com/rs/zerolog/log"
	"testing"
	"time"
)

func TestTime01(t *testing.T) {
	nano := time.Now().Unix()
	log.Info().Msgf("nano: %d", nano)
}

func TestTime02(t *testing.T) {
	now := time.Now() //获取当前时间
	log.Info().Msgf("current time_demo:%v", now)
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	log.Info().Msgf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
}

func TestTime03(t *testing.T) {
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	log.Info().Msgf("current timestamp1:%v", timestamp1)
	log.Info().Msgf("current timestamp2:%v", timestamp2)
}

func TestTime04(t *testing.T) {
	timeObj := time.Unix(time.Now().Unix(), 0) //将时间戳转为时间格式
	log.Info().Msgf("时间格式s= %s", timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	log.Info().Msgf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
}

func TestTime05(t *testing.T) {
	log.Info().Msgf("Duration = %v", time.Second)
}

func TestTime06(t *testing.T) {
	now := time.Now()
	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	// 时间操作 Add、Sub、Equal、Before、After
	log.Info().Msgf("当前时间加1小时后的时间:%v", later)
}

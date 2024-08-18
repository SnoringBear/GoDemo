package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTime01(t *testing.T) {
	nano := time.Now().Unix()
	fmt.Println(nano)
}

func TestTime02(t *testing.T) {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func TestTime03(t *testing.T) {
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

func TestTime04(t *testing.T) {
	timeObj := time.Unix(time.Now().Unix(), 0) //将时间戳转为时间格式
	fmt.Printf("时间格式s= %s \n", timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func TestTime05(t *testing.T) {
	fmt.Printf("Duration = %d\n", time.Second)
}

func TestTime06(t *testing.T) {
	now := time.Now()
	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	// 时间操作 Add、Sub、Equal、Before、After
	fmt.Println(later)
}

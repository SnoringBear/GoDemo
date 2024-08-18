package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeTwo01(t *testing.T) {
	// 在 Go 中，你可以使用 time.Date 函数来构造一个指定日期和时间的 time.Time 对象，
	// 然后通过调用 Unix() 或 UnixNano() 获取该时间的 Unix 时间戳

	// 指定日期和时间
	year := 2023
	month := time.August // 或者使用 time.Month 类型的其他月份
	day := 18
	hour := 12
	minute := 30
	second := 0

	// 创建指定时间的 time.Time 对象
	specifiedTime := time.Date(year, month, day, hour, minute, second, 0, time.UTC)

	// 获取秒级时间戳
	timestamp := specifiedTime.Unix()

	// 获取毫秒级时间戳
	timestampMillis := specifiedTime.UnixNano() / int64(time.Millisecond)

	fmt.Println("指定时间的秒级时间戳:", timestamp)
	fmt.Println("指定时间的毫秒级时间戳:", timestampMillis)
}

func TestTimeTwo02(t *testing.T) {
	// 获取当前时间
	now := time.Now()

	// 获取服务器当前的时区
	location := now.Location()

	// 获取服务器当前时区的名称和偏移量
	zone, offset := now.Zone()

	offsetHours := offset / 3600

	fmt.Println("服务器当前时区:", location)
	fmt.Println("服务器当前时区名称:", zone)
	fmt.Println("服务器时区偏移量（秒）:", offset)
	fmt.Println("服务器时区偏移量（小时）:", offsetHours)
}

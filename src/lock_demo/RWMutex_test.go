package lock_demo

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRW01(t *testing.T) {
	main1()
}

var mu1 sync.RWMutex
var count int

func main1() {
	go A1()
	time.Sleep(2 * time.Second)
	mu1.Lock()
	defer mu1.Unlock()
	count++
	fmt.Println(count)
}
func A1() {
	mu1.RLock()
	defer mu1.RUnlock()
	B1()
}
func B1() {
	time.Sleep(5 * time.Second)
	C1()
}
func C1() {
	mu1.RLock()
	defer mu1.RUnlock()
}

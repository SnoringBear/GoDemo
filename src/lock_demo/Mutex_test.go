package lock_demo

import (
	"fmt"
	"sync"
	"testing"
)

func TestErrorLock01(t *testing.T) {
	main()
	// TestLock01 方法运行时会报 panic，原因是它调用了 main()，而 main() 函数内部会调用 A()，
	// A() 会加锁 mu.Lock()，然后在 defer 里解锁。接着 A() 又调用 B()，B() 再调用 C()，
	// 而 C() 又尝试对同一个 mu 进行加锁 mu.Lock()，但此时锁还没有被释放（因为 A() 的 defer 还没执行），
	// 导致死锁，最终程序会 panic（fatal error: all goroutines are asleep - deadlock!）
	// 根本原因：Go 的 sync.Mutex 不是可重入锁，不能在同一个 goroutine 里多次加锁。
}

var mu sync.Mutex
var chain string

func main() {
	chain = "main"
	A()
	fmt.Println(chain)
}
func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> A"
	B()
}
func B() {
	chain = chain + " --> B"
	C()
}
func C() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> C"
}

func TestErrorLock02(t *testing.T) {

}

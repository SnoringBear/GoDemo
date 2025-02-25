package concurrent_demo

import (
	"fmt"
	"sync"
	"testing"
)

var mutex sync.Mutex
var wg sync.WaitGroup

func TestLock01(t *testing.T) {
	// golang 常见的锁  sync.Mutex（互斥锁）   sync.RWMutex（读写互斥锁）    sync.Once     sync.WaitGroup
	p := &Person{age: 1}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go p.changeAge()
	}
	wg.Wait()
	fmt.Printf("结果 = %d \n", p.age)
}

type Person struct {
	age int
}

func (p *Person) changeAge() {
	mutex.Lock()
	p.age++
	wg.Done()
	defer mutex.Unlock()
}

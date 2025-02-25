package lock_demo

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// TestWaitGroup WaitGroup：可以用来协调多个coroutine，等待多个coroutine执行完成
func TestWaitGroup(t *testing.T) {
	var sm map[string]string

	sm = make(map[string]string)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		sm["key1"] = "value1"
	}()

	go func() {
		defer wg.Done()
		sm["key2"] = "value2"
	}()

	wg.Wait()

	for s, s2 := range sm {
		fmt.Println(s, ":", s2)
	}
}

// TestSyncCond  sync.Cond:条件变量用来协调想要访问共享资源的那些 goroutine，当共享资源的状态发生变化的时候，它可以用来通知被互斥锁阻塞的 goroutine
func TestSyncCond(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})
	ready := false

	wg := sync.WaitGroup{}
	wg.Add(2)

	// 每次打印的顺序有可能不一样，下面两个coroutine执行的顺序不是固定，而是随机的

	go func() {
		defer wg.Done()
		cond.L.Lock()
		fmt.Println("Goroutine 2 lock")
		time.Sleep(1 * time.Minute)
		ready = true
		cond.L.Unlock()
		fmt.Println("Goroutine 2 unlock")
		cond.Signal() // Notify one waiting goroutine
		fmt.Println("Goroutine send signal")
	}()

	// Goroutine waiting for condition
	go func() {
		defer wg.Done()
		cond.L.Lock()
		fmt.Println("Goroutine 11 is running")
		for !ready {
			fmt.Println("Wait")
			cond.Wait()
		}
		fmt.Println("Goroutine 12 is running")
		cond.L.Unlock()
	}()

	wg.Wait()
}

// TestSyncOnce  sync.Once:
func TestSyncOnce(t *testing.T) {
	var once sync.Once
	wg := sync.WaitGroup{}
	wg.Add(3)

	initialize := func() {
		fmt.Println("Initialized")
	}

	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			once.Do(initialize)
		}()
	}

	wg.Wait()
}

// TestSyncRWMutex  sync.RWMutex:
func TestSyncRWMutex(t *testing.T) {
	var rwMu sync.RWMutex
	counter := 0

	wg := sync.WaitGroup{}
	wg.Add(3)

	// Reader 1
	go func() {
		defer wg.Done()
		rwMu.RLock()
		fmt.Println("Reader 1:", counter)
		rwMu.RUnlock()
	}()

	// Reader 2
	go func() {
		defer wg.Done()
		rwMu.RLock()
		fmt.Println("Reader 2:", counter)
		rwMu.RUnlock()
	}()

	// Writer
	go func() {
		defer wg.Done()
		rwMu.Lock()
		counter++
		rwMu.Unlock()
	}()

	wg.Wait()
}

// TestSyncMutex  sync.Mutex:
func TestSyncMutex(t *testing.T) {
	var mu sync.Mutex
	counter := 0

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		mu.Lock()
		counter++
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		counter++
		mu.Unlock()
	}()

	wg.Wait()
	fmt.Println("Counter:", counter)
}

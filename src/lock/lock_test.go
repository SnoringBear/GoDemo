package lock

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	var sm sync.Map

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		sm.Store("key1", "value1")
	}()

	go func() {
		defer wg.Done()
		sm.Store("key2", "value2")
	}()

	wg.Wait()

	sm.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}

func TestMap(t *testing.T) {
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

func TestSyncCond(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})
	ready := false

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		cond.L.Lock()
		fmt.Println("Goroutine 2 is preparing")
		ready = true
		cond.L.Unlock()
		cond.Signal() // Notify one waiting goroutine
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

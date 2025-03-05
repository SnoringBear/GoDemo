package file_demo

import (
	"fmt"
	"sync"
	"testing"
)

var wg2 sync.WaitGroup

func TestCh01(t *testing.T) {
	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "world"

	for i := 0; i < 2; i++ {
		wg2.Add(1)
		go chDemo01(ch)
	}
	wg2.Wait()
}

func chDemo01(ch chan string) {
	defer wg2.Done()
	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		default:
			return
		}
	}
}

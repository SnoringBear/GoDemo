package runtime_demo

import (
	"fmt"
	"runtime"
	"testing"
)

func TestOS1(t *testing.T) {
	fmt.Printf("runtime.GOOS:%#v \n", runtime.GOOS)
	fmt.Printf("runtime.GOARCH:%#v \n", runtime.GOARCH)
}

func TestOS2(t *testing.T) {
	runtime.GC()
}

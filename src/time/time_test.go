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

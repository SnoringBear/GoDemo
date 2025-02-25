package time_demo

import (
	"github.com/rs/zerolog/log"
	"testing"
	"time"
)

func TestTick01(t *testing.T) {
	// 定时器
	tick := time.Tick(time.Second)
	for i := range tick {
		log.Info().Msgf("i:%v", i)
	}
}

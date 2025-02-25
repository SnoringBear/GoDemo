package map_demo

import (
	"github.com/rs/zerolog/log"
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	var sm sync.Map

	sm.Store("key1", "value1")
	sm.Store("key2", "value2")

	log.Info().Msg("开始遍历map")
	sm.Range(func(key, value interface{}) bool {
		log.Info().Msgf("key:%s value:%s", key, value)
		return true
	})

}

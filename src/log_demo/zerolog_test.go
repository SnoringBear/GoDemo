package log_demo

import (
	"github.com/rs/zerolog/log"
	"testing"
)

func TestLog01(t *testing.T) {
	log.Info().Msg("Hello Zero log")
}

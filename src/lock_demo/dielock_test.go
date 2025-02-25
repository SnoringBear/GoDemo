package lock_demo

import (
	"github.com/rs/zerolog/log"
	"sync"
	"testing"
)

func TestDieLock01(t *testing.T) {
	var mu1, mu2 sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine 1
	go func() {
		defer wg.Done()
		mu1.Lock()
		log.Info().Msgf("Goroutine 1: locked mu1")

		// Simulate some work with sleep
		log.Info().Msgf("Goroutine 1: waiting to lock mu2...")

		mu2.Lock()
		log.Info().Msgf("Goroutine 1: locked mu2")

		// Unlock both mutexes
		mu2.Unlock()
		mu1.Unlock()
	}()

	// Goroutine 2
	go func() {
		defer wg.Done()
		mu2.Lock()
		log.Info().Msgf("Goroutine 2: locked mu2")

		// Simulate some work with sleep
		log.Info().Msgf("Goroutine 2: waiting to lock mu1...")
		mu1.Lock()
		log.Info().Msgf("Goroutine 2: locked mu1")

		// Unlock both mutexes
		mu1.Unlock()
		mu2.Unlock()
	}()

	// Wait for both goroutines to finish
	wg.Wait()
}

func TestDieLock02(t *testing.T) {
	var mu1, mu2 sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine 1
	go func() {
		defer wg.Done()
		mu1.Lock()
		log.Info().Msgf("Goroutine 1: locked mu1")

		// Simulate some work with sleep
		log.Info().Msgf("Goroutine 1: waiting to lock mu2...")
		mu2.Lock()
		log.Info().Msgf("Goroutine 1: locked mu2")

		// Unlock both mutexes
		mu1.Unlock()
		mu2.Unlock()
	}()

	// Goroutine 2
	go func() {
		defer wg.Done()
		mu2.Lock()
		log.Info().Msgf("Goroutine 2: locked mu2")

		// Simulate some work with sleep
		log.Info().Msgf("Goroutine 2: waiting to lock mu1...")
		mu1.Lock()
		log.Info().Msgf("Goroutine 2: locked mu1")

		// Unlock both mutexes
		mu2.Unlock()
		mu1.Unlock()
	}()

	// Wait for both goroutines to finish
	wg.Wait()
}

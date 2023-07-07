package semaphore_test

import (
	"math/rand"
	"runtime"
	"testing"
	"time"

	"github.com/g3offrey/semaphore"
)

func GetRandomDuration(max time.Duration) time.Duration {
	return time.Duration(rand.Int63n(int64(max)))
}

func TestSemaphore(t *testing.T) {
	capacity := 3
	sem := semaphore.Make(capacity)
	initialNumberOfGoRoutines := runtime.NumGoroutine()
	expectedNumberOfGoRoutines := initialNumberOfGoRoutines + capacity

	for i := 0; i < 50; i++ {
		sem.Acquire()

		go func() {
			defer sem.Release()

			time.Sleep(GetRandomDuration(100) * time.Millisecond)
			if runtime.NumGoroutine() > expectedNumberOfGoRoutines {
				t.Errorf("Too many goroutines running concurrently, expecting %d, got %d", expectedNumberOfGoRoutines, runtime.NumGoroutine())
			}
		}()
	}

	sem.Wait()

	if runtime.NumGoroutine() != initialNumberOfGoRoutines {
		t.Errorf("Expected no goroutine, got %d", runtime.NumGoroutine()-initialNumberOfGoRoutines)
	}
}

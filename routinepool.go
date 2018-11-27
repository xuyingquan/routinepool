package routinepool

import (
	"errors"
	"math"
)

const (
	// DefaultAntsPoolSize is the default capacity for a default goroutine pool.
	DefaultAntsPoolSize = math.MaxInt32

	// DefaultCleanIntervalTime is the interval time to clean up goroutines.
	DefaultCleanIntervalTime = 5
)

// Init a instance pool when importing ants.
var defaultRoutinePool, _ = NewPool(DefaultAntsPoolSize)

// Submit submits a task to pool.
func Submit(task f) error {
	return defaultRoutinePool.Submit(task)
}

// Running returns the number of the currently running goroutines.
func Running() int {
	return defaultRoutinePool.Running()
}

// Cap returns the capacity of this default pool.
func Cap() int {
	return defaultRoutinePool.Cap()
}

// Free returns the available goroutines to work.
func Free() int {
	return defaultRoutinePool.Free()
}

// Release Closes the default pool.
func Release() {
	defaultRoutinePool.Release()
}

// Errors for the Ants API.
var (
	ErrInvalidPoolSize   = errors.New("invalid size for pool")
	ErrInvalidPoolExpiry = errors.New("invalid expiry for pool")
	ErrPoolClosed        = errors.New("this pool has been closed")
)


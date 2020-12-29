// Code generated by github.com/vektah/dataloaden, DO NOT EDIT.

package model

import (
	"sync"
	"time"
)

// LimitSpeedLineLoaderConfig captures the config to create a new LimitSpeedLineLoader
type LimitSpeedLineLoaderConfig struct {
	// Fetch is a method that provides the data for the loader
	Fetch func(keys []string) ([]*LimitSpeedLine, []error)

	// Wait is how long wait before sending a batch
	Wait time.Duration

	// MaxBatch will limit the maximum number of keys to send in one batch, 0 = not limit
	MaxBatch int
}

// NewLimitSpeedLineLoader creates a new LimitSpeedLineLoader given a fetch, wait, and maxBatch
func NewLimitSpeedLineLoader(config LimitSpeedLineLoaderConfig) *LimitSpeedLineLoader {
	return &LimitSpeedLineLoader{
		fetch:    config.Fetch,
		wait:     config.Wait,
		maxBatch: config.MaxBatch,
	}
}

// LimitSpeedLineLoader batches and caches requests
type LimitSpeedLineLoader struct {
	// this method provides the data for the loader
	fetch func(keys []string) ([]*LimitSpeedLine, []error)

	// how long to done before sending a batch
	wait time.Duration

	// this will limit the maximum number of keys to send in one batch, 0 = no limit
	maxBatch int

	// INTERNAL

	// lazily created cache
	cache map[string]*LimitSpeedLine

	// the current batch. keys will continue to be collected until timeout is hit,
	// then everything will be sent to the fetch method and out to the listeners
	batch *limitSpeedLineLoaderBatch

	// mutex to prevent races
	mu sync.Mutex
}

type limitSpeedLineLoaderBatch struct {
	keys    []string
	data    []*LimitSpeedLine
	error   []error
	closing bool
	done    chan struct{}
}

// Load a LimitSpeedLine by key, batching and caching will be applied automatically
func (l *LimitSpeedLineLoader) Load(key string) (*LimitSpeedLine, error) {
	return l.LoadThunk(key)()
}

// LoadThunk returns a function that when called will block waiting for a LimitSpeedLine.
// This method should be used if you want one goroutine to make requests to many
// different data loaders without blocking until the thunk is called.
func (l *LimitSpeedLineLoader) LoadThunk(key string) func() (*LimitSpeedLine, error) {
	l.mu.Lock()
	if it, ok := l.cache[key]; ok {
		l.mu.Unlock()
		return func() (*LimitSpeedLine, error) {
			return it, nil
		}
	}
	if l.batch == nil {
		l.batch = &limitSpeedLineLoaderBatch{done: make(chan struct{})}
	}
	batch := l.batch
	pos := batch.keyIndex(l, key)
	l.mu.Unlock()

	return func() (*LimitSpeedLine, error) {
		<-batch.done

		var data *LimitSpeedLine
		if pos < len(batch.data) {
			data = batch.data[pos]
		}

		var err error
		// its convenient to be able to return a single error for everything
		if len(batch.error) == 1 {
			err = batch.error[0]
		} else if batch.error != nil {
			err = batch.error[pos]
		}

		if err == nil {
			l.mu.Lock()
			l.unsafeSet(key, data)
			l.mu.Unlock()
		}

		return data, err
	}
}

// LoadAll fetches many keys at once. It will be broken into appropriate sized
// sub batches depending on how the loader is configured
func (l *LimitSpeedLineLoader) LoadAll(keys []string) ([]*LimitSpeedLine, []error) {
	results := make([]func() (*LimitSpeedLine, error), len(keys))

	for i, key := range keys {
		results[i] = l.LoadThunk(key)
	}

	limitSpeedLines := make([]*LimitSpeedLine, len(keys))
	errors := make([]error, len(keys))
	for i, thunk := range results {
		limitSpeedLines[i], errors[i] = thunk()
	}
	return limitSpeedLines, errors
}

// LoadAllThunk returns a function that when called will block waiting for a LimitSpeedLines.
// This method should be used if you want one goroutine to make requests to many
// different data loaders without blocking until the thunk is called.
func (l *LimitSpeedLineLoader) LoadAllThunk(keys []string) func() ([]*LimitSpeedLine, []error) {
	results := make([]func() (*LimitSpeedLine, error), len(keys))
	for i, key := range keys {
		results[i] = l.LoadThunk(key)
	}
	return func() ([]*LimitSpeedLine, []error) {
		limitSpeedLines := make([]*LimitSpeedLine, len(keys))
		errors := make([]error, len(keys))
		for i, thunk := range results {
			limitSpeedLines[i], errors[i] = thunk()
		}
		return limitSpeedLines, errors
	}
}

// Prime the cache with the provided key and value. If the key already exists, no change is made
// and false is returned.
// (To forcefully prime the cache, clear the key first with loader.clear(key).prime(key, value).)
func (l *LimitSpeedLineLoader) Prime(key string, value *LimitSpeedLine) bool {
	l.mu.Lock()
	var found bool
	if _, found = l.cache[key]; !found {
		// make a copy when writing to the cache, its easy to pass a pointer in from a loop var
		// and end up with the whole cache pointing to the same value.
		cpy := *value
		l.unsafeSet(key, &cpy)
	}
	l.mu.Unlock()
	return !found
}

// Clear the value at key from the cache, if it exists
func (l *LimitSpeedLineLoader) Clear(key string) {
	l.mu.Lock()
	delete(l.cache, key)
	l.mu.Unlock()
}

func (l *LimitSpeedLineLoader) unsafeSet(key string, value *LimitSpeedLine) {
	if l.cache == nil {
		l.cache = map[string]*LimitSpeedLine{}
	}
	l.cache[key] = value
}

// keyIndex will return the location of the key in the batch, if its not found
// it will add the key to the batch
func (b *limitSpeedLineLoaderBatch) keyIndex(l *LimitSpeedLineLoader, key string) int {
	for i, existingKey := range b.keys {
		if key == existingKey {
			return i
		}
	}

	pos := len(b.keys)
	b.keys = append(b.keys, key)
	if pos == 0 {
		go b.startTimer(l)
	}

	if l.maxBatch != 0 && pos >= l.maxBatch-1 {
		if !b.closing {
			b.closing = true
			l.batch = nil
			go b.end(l)
		}
	}

	return pos
}

func (b *limitSpeedLineLoaderBatch) startTimer(l *LimitSpeedLineLoader) {
	time.Sleep(l.wait)
	l.mu.Lock()

	// we must have hit a batch limit and are already finalizing this batch
	if b.closing {
		l.mu.Unlock()
		return
	}

	l.batch = nil
	l.mu.Unlock()

	b.end(l)
}

func (b *limitSpeedLineLoaderBatch) end(l *LimitSpeedLineLoader) {
	b.data, b.error = l.fetch(b.keys)
	close(b.done)
}

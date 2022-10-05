package Cache

import (
	"fmt"
	"sync"
)

type LocalCache struct {
	mu           sync.RWMutex
	phoneNumbers map[string]bool
}

func NewLocalCache() *LocalCache {
	lc := &LocalCache{
		phoneNumbers: make(map[string]bool),
	}
	return lc
}

func (lc *LocalCache) Add(u string) {
	fmt.Println("adding")
	lc.mu.Lock()
	defer lc.mu.Unlock()
	lc.phoneNumbers[u] = true
}

func (lc *LocalCache) Get(id string) bool {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	_, ok := lc.phoneNumbers[id]
	if !ok {
		return false
	}
	return true
}

func (lc *LocalCache) Delete(id string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	delete(lc.phoneNumbers, id)
}

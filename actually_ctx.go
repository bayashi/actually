package actually

import (
	"runtime"
	"strconv"
	"strings"
	"sync"
)

// global context of actually
var aCtx = &aContext{
	failNowStates: make(map[int64]bool),
}

type aContext struct {
	mu            sync.RWMutex
	failNowStates map[int64]bool // [goroutineID]bool
}

func (ac *aContext) failNowOn() {
	ac.mu.Lock()
	defer ac.mu.Unlock()
	ac.failNowStates[goroutineID()] = true
}

func (ac *aContext) failNotNow() {
	ac.mu.Lock()
	defer ac.mu.Unlock()
	delete(ac.failNowStates, goroutineID())
}

func (ac *aContext) failNowState() bool {
	ac.mu.RLock()
	defer ac.mu.RUnlock()
	return ac.failNowStates[goroutineID()]
}

func goroutineID() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.ParseInt(idField, 10, 64)
	if err != nil {
		panic(err)
	}

	return id
}

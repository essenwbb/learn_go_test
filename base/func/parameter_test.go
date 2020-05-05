package _func

import (
	"sync"
	"testing"
)

func sum(ops ...int) int {
	s := 0
	for _, v := range ops {
		s += v
	}
	return s
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3))
}

func TestDefer(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Logf("recover from %s", err)
		}
	}()
	t.Log("start")
	panic("Fatal error")
}

func TestShareMemRwLock(t *testing.T) {
	t.Helper()
	counter := 0
	var rwMut sync.RWMutex
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			rwMut.Lock()
			defer func() {
				rwMut.Unlock()
				wg.Done()
			}()
			counter++
		}()
	}
	wg.Wait()
	t.Log(counter)
}

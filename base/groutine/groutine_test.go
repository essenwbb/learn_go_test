package groutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(i int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("the result is from %d", i)

}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ch <- runTask(i)
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Logf("The begin NumGoroutine %d", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Millisecond * 100)
	t.Logf("The after NumGoroutine %d", runtime.NumGoroutine())
}

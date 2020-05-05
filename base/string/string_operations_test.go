package string_operations

import (
	"testing"
)

func TestZhSupport(t *testing.T) {
	t.Helper()
	ZhString := "黑化肥会挥发"
	t.Logf("len(string): %d", len(ZhString))

	t.Logf("len([]rune): %d", len([]rune(ZhString)))
	for i, v := range []rune(ZhString) {
		t.Logf("%d:%x", i, v)
	}
}

package _map

import (
	"encoding/json"
	"testing"
)

func TestFuncParameter(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2}
	byte1, _ := json.Marshal(m1)
	t.Log(string(byte1))
}

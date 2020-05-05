package fib

import (
	"fmt"
	"testing"
)

func Test_fibonacci(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		f := Fibonacci()
		fmt.Println(arrayNLinesFileContent(f, 40))
	})
}

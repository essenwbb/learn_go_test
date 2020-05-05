package fib

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Fibonacci() IntGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type IntGen func() int

func (i IntGen) Read(p []byte) (n int, err error) {
	next := i()
	sprints := fmt.Sprintf("%d\n", next)
	// TODO: incorrect if p is too small!
	return strings.NewReader(sprints).Read(p)
}

func arrayNLinesFileContent(reader io.Reader, n int) []interface{} {
	scanner := bufio.NewScanner(reader)
	res := make([]interface{}, 0)
	for i := 0; i < n; i++ {
		scanner.Scan()
		res = append(res, scanner.Text())
	}
	return res
}

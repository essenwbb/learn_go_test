package main

import (
	"bufio"
	"fmt"
	"learb_go_test/base/functional/fib"
	"learb_go_test/base/util"
	"os"
)

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	util.PanicErr(err)
	defer util.DeferPanicErr(file.Close)

	writer := bufio.NewWriter(file)
	defer util.DeferPanicErr(writer.Flush)

	f := fib.Fibonacci()
	for i := 0; i < 30; i++ {
		_, err = fmt.Fprintln(writer, f())
		util.PanicErr(err)
	}
}

func main() {
	writeFile("fib.txt")
}

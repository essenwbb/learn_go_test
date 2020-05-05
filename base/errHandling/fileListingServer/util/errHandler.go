package util

import (
	"fmt"
	"os"
)

func PanicErr(err error) {
	if err != nil {
		if pathError, ok := err.(*os.PathError); ok {
			fmt.Printf("Error: %s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		} else {
			panic(err)
		}
	}
}

func DeferPanicErr(f func() error) {
	err := f()
	PanicErr(err)
}

package util

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func DeferPanicErr(f func() error) {
	err := f()
	PanicErr(err)
}

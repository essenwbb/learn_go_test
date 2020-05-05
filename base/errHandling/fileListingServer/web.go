package main

import (
	"github.com/essenwbb/learb_go_test/base/errHandling/fileListingServer/fileListing"
	"github.com/essenwbb/learb_go_test/base/errHandling/fileListingServer/util"
	"net/http"
	"os"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}

}

func main() {
	prefix := "/list/"
	http.HandleFunc(prefix, errWrapper(fileListing.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	util.PanicErr(err)
}
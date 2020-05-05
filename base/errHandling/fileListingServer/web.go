package main

import (
	"github.com/essenwbb/learb_go_test/base/errHandling/fileListingServer/fileListing"
	"github.com/essenwbb/learb_go_test/base/errHandling/fileListingServer/util"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type userError interface {
	error
	Message() string
}

type appHandler func(http.ResponseWriter, *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {

			if userErr, ok := err.(userError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}

			log.Warnf("Error Handling request: %s", err.Error())
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

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//util.PanicErr(err)
	//log.Infof("Current path: %s", dir)

	http.HandleFunc("/", errWrapper(fileListing.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	util.PanicErr(err)
}

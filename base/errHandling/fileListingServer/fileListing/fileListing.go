package fileListing

import (
	"github.com/essenwbb/learb_go_test/base/errHandling/fileListingServer/util"
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter, request *http.Request) (err error) {
	prefix := "/list/"
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer util.DeferPanicErr(file.Close)

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	_, err = writer.Write(all)
	if err != nil {
		return
	}
	return nil
}

package fileListing

import (
	"github.com/essenwbb/learb_go_test/base/errHandling/fileListingServer/util"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	prefix    = "/list/"
	sharedDir = "/Users/wubinbin/go/src/github.com/essenwbb/learb_go_test/"
)

type userError string

func (u userError) Error() string {
	return u.Message()
}

func (u userError) Message() string {
	return string(u)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) (err error) {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must start with" + prefix)
	}

	path := request.URL.Path[len(prefix):]
	file, err := os.Open(sharedDir + path)
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

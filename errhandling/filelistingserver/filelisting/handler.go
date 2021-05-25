package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter,
	request *http.Request) error {
	//check if url has prefix
	if strings.Index(
		request.URL.Path, prefix) != 0 {
		return userError("path must start " +
			"with " + prefix)
	}
	path := request.URL.Path[len(prefix):] // /list/xxxxx ->  xxxxx
	file, err := os.Open(path)
	if err != nil {
		//panic(err)

		//not good, users will get confused: what happened???
		//do as below!

		//http.Error(writer,
		//	err.Error(),
		//	http.StatusInternalServerError)

		//well. then users can know what's wrong is it,
		//but it's not safe to give users our internal error messages,
		//so we need to pack them!(see errWrapper in web.go)

		//now we let this func return err to the errWrapper to deal with them
		return err

	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}

	writer.Write(all)
	return nil
}

package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter,
	request *http.Request) error {
	path := request.URL.Path[len("/list/"):] // /list/xxxxx ->  xxxxx
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

package main

import (
	"goLearning/errhandling/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	//deal with all errors here!
	return func(writer http.ResponseWriter,
		request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error occurred "+
				"handling request: %s", err.Error())
			if userError, ok := err.(userError); ok {
				http.Error(writer,
					userError.Message(),
					http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				//statusText is a helper func to get some easy msgs(like here is Not Found)
				http.StatusText(code), //wrap the internal error to a msg here, not err.Error() anymore!
				code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	//url here has a potential risk(if here is / and in handler is /list/, then it will cause slice out of range)
	//so we need to deal with it by our defer & recover (let it in errWrapper)
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

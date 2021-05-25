package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

//the name of the test files should be xxx_test.go

// 2 ways to test it
// 1. mock request & response writer (quickly &test granularity is small like a unit test)
// 2. create a server to test it (slow but the  integration level & coverage is high)

func errPanic(writer http.ResponseWriter,
	request *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter,
	request *http.Request) error {
	return testingUserError("a user error")
}

func errNotFound(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter,
	request *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter,
	request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "a user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

//way 1(mock)
func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.google.com", nil)
		f(response, request)

		//response is a *ResponseRecorder, but we can use .Result() to get the *Response
		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

//way2(server)
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)

		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response,
	expectedCode int, expectedMessage string,
	t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode ||
		body != expectedMessage {
		t.Errorf("expected (%d, %s), "+
			"got (%d, %s)",
			expectedCode, expectedMessage,
			resp.StatusCode, body)
	}
}

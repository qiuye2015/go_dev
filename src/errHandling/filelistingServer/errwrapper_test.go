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

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()

}
func (e testingUserError) Message() string {
	return string(e)
}

func errUeserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}
func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}
func errNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}
func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown error")
}
func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

// 表格驱动测试
var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUeserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.imooc.com",
			nil)
		f(response, request)
		vertifyResponse(t,response.Result(),tt.code,tt.message)
	}

}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt:= range tests{
		f:= errWrapper(tt.h)
		server:= httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)

		vertifyResponse(t,response,tt.code,tt.message)

	}
}

func vertifyResponse(t *testing.T, response *http.Response,
	exceptedCode int, exceptedMessage string) {
	b, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(b), "\n")
	if response.StatusCode != exceptedCode ||
		body !=exceptedMessage {
		t.Errorf("except (%d, %s); "+
			"got (%d, %s)",
			exceptedCode, exceptedMessage,
			response.StatusCode, body)
	}
}

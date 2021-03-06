package packr

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_HTTPBox(t *testing.T) {
	r := require.New(t)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(testBox))

	req, err := http.NewRequest("GET", "/hello.txt", nil)
	r.NoError(err)

	res := httptest.NewRecorder()

	mux.ServeHTTP(res, req)

	r.Equal(200, res.Code)
	r.Equal("hello world!", strings.TrimSpace(res.Body.String()))
}

// func Test_HTTPBox_CaseInsensitive(t *testing.T) {
//
// 	mux := http.NewServeMux()
// 	testBox.AddString("myfile.txt", "this is my file")
// 	mux.Handle("/", http.FileServer(testBox))
//
// 	for _, path := range []string{"/MyFile.txt", "/myfile.txt", "/Myfile.txt"} {
// 		t.Run(path, func(st *testing.T) {
// 			r := require.New(st)
//
// 			req, err := http.NewRequest("GET", path, nil)
// 			r.NoError(err)
//
// 			res := httptest.NewRecorder()
//
// 			mux.ServeHTTP(res, req)
//
// 			r.Equal(200, res.Code)
// 			r.Equal("this is my file", strings.TrimSpace(res.Body.String()))
// 		})
// 	}
// }

func Test_HTTPBox_NotFound(t *testing.T) {
	r := require.New(t)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(testBox))

	req, err := http.NewRequest("GET", "/notInBox.txt", nil)
	r.NoError(err)

	res := httptest.NewRecorder()

	mux.ServeHTTP(res, req)

	r.Equal(404, res.Code)
}

func Test_HTTPBox_Handles_IndexHTML(t *testing.T) {
	r := require.New(t)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(testBox))

	req, err := http.NewRequest("GET", "/", nil)
	r.NoError(err)

	res := httptest.NewRecorder()

	mux.ServeHTTP(res, req)

	r.Equal("<h1>Index!</h1>", strings.TrimSpace(res.Body.String()))
}

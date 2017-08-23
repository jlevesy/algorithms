package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/jlevesy/algorithm/strings"
)

func hasUniqueHandler(w http.ResponseWriter, r *http.Request) {
	in := r.URL.Query().Get("str")
	res := strings.HasUniqueChars(in)
	w.Write([]byte(fmt.Sprintf("%v", res)))
}

func hasUnique2Handler(w http.ResponseWriter, r *http.Request) {
	in := r.URL.Query().Get("str")
	fmt.Println("HEY I HAVE A REQUEST WITH ARG:", in)
	res := strings.HasUniqueChars2(in)
	w.Write([]byte(fmt.Sprintf("%v", res)))
}

func compressHandler(w http.ResponseWriter, r *http.Request) {
	in := r.URL.Query().Get("str")
	res := strings.CompressString([]rune(in))
	w.Write([]byte(string(res)))
}

func main() {
	http.HandleFunc("/v1/has-unique", hasUniqueHandler)
	http.HandleFunc("/v1/has-unique-2", hasUnique2Handler)
	http.HandleFunc("/v1/compress", compressHandler)
	http.ListenAndServe(":8080", nil)
}

package server

import (
	"net/http"
)

var stdSrv *http.Server

func StartStd(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world\n"))
	})

	stdSrv = &http.Server{Addr: addr, Handler: mux}
	return stdSrv.ListenAndServe()
}

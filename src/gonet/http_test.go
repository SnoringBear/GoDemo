package gonet

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHttp01(t *testing.T) {
	http.HandleFunc("/handler1", handler1)
	http.HandleFunc("/handler2", handler2)

	// 使用 HTTP/2 会自动启用多路复用
	http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	// epoll   kqueue
	// epoll 采用的是事件驱动模型，而非轮询模型。它能够在大量的文件描述符中只处理那些发生了事件的文件描述符，因此在大规模并发时性能更好
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is handler 1")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is handler 2")
}

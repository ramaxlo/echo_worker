package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net"
	"net/http"
	"os"
)

func echo_handler(ws *websocket.Conn) {
	fmt.Printf("Conn: %s\n", ws.RemoteAddr().String())
	io.Copy(ws, ws)
}

func main() {
	file := os.NewFile(uintptr(3), "socket")
	ln, err := net.FileListener(file)
	if err != nil {
		fmt.Println("Fail to create listener")
		return
	}

	http.Handle("/echo", websocket.Handler(echo_handler))
	fmt.Println("Start to serve")
	err = http.Serve(ln, nil)
	if err != nil {
		fmt.Println("Fail to serve")
		return
	}

	fmt.Println("Serve done")
}

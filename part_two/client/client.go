package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

func main() {
	log.Println("starting server")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}
	for {
		if conn, err := ln.Accept(); err == nil {
			go handleConnection(conn)

		}
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		req, err := http.ReadRequest(reader)
		if err != nil {
			if err != io.EOF {
				log.Printf("Failed to read request: %s", err)
			}
			return
		}
		if be, err := net.Dial("tcp", "127.0.0.1:8081"); err == nil {
			be_reader := bufio.NewReader(be)
			if err := req.Write(be); err == nil {
				if resp, err := http.ReadResponse(be_reader, req); err == nil {

					resp.Close = true
					if err := resp.Write(conn); err == nil {

						fmt.Printf("%s: %d", req.URL.Path, resp.StatusCode)
					}
				}
			}
		}
	}

}

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

type StreamServer struct{}

func (ss *StreamServer) ConnectAndReadFile() {
	ln, err := net.Listen("tcp", ":7777")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go ss.Process(conn)
	}
}

func (ss *StreamServer) Process(conn net.Conn) {
	buf := new(bytes.Buffer)
	for {
		var size int64
		binary.Read(conn, binary.LittleEndian, &size)

		qtdBytes, err := io.CopyN(buf, conn, size)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(buf.Bytes())
		fmt.Println("Received", qtdBytes, "bytes from client")
		break
	}
}

func main() {
	ss := &StreamServer{}
	ss.ConnectAndReadFile()
}

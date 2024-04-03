package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func main() {
	file := make([]byte, 2048000000000)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		panic(err)
	}

	conn, err := net.Dial("tcp", ":7777")
	if err != nil {
		panic(err)
	}

	binary.Write(conn, binary.LittleEndian, int64(len(file)))

	qtsBytes, err := io.CopyN(conn, bytes.NewReader(file), int64(len(file)))
	if err != nil {
		panic(err)
	}
	fmt.Println("Sent", qtsBytes, "bytes to server")
}

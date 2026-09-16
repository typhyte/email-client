package main

// client code from https://www.golinuxcloud.com/golang-tcp-server-client/
import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	fmt.Println("Starting!")
	const addr = "192.168.17.191:25"
	// connection
	// TODO: figure out how to connect to the smtp server.
	//       this current function *can* connect, but it
	//       *can't* actually authenticate. 554 error.
	conn, err := net.Dial("tcp", addr)
	fmt.Println("Connecting...")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// initial message
	_, err = fmt.Fprintf(conn, "hello")
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Print(string(buf[:n]))
}

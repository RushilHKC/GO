package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	//write to the server
	fmt.Fprintf(conn, "GET /index.html\n")

	//Read the response from the server
	bs := make([]byte, 1024)
	n, er := conn.Read(bs)

	if er != nil {
		log.Fatal(er)
	}

	fmt.Println(string(bs[:n]))

	

}

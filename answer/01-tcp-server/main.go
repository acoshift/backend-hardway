package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatal("tcp listen error:", err)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal("tcp accept error:", err)
		}

		go func() {
			defer conn.Close()
			reader := bufio.NewReader(conn)
			for {
				s, _, err := reader.ReadLine()
				if err != nil {
					log.Printf("connection: %v\n", err)
					return
				}
				log.Printf("receive: %s\n", s)
				fmt.Fprintf(conn, "server: %s\n", s)
			}
		}()
	}
}

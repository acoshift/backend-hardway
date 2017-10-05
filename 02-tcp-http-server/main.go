package main

import (
	"bufio"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:3333")
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
					log.Print("connection:", err)
					return
				}
				log.Printf("receive: %s", s)
				if len(s) == 0 {
					p := "HTTP/1.1 200 OK\n"
					p += "Date: " + time.Now().Format(http.TimeFormat) + "\n"
					p += "Connection: close\n"
					p += "Content-Type: text/html\n"
					p += "\n"
					p += "<h1>Hello from TCP Server</h1><p>:P</p>"
					conn.Write([]byte(p))
					conn.Close()
				}
			}
			conn.Write([]byte("Received"))
		}()
	}
}

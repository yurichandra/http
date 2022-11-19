package http

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"strings"
)

var methodMap = map[string]bool{
	"GET":    true,
	"POST":   true,
	"PATCH":  true,
	"DELETE": true,
	"PUT":    true,
}

func Serve(port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Println(err)
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			return err
		}

		go func() {
			defer conn.Close()
			scanner := bufio.NewScanner(conn)

			for scanner.Scan() {
				text := scanner.Text()

				// Separate strings by white space character.
				texts := strings.Fields(text)

				var method string
				if len(texts) > 0 {
					method = texts[0]
				}

				var response, finalResponse string

				if len(texts) > 0 && methodMap[method] {
					if method == "POST" {
						response = base64.StdEncoding.EncodeToString([]byte(text))
					} else {
						response = "Hello, Gunners!"
					}

					finalResponse = fmt.Sprintf("HTTP/1.1 200 OK \r\n"+"Content-Length: %d\r\n"+"Content-Type: text/html\r\n"+"\r\n"+"%s", len(response), response)
					conn.Write([]byte(finalResponse))
				}
			}
		}()
	}
}

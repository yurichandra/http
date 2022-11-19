package main

import (
	"github.com/yurichandra/http/http"
	"log"
)

func main() {
	if err := http.Serve(":8000"); err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var Version = ""

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world, version: '%s', time: '%s'", Version, time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
	fmt.Println(Version)
	var port int
	flag.IntVar(&port, "p", 8080, "端口号,默认为8080")
	flag.Parse()
	fmt.Println("监控在端口", port)
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("错误: ", err)
	}
}

package main

import (
	"fmt"
	"net/http"
	"time"
)

func printTime(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	fmt.Fprintf(w, "time is %s\n", t.Format("2006-01-02 15:04:05"))
}

func printHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", printTime)
	http.HandleFunc("/hello", printHello)

	fmt.Println("about to start listening on port 8000")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

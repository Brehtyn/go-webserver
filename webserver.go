package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/bunny", bunnyHandler)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))
	http.ListenAndServe("localhost:8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> You are ready to Go </h1>")
	fmt.Fprintf(w, "<a href='bunny'> What is Go Bunny </a>")
}

func bunnyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Go Bunny </h1>")
	fmt.Fprintf(w, "<p> Bigitals Logo, iconic for me! </p>")
	fmt.Fprintf(w, "<img src='img/simple_logo.png' />")
}

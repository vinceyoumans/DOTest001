package main

import "fmt"

import "net/http"

import "log"

func main() {
	fmt.Println("GoDocker")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello World")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}


package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Helloo......")
}

func Bye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Byeee.....")
}

func main() {

	fmt.Println("starting the server at address :8080")

	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/bye", Bye)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Errorf("opps got an error %w", err)
	}

}

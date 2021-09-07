package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request)  {
	fmt.Println("Hello")

	_, err := w.Write([]byte("Hello"))
	if err != nil {
		return 
	}
}

func main() {
	http.HandleFunc("/", home)
	_ = http.ListenAndServe(":3000", nil)
}
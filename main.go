package main

import (
	"os"
	"fmt"
	"net/http"
)

type Env struct {
  UserName string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	  user := os.Getenv("USER_NAME")
		fmt.Fprintf(w, "HELLO, %v\n", user)
	})
	http.ListenAndServe(":80", nil)
}

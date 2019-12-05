// zor performs authentication and authorization for the application
package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/auth", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

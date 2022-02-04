// webserverjstest
package main

import (
	"fmt"
	"net/http"
	"log"
)

type Home struct {
}

func (h *Home) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello http api"))
}

func main() {
	fmt.Println("Start WebServer js test")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

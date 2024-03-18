// main.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cgi"
	"os"
)

func main() {
	fmt.Println("Starting server on http://0.0.0.0:1234")
	log.Fatal(http.ListenAndServe(":1234", &cgi.Handler{
		Path: os.Getenv("CGI_HANDLER"),
		Dir:  os.Getenv("DOCUMENT_ROOT"),

		InheritEnv: []string{"SCRIPT_FILENAME", "REDIRECT_STATUS"},
	}))
}
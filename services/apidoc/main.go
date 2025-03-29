package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	scalar "github.com/MarceloPetrucio/go-scalar-api-reference"
)

//go:embed openapi.yaml
var apiDoc string

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", apidoc)

	server := http.Server{
		Addr:              os.Getenv("ADDR"),
		Handler:           router,
		ReadTimeout:       3 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       10 * time.Second,
	}

	log.Println("starting api")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func apidoc(w http.ResponseWriter, _ *http.Request) {
	htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
		DarkMode:    true,
		SpecContent: apiDoc,
		ShowSidebar: true,
		CustomOptions: scalar.CustomOptions{
			PageTitle: os.Getenv("PAGE_TITLE"),
		},
	})

	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Fprintln(w, htmlContent)
}

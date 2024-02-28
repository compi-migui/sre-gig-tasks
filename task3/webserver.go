package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/pkg/browser"
)


func main() {
	var wg sync.WaitGroup

	workdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(1)
	go func(path string) {
		defer wg.Done()
		log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(path))))
	}(workdir)

	browser.OpenURL("localhost:8080")
	wg.Wait()

}

package main

// look up lisp doc via html2text

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jaytaylor/html2text"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s TOPIC", os.Args[0])
	}
	url := "http://l1sp.org/" + os.Args[1]
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	text, err := html2text.FromString(string(contents))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text)
}

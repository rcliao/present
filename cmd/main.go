package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/russross/blackfriday.v2"
)

func main() {
	content, err := ioutil.ReadFile("./slides.md")
	if err != nil {
		panic(err)
	}
	markdownContent := string(parse(content))
	http.HandleFunc("/", render(markdownContent))
	http.ListenAndServe(":9000", nil)
}

func render(content string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, content)
	}
}

func parse(content []byte) []byte {
	output := blackfriday.Run(content)
	return output
}

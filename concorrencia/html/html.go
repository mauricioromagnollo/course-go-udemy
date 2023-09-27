package html

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

// Obtem um título de uma página HTML via requisição HTTP <title>
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1] // 1 é o primeiro grupo de captura
		}(url) // IIFE
	}

	return c
}

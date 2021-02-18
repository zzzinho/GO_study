package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gocolly/colly"
)

func generateFormData() map[string][]byte {
	f, _ := os.Open("gocolly.jpg")
	defer f.Close()

	imgData, _ := ioutil.ReadAll(f)

	return map[string][]byte{
		"firstname": []byte("one"),
		"lastname":  []byte("two"),
		"email":     []byte("onetwo@example.com"),
		"file":      imgData,
	}
}

func setupServer() {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10000000)
		if err != nil {
			fmt.Println("server: Error")
			w.WriteHeader(500)
			w.Write([]byte("<html><body>Internal Server Error</body></html>"))
			return
		}
		w.WriteHeader(200)
		fmt.Println("server: OK")
		w.Write([]byte("<html><body>Success</body></html>"))
	}

	go http.ListenAndServe(":8080", handler)
}

func main() {
	setupServer()

	c := colly.NewCollector(colly.AllowURLRevisit(), colly.MaxDepth(5))

	c.OnHTML("html", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		time.Sleep(1 * time.Second)
		e.Request.PostMultipart("http://localhost:8080/", generateFormData())
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Posting gocolly.jpb to", r.URL.String())
	})
	c.PostMultipart("http://localhost:8080/", generateFormData())
	c.Wait()
}

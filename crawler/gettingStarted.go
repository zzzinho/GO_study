package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// 요청 전에 호출
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	// error 발생시
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Soemthing went wrong:", err)
	})
	// response header 수신 후
	c.OnResponseHeaders(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	// response 수신 후
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Reqeuset.URL)
	})
	// OnResponse 직후 받은 content가 HTML일 때
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println("First column of a table row:", e.Text)
	})
	// OnHTML 직후 받은 content가 HTML 또는 XML 일 때
	c.OnXML("//h1", func(e *colly.XMLElement) {
		fmt.Println(e.Text)
	})
	// OnXML 콜백 후
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

}

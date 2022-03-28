package main

import (
	"flag"
	"fmt"
	"github.com/kevin-glare/hardcode-dev-go/hw2/pkg/crawler"
	"github.com/kevin-glare/hardcode-dev-go/hw2/pkg/crawler/spider"
	"strings"
)

var (
	urls  = [2]string{"https://go.dev", "https://golang.org"}
	depth = 2
	s string
)

func main() {
	parseFlag()
	data := scanURLs()
	if len(data) == 0 || len(s) == 0 {
		fmt.Println(data)
		return
	}

	result := scanData(data)
	fmt.Print(result)

}
func parseFlag() {
	flag.StringVar(&s, "s", "", "search query")
	flag.Parse()

	if len(s) > 0 {
		s = strings.ToLower(s)
	}
}

func scanURLs() []crawler.Document {
	var data []crawler.Document
	spider := spider.New()

	for _, url := range urls {
		result, err := spider.Scan(url, depth)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		data = append(data, result...)
	}

	return data
}

func scanData(data []crawler.Document) []crawler.Document {
	var result []crawler.Document

	for _, site := range data {
		title := strings.ToLower(site.Title)

		if strings.Contains(title, s) {
			result = append(result, site)
		}
	}

	return result
}

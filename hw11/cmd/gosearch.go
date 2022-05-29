package main

import (
	"fmt"
	"github.com/kevin-glare/hardcode-dev-go/hw11/pkg/netsrv"
	"log"
	"sort"

	"github.com/kevin-glare/hardcode-dev-go/hw11/pkg/crawler"
	"github.com/kevin-glare/hardcode-dev-go/hw11/pkg/crawler/spider"
	"github.com/kevin-glare/hardcode-dev-go/hw11/pkg/index/hash"
)

var (
	urls  = [2]string{"https://go.dev", "https://golang.org"}
	depth = 1
)

func main() {
	index := hash.New()
	var docs []crawler.Document

	docs = parsingURLs(index)
	sort.SliceStable(docs, func(i, j int) bool {
		return docs[i].ID < docs[j].ID
	})

	inCh := make(chan string)
	outCh := make(chan []crawler.Document)

	server, err := netsrv.NewServer(inCh, outCh)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer server.Close()

	server.Run()

	results(index, docs, inCh, outCh)
}

func results(index *hash.Index, docs []crawler.Document, inCh chan string, outCh chan []crawler.Document) {
	for {
		for query := range inCh {
			findDoc := make([]crawler.Document, 0)
			indices := index.Search(query)
			for _, val := range indices {
				for _, doc := range docs {
					if doc.ID == val {
						findDoc = append(findDoc, doc)
						break
					}
				}
			}

			outCh <- findDoc
		}
	}
}


func parsingURLs(index *hash.Index) []crawler.Document {
	spider := spider.New()
	counter := 0

	var data []crawler.Document

	for _, url := range urls {
		result, err := spider.Scan(url, depth)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		for i, _ := range result {
			result[i].ID = counter
			counter++
		}

		data = append(data, result...)
		index.Add(result)
	}

	return data
}

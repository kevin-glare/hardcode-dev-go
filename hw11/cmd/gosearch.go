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
	index *hash.Index
	docs  []crawler.Document
)

func main() {
	index = hash.New()
	docs = parsingURLs(index)
	sort.SliceStable(docs, func(i, j int) bool {
		return docs[i].ID < docs[j].ID
	})

	server, err := netsrv.NewServer(results)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer server.Close()

	err = server.Run()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func results(query string) []crawler.Document {
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

	return findDoc
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

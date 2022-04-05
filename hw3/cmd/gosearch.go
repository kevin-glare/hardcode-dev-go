package main

import (
	"fmt"
	"github.com/kevin-glare/hardcode-dev-go/hw3/pkg/crawler"
	"github.com/kevin-glare/hardcode-dev-go/hw3/pkg/crawler/spider"
	"github.com/kevin-glare/hardcode-dev-go/hw3/pkg/index/hash"
	"math/rand"
	"sort"
)

var (
	urls  = [2]string{"https://go.dev", "https://golang.org"}
	depth = 2
	ids   = make(map[int]int)
)

const (
	N = 100_000
)

func main() {
	index, docs := scanURLs()

	// The sort is not guaranteed to be stable.
	sort.SliceStable(docs, func(i, j int) bool {
		return docs[i].ID < docs[j].ID
	})

	var query string
	for {
		fmt.Print("Query: ")
		fmt.Scanf("%s\n", &query)

		if len(query) == 0 {
			return
		}

		indices := index.Search(query)

		printResult(indices, docs)
	}
}

func scanURLs() (*hash.Index, []crawler.Document) {
	var data []crawler.Document

	spider := spider.New()
	index := hash.New()

	for _, url := range urls {
		result, err := spider.Scan(url, depth)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		for i, _ := range result {
			result[i].ID = generationID()
		}

		data = append(data, result...)
		index.Add(result)
	}

	return index, data
}

func generationID() int {
	id := rand.Intn(N)
	if _, ok := ids[id]; ok {
		return generationID()
	}

	ids[id] = id
	return id
}

func printResult(indices []int, docs []crawler.Document) {
	docsLen := len(docs)

	for _, val := range indices {
		i := sort.Search(docsLen, func(i int) bool { return docs[i].ID == val })

		if i < docsLen && docs[i].ID == val {
			fmt.Println(docs[i])
		}
	}
}

package api

import (
	"fmt"
	"github.com/kevin-glare/hardcode-dev-go/hw5/pkg/crawler"
	"github.com/kevin-glare/hardcode-dev-go/hw5/pkg/crawler/spider"
	"github.com/kevin-glare/hardcode-dev-go/hw5/pkg/index/hash"
	"sort"
	"sync"
)

var (
	urls  = [2]string{"https://go.dev", "https://golang.org"}
	depth = 1
)

type Store struct {
	index *hash.Index

	sync.Mutex
	Docs  []crawler.Document
}
func NewStore() *Store {
	index := hash.New()

	docs := parsingURLs(index)
	sort.SliceStable(docs, func(i, j int) bool {
		return docs[i].ID < docs[j].ID
	})

	return &Store{
		index: index,
		Docs:  docs,
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

func (r *Store) FindByQuery(query string) []crawler.Document {
	findDoc := make([]crawler.Document, 0)
	indices := r.index.Search(query)

	for _, val := range indices {
		for _, doc := range r.Docs {
			if doc.ID == val {
				findDoc = append(findDoc, doc)
				break
			}
		}
	}

	return findDoc
}

func (r *Store) FindByID(id int) *crawler.Document {
	for i, d := range r.Docs {
		if d.ID == id {
			return &r.Docs[i]
		}
	}
	return nil
}

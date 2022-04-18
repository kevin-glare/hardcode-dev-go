package main

import (
	"encoding/json"
	"fmt"
	"github.com/kevin-glare/hardcode-dev-go/hw5/pkg/crawler"
	"github.com/kevin-glare/hardcode-dev-go/hw5/pkg/crawler/spider"
	"github.com/kevin-glare/hardcode-dev-go/hw5/pkg/index/hash"
	"io"
	"os"
	"sort"
)

var (
	urls  = [2]string{"https://go.dev", "https://golang.org"}
	depth = 2
)

func main() {
	index := hash.New()
	docPath := "./assets/docs.json"
	var docs []crawler.Document
	var err error

	fileExist := doesFileExist(docPath)

	if fileExist {
		docs, err = parsingDocument(index, docPath)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		docs = parsingURLs(index)
	}

	sort.SliceStable(docs, func(i, j int) bool {
		return docs[i].ID < docs[j].ID
	})

	if !fileExist {
		if err = saveDocsToFile(docs, docPath); err != nil {
			fmt.Println(err.Error())
			return
		}
	}

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

func parsingDocument(index *hash.Index, docPath string) ([]crawler.Document, error) {
	file, err := os.Open(docPath)

	defer file.Close()
	if err != nil {
		return nil, err
	}

	var docs []crawler.Document
	message, _ := io.ReadAll(file)
	err = json.Unmarshal([]byte(string(message)), &docs)
	if err != nil {
		return nil, err
	}

	index.Add(docs)

	return docs, err
}

func saveDocsToFile(docs []crawler.Document, docPath string) error {
	file, err := os.Create(docPath)
	defer file.Close()

	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(docs)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)

	return err
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

func printResult(indices []int, docs []crawler.Document) {
	for _, val := range indices {
		for _, doc := range docs {
			if doc.ID == val {
				fmt.Println(docs)
				break
			}
		}
	}
}

func doesFileExist(pathToFile string) bool {
	file, err := os.Open(pathToFile)
	defer file.Close()

	if err != nil {
		return false
	}

	return true
}

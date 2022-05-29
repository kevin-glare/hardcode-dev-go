package index

// Обратный индекс отсканированных документов.

import "github.com/kevin-glare/hardcode-dev-go/hw11/pkg/crawler"

// Interface определяет контракт службы индексирования документов.
type Interface interface {
	Add([]crawler.Document)
	Search(string) []int
}

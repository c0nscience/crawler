package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

type crawlerClient struct {
	url string
}

func New(url string) *crawlerClient {
	return &crawlerClient{
		url,
	}
}

func (me *crawlerClient) InStock(size string) bool {
	resp, err := http.Get(me.url)
	if err != nil {
		log.Error().Err(err)
		return false
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Error().Err(err)
		return false
	}

	el := doc.Find("li[data-availability]").
		FilterFunction(func(_ int, sel *goquery.Selection) bool {
			val, _ := sel.Attr("data-availability")
			return val == "available" || val == "back-in-stock"
		}).
		FilterFunction(func(i int, selection *goquery.Selection) bool {
			return strings.TrimSpace(selection.Text()) == size
		})
	return len(el.Nodes) > 0
}

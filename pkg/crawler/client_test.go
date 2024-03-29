package crawler_test

import (
	"crawler/pkg/crawler"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCrawlerClient_InStock_NewPage(t *testing.T) {

	tests := []struct {
		size    string
		inStock bool
	}{
		{"S 36/38", true},
		{"M 40/42", true},
		{"L 44/46", true},
		{"XL 48/50", true},
		{"XXL 52/54", true},
	}

	for _, test := range tests {
		v := ""
		if test.inStock {
			v = "be"
		} else {
			v = "not be"
		}
		t.Run(fmt.Sprintf("%s should %s in stock", test.size, v), func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				b, err := os.ReadFile("new-page.txt")
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.Write(b)
			}))
			defer svr.Close()

			cli := crawler.New(svr.URL)

			inStock := cli.InStock(test.size)

			assert.Equal(t, test.inStock, inStock)
		})
	}

}

func TestCrawlerClient_InStock_NewPage_Alternative(t *testing.T) {

	tests := []struct {
		size    string
		inStock bool
	}{
		{"S 36/38", true},
		{"M 40/42", true},
		{"L 44/46", false},
		{"XL 48/50", false},
		{"XXL 52/54", true},
	}

	for _, test := range tests {
		v := ""
		if test.inStock {
			v = "be"
		} else {
			v = "not be"
		}
		t.Run(fmt.Sprintf("%s should %s in stock", test.size, v), func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				b, err := os.ReadFile("new-page-alternative.txt")
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.Write(b)
			}))
			defer svr.Close()

			cli := crawler.New(svr.URL)

			inStock := cli.InStock(test.size)

			assert.Equal(t, test.inStock, inStock)
		})
	}

}

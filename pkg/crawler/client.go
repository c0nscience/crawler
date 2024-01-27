package crawler

import (
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"regexp"
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

	out, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err)
		return false
	}

	r := regexp.MustCompile(`<li.*?data-availability="available".*?> (` + size + `)</.*?</li>`)

	b := r.FindStringSubmatch(string(out))
	return b != nil
}

package main

import (
	"fmt"
	"net"
	"net/http"
	"regexp"
	"time"

	"github.com/gocolly/colly/v2"
)

var (
	startUrl        = "https://www.continente.pt/"
	urlFiltersRegex = regexp.MustCompile(`www\.continente`)
)

func main() {
	defer Timer("main")()

	c := colly.NewCollector(
		colly.Async(),
		colly.URLFilters(urlFiltersRegex),
	)
	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 10 * time.Second,
	})

	fmt.Println("starting...")

	NewExecutor(c).start(startUrl)
}

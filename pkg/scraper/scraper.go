package scraper

import (
	"github.com/chromedp/chromedp"
)

type Scraper interface {
	Scrape(actions ...chromedp.Action) error
}

func New() Scraper {
	return &chromedpScraper{}
}

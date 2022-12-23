package scraper

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
)

type chromedpScraper struct {
}

func (cdp *chromedpScraper) Scrape(actions ...chromedp.Action) error {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	return chromedp.Run(ctx, actions...)
}

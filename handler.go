package main

import (
	"github.com/gocolly/colly/v2"
)

type handler struct {
}

func (h *handler) RegisterHandlers(c *colly.Collector, s *Scrapper) {}

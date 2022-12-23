package controller

import (
	"webscrapper/pkg/scraper"

	"github.com/chromedp/chromedp"
)

type chromedpNodeController struct {
	nodes []Node
}

func newChromedpNodeController() Controller {
	return &chromedpNodeController{}
}

func (cdpc *chromedpNodeController) Scrape() error {
	return scraper.New().Scrape(cdpc.actions()...)
}

func (cdpc *chromedpNodeController) Click() Controller {
	cdpc.nodes = append(cdpc.nodes, newClickNode())
	return cdpc
}

func (cdpc *chromedpNodeController) Find(selector string) Controller {
	cdpc.nodes = append(cdpc.nodes, newFindNode(selector))
	return cdpc
}

func (cdpc *chromedpNodeController) Get(selector string) Controller {
	cdpc.nodes = append(cdpc.nodes, newGetNode(selector))
	return cdpc
}

func (cdpc *chromedpNodeController) Navigate(url string) Controller {
	cdpc.nodes = append(cdpc.nodes, newNavigateNode(url))
	return cdpc
}

func (cdpc *chromedpNodeController) Text(output *string) Controller {
	cdpc.nodes = append(cdpc.nodes, newTextNode(output))
	return cdpc
}

func (cdpc *chromedpNodeController) TextAll(output *[]string) Controller {
	cdpc.nodes = append(cdpc.nodes, newTextAllNode(output))
	return cdpc
}

func (cdpc *chromedpNodeController) actions() []chromedp.Action {
	grouper := newChromedpActionGrouper()
	for _, node := range cdpc.nodes {
		node.Accept(grouper)
	}
	return grouper.Actions()
}

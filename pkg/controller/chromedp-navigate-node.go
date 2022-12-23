package controller

type NavigateNode interface {
	Node
	URL() string
}

type chromedpNavigateNode struct {
	url string
}

func newNavigateNode(url string) *chromedpNavigateNode {
	return &chromedpNavigateNode{url}
}

func (navigateNode *chromedpNavigateNode) Accept(grouper ActionGrouper) error {
	return grouper.VisitNavigateNode(navigateNode)
}

func (navigateNode *chromedpNavigateNode) URL() string {
	return navigateNode.url
}

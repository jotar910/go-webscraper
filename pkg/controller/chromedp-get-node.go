package controller

type GetNode interface {
	Node
	Selector() string
}

type chromedpGetNode struct {
	selector string
}

func newGetNode(selector string) *chromedpGetNode {
	return &chromedpGetNode{selector}
}

func (GetNode *chromedpGetNode) Accept(grouper ActionGrouper) error {
	return grouper.VisitGetNode(GetNode)
}

func (GetNode *chromedpGetNode) Selector() string {
	return GetNode.selector
}

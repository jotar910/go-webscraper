package controller

type FindNode interface {
	Node
	Selector() string
}

type chromedpFindNode struct {
	selector string
}

func newFindNode(selector string) *chromedpFindNode {
	return &chromedpFindNode{selector}
}

func (findNode *chromedpFindNode) Accept(grouper ActionGrouper) error {
	return grouper.VisitFindNode(findNode)
}

func (findNode *chromedpFindNode) Selector() string {
	return findNode.selector
}

package controller

type ClickNode interface {
	Node
}

type chromedpClickNode struct {
}

func newClickNode() *chromedpClickNode {
	return &chromedpClickNode{}
}

func (clickNode *chromedpClickNode) Accept(grouper ActionGrouper) error {
	return grouper.VisitClickNode(clickNode)
}

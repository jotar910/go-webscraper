package controller

type TextAllNode interface {
	Node
	OutputValues() *[]string
}

type chromedpTextAllNode struct {
	output *[]string
}

func newTextAllNode(output *[]string) *chromedpTextAllNode {
	return &chromedpTextAllNode{output}
}

func (textAllNode *chromedpTextAllNode) Accept(grouper ActionGrouper) error {
	return grouper.VisitTextAllNode(textAllNode)
}

func (textAllNode *chromedpTextAllNode) OutputValues() *[]string {
	return textAllNode.output
}

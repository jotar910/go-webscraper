package controller

type TextNode interface {
	Node
	OutputValue() *string
}

type chromedpTextNode struct {
	output *string
}

func newTextNode(output *string) *chromedpTextNode {
	return &chromedpTextNode{output}
}

func (textNode *chromedpTextNode) Accept(grouper ActionGrouper) error {
	return grouper.VisitTextNode(textNode)
}

func (textNode *chromedpTextNode) OutputValue() *string {
	return textNode.output
}

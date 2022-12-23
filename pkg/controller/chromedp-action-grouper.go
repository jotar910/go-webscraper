package controller

import (
	"strings"

	"github.com/chromedp/chromedp"
)

type chromedpActionGrouper struct {
	selectorSb *strings.Builder
	actions    []chromedp.Action
}

func newChromedpActionGrouper() *chromedpActionGrouper {
	return &chromedpActionGrouper{
		selectorSb: new(strings.Builder),
		actions:    make([]chromedp.Action, 0),
	}
}

func (grouper *chromedpActionGrouper) Actions() []chromedp.Action {
	return grouper.actions
}

func (grouper *chromedpActionGrouper) VisitClickNode(node ClickNode) error {
	grouper.actions = append(grouper.actions, chromedp.Click(grouper.selectorSb.String(), chromedp.NodeVisible))
	return nil
}

func (grouper *chromedpActionGrouper) VisitFindNode(node FindNode) error {
	if grouper.selectorSb.Len() > 0 {
		grouper.selectorSb.WriteRune(' ')
	}
	grouper.selectorSb.WriteString(node.Selector())
	return nil
}

func (grouper *chromedpActionGrouper) VisitGetNode(node GetNode) error {
	grouper.selectorSb.Reset()
	grouper.selectorSb.WriteString(node.Selector())
	return nil
}

func (grouper *chromedpActionGrouper) VisitNavigateNode(node NavigateNode) error {
	grouper.selectorSb.Reset()
	grouper.actions = append(grouper.actions, chromedp.Navigate(node.URL()))
	return nil
}

func (grouper *chromedpActionGrouper) VisitTextAllNode(node TextAllNode) error {
	selector := grouper.selectorSb.String()
	grouper.actions = append(grouper.actions,
		chromedp.WaitVisible(selector),
		chromedp.Evaluate(`[...document.querySelectorAll('`+selector+`')].map((e) => e.innerText)`, node.OutputValues()),
	)
	return nil
}

func (grouper *chromedpActionGrouper) VisitTextNode(node TextNode) error {
	grouper.actions = append(grouper.actions, chromedp.Text(grouper.selectorSb.String(), node.OutputValue()))
	return nil
}

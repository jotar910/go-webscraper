package controller

type ActionGrouper interface {
	VisitClickNode(node ClickNode) error
	VisitFindNode(node FindNode) error
	VisitGetNode(node GetNode) error
	VisitNavigateNode(node NavigateNode) error
	VisitTextAllNode(node TextAllNode) error
	VisitTextNode(node TextNode) error
}

package controller

type Node interface {
	Accept(grouper ActionGrouper) error
}

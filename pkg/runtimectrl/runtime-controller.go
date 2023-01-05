package runtimectrl

type RuntimeController interface {
	Clone() RuntimeController
}

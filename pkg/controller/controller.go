package controller

type Controller interface {
	Click() Controller
	Find(selector string) Controller
	Get(selector string) Controller
	Navigate(url string) Controller
	Text(output *string) Controller
	TextAll(output *[]string) Controller

	Scrape() error
	Clone() Controller
}

func New() Controller {
	return newChromedpNodeController()
}

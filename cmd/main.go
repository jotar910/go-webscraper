package main

import (
	"fmt"
	"webscrapper/pkg/controller"
	"webscrapper/pkg/runtime"
)

func main() {
	// scraper.New().Scrape(func(doc *goquery.Document) {
	// a, e := doc.Find(".panel > div > div.row ~ div > div span.text-muted ~ h2 > div > span").Html()
	// a, e := doc.Find(".panel > div > div.row ~ div > div span.text-muted ~ h2 > div > span").Html()
	// doc.FindSelection(".panel")
	// fmt.Printf("%d\n%s\n%v\n", doc.Find(".panel > div > div.row").Size(), a, e)
	/* time.Sleep(5 * time.Second)
	a, e = doc.Find(".panel > div > div.row ~ div > div span.text-muted ~ h2 > div > span").Html()
	fmt.Printf("%d\n%s\n%v\n", doc.Size(), a, e) */
	/* time.Sleep(5 * time.Second)
		doc.Find("body").Each(func(i int, s *goquery.Selection) {
			a, e := s.Html()
			fmt.Printf("%v (%d) - %v\n", a, len(a), e)
			s.Children().Filter("div").Each(func(i int, s *goquery.Selection) {
				s.Children().Filter("div.row").Each(func(i int, s *goquery.Selection) {
					s.Siblings().Filter("div").Each(func(i int, s *goquery.Selection) {
						s.Children().Filter("div").Each(func(i int, s *goquery.Selection) {
							s.Find("span.text-muted").Each(func(i int, s *goquery.Selection) {
								s.Siblings().Filter("h2").Each(func(i int, s *goquery.Selection) {
									s.Children().Filter("div").Each(func(i int, s *goquery.Selection) {
										s.Children().Filter("span").Each(func(i int, s *goquery.Selection) {
											a, e := s.Html()
											fmt.Printf("%v (%d) - %v\n", strings.Contains(a, "Define Microservice Architecture"), len(a), e)
										})
									})
								})
							})
						})
					})
				})
			})
		})
	}, "https://www.fullstack.cafe/interview-questions/microservices") */
	// var example string
	/*
		var titles []string
		selector := `.panel > div > div.row ~ div > div span.text-muted ~ h2 > div > span`
		scraper.New().Scrape(
			chromedp.Navigate(`https://www.fullstack.cafe/interview-questions/microservices`),
			chromedp.WaitVisible(selector),
			chromedp.Evaluate(`[...document.querySelectorAll('`+selector+`')].map((e) => e.innerText)`, &titles),
			// chromedp.Query(`.panel > div > div.row ~ div > div span.text-muted ~ h2 > div > span`, chromedp.NodeVisible, chromedp.ByQueryAll),
			chromedp.QueryAfter(`.panel > div > div.row ~ div > div span.text-muted ~ h2 > div > span`, func(ctx context.Context, eci runtime.ExecutionContextID, n ...*cdp.Node) error {
				for _, node := range n {
					j := node.FullXPath()
					fmt.Printf("%+v\n\n", j)
				}
				fmt.Printf("Total: %d\n", len(n))
				return nil
			}, chromedp.NodeVisible, chromedp.ByQueryAll),
		)
		for _, title := range titles {
			fmt.Printf("%+v\n", title)
		}
		fmt.Printf("Total: %d\n", len(titles))
	*/
	/* var titles []string
	var title string
	controller.New().
		Navigate("https://www.fullstack.cafe/interview-questions/microservices").
		Get(".panel > div > div.row ~ div > div span.text-muted ~ h2 > div > span").
		Text(&title).
		TextAll(&titles).
		Scrape()
	for _, title := range titles {
		fmt.Printf("%+v\n", title)
	}
	fmt.Printf("Total: %d\n", len(titles))
	fmt.Printf("First: %s\n", title) */
	runtime.New(controller.New()).
		Run("Get questions from fullstack cafe", func(run runtime.Runtime[controller.Controller]) {
			run.Do("Questions regarding microservices", func(ctx *runtime.Context[controller.Controller]) {
				var titles []string
				var title string

				ctx.C.
					Navigate("https://www.fullstack.cafe/interview-questions/microservices").
					Get(".panel > div > div.row ~ div > div span.text-muted ~ h2 > div > span").
					Text(&title).
					TextAll(&titles).
					Scrape()

				for _, title := range titles {
					fmt.Printf("%+v\n", title)
				}
				fmt.Printf("Total: %d\n", len(titles))
				fmt.Printf("First: %s\n", title)
			})
		})
}

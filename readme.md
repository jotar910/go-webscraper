# Welcome to my Golang web scraper!👋

My web scraper is designed to make it easy for you to extract data from web pages. It is composed of three main components: the scraper, the controller, and the runtime.

## Scraper component

The scraper component uses the [chromedp](https://github.com/chromedp/chromedp) tool to perform the actual scraping of web pages. Chromedp is a Go library for driving Chrome or Chromium. It allows us to execute JavaScript and interact with Single Page Applications (SPAs) as a real user would, ensuring that accurate and up-to-date data can be extracted. Chromedp also provides a simple and efficient API, making it easy to use in my web scraper.

## Controller component

The controller component can be seen as a wrapper of the scraper component, abstracting some aspects of it and implementing a functional approach. It takes advantage of the visitor pattern to build the chain of actions that will be used in the scraper. This allows us to achieve flexibility and clearly specify the actions to be taken during the scraping phase.

## Runtime component

The runtime component is responsible for the execution lifecycle of the scraper. It utilizes a lifecycle approach, which includes the following hooks:

* "before" runs before everything. This can be used for tasks such as setting up a database connection or authenticating with a website.
* "after" runs after everything. This can be used for tasks such as closing a database connection or saving scraped data to a file.
* "beforeEach" runs before every iteration. This can be used for tasks such as resetting state or setting up the environment for the next iteration.
* "afterEach" runs after every iteration. This can be used for tasks such as cleaning up the environment or logging the results of the iteration.
* "do" represents the scraping code to run on every iteration. This is where the actual scraping takes place.

We are constantly working to improve my web scraper. Our next goal is to create a domain specific language (DSL) that will allow you to scrape pages using either a CLI or a web page interface. This DSL will provide a simple and intuitive way to specify the actions to be taken during the scraping process.

In addition to using my web scraper as a standalone tool, you can also use it as a library. By utilizing the plugin pattern, you can easily incorporate my web scraper into your own Go projects as a flexible and efficient data extraction solution.

Thank you for choosing my Golang web scraper. If you have any questions or suggestions, don't hesitate to reach out. We hope you find my scraper useful and efficient in your data extraction needs.
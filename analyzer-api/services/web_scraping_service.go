package services

import (
	"example/sample/dto"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

func WebScraper(baseURL string) *dto.WebPageDetail {

	var scrapedData = dto.WebPageDetail{}
	var wg = sync.WaitGroup{}

	urlStruct, _ := url.Parse(baseURL)
	host := strings.Split(baseURL, urlStruct.Path)[0]

	c := colly.NewCollector(
		colly.AllowedDomains(urlStruct.Hostname()),
	)

	func() {
		wg.Add(1)
		go c.OnHTML("head", func(e *colly.HTMLElement) {
			scrapedData.Title = e.ChildText("title")
			defer wg.Done()
		})
	}()

	go queryCollector("h1", c, &wg, &scrapedData)
	go queryCollector("h2", c, &wg, &scrapedData)
	go queryCollector("h3", c, &wg, &scrapedData)
	go queryCollector("h4", c, &wg, &scrapedData)
	go queryCollector("h5", c, &wg, &scrapedData)
	go queryCollector("h6", c, &wg, &scrapedData)
	go handleWebLinks(host, c, &wg, &scrapedData)
	go handleWebLogin(c, &wg, &scrapedData)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(baseURL)

	return &scrapedData

}

func handleWebLogin(c *colly.Collector,
	wg *sync.WaitGroup, scrapedData *dto.WebPageDetail) {
	wg.Add(1)
	defer wg.Done()
	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEachWithBreak("a", func(i int, link *colly.HTMLElement) bool {
			text := strings.TrimSpace(strings.ToLower(link.Text))

			if text == "log in" || text == "sign in" || text == "login" {
				scrapedData.IsWithLogin = true
				return false
			}
			return true
		})

		if scrapedData.IsWithLogin {
			return
		}

		e.ForEachWithBreak("button", func(i int, link *colly.HTMLElement) bool {
			text := strings.TrimSpace(strings.ToLower(link.Text))
			if text == "log in" || text == "sign in" {
				scrapedData.IsWithLogin = true
				return false
			}
			return true
		})
	})
}

func handleWebLinks(host string, c *colly.Collector,
	wg *sync.WaitGroup, scrapedData *dto.WebPageDetail) {

	wg.Add(1)
	defer wg.Done()
	c.OnHTML("body", func(e *colly.HTMLElement) {

		for _, v := range e.ChildAttrs("a", "href") {

			if !strings.Contains(v, "://") && !strings.HasPrefix(v, "//") {

				scrapedData.InternalLink++

				if strings.HasPrefix(v, "#") {
					if e.ChildAttr(v, "id") == "" {
						scrapedData.InternalDeadIdLink++
						continue
					}
					continue
				}

				go validateWebLink(host+v, "internalPath", wg, scrapedData)
				continue
			}

			scrapedData.ExternalLink++

			go validateWebLink(v, "external", wg, scrapedData)
		}

	})

}

func queryCollector(goquerySelector string, c *colly.Collector,
	wg *sync.WaitGroup, scrapedData *dto.WebPageDetail) {

	wg.Add(1)
	defer wg.Done()

	c.OnHTML(goquerySelector, func(e *colly.HTMLElement) {

		switch goquerySelector {
		case "h1":
			scrapedData.H1++
		case "h2":
			scrapedData.H2++
		case "h3":
			scrapedData.H3++
		case "h4":
			scrapedData.H4++
		case "h5":
			scrapedData.H5++
		case "h6":
			scrapedData.H6++

		}
	})

}

// some working external links in the app get failed with http.GET method
// witch contains unsupported protocol scheme eg: {//, ://}, those links not handled
func validateWebLink(url string, linkType string, wg *sync.WaitGroup, scrapedData *dto.WebPageDetail) {
	wg.Add(1)
	defer wg.Done()

	if status, _ := ValidateURL(url); status == 200 {
		return
	}

	switch linkType {
	case "internalPath":
		scrapedData.InternalDeadPathLink++
	case "external":
		scrapedData.ExternalDeadLink++
	}
}

func ValidateURL(_url string) (int, string) {

	validatedURL, err := url.ParseRequestURI(_url)

	if err != nil {
		return 400, err.Error()
	}

	_res, err := http.Get(validatedURL.String())

	if err != nil {
		return 400, err.Error()
	}

	return _res.StatusCode, _res.Status

}

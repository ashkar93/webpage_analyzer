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

	//baseURL = "https://en.wikipedia.org/wiki/Web_scraping"
	urlStruct, _ := url.Parse(baseURL)
	host := strings.Split(baseURL, urlStruct.Path)[0]

	c := colly.NewCollector(
		colly.AllowedDomains(urlStruct.Hostname()),
	)

	func() {
		wg.Add(1)
		go c.OnHTML("head", func(e *colly.HTMLElement) {
			scrapedData.Title = e.ChildText("title")
		})
		defer wg.Done()
	}()

	func() {
		wg.Add(1)
		go c.OnHTML("h1", func(e *colly.HTMLElement) {
			scrapedData.H1++
		})
		defer wg.Done()
	}()

	func() {
		wg.Add(1)
		go c.OnHTML("h2", func(e *colly.HTMLElement) {
			scrapedData.H2++
		})
		defer wg.Done()
	}()

	func() {
		wg.Add(1)
		go c.OnHTML("h3", func(e *colly.HTMLElement) {
			scrapedData.H3++
		})
		defer wg.Done()
	}()

	func() {
		wg.Add(1)
		go c.OnHTML("h4", func(e *colly.HTMLElement) {
			scrapedData.H4++
		})
		defer wg.Done()
	}()

	func() {
		wg.Add(1)
		go c.OnHTML("h5", func(e *colly.HTMLElement) {
			scrapedData.H5++
		})
		defer wg.Done()
	}()

	func() {
		wg.Add(1)
		go c.OnHTML("h6", func(e *colly.HTMLElement) {
			scrapedData.H6++
		})
		defer wg.Done()
	}()

	func() {
		wg.Add(1)
		go c.OnHTML("body", func(e *colly.HTMLElement) {
			for _, v := range e.ChildAttrs("a", "href") {

				if !strings.Contains(v, "://") && !strings.HasPrefix(v, "//") {

					scrapedData.InternalLink++

					if strings.HasPrefix(v, "#") && e.ChildAttr(v, "id") == "" {
						scrapedData.InternalDeadIdLink++
						continue
					}

					wg.Add(1)
					go validateLink(host+v, "internalPath", &wg, &scrapedData)
					continue
				}

				scrapedData.ExternalLink++
				wg.Add(1)
				go validateLink(v, "external", &wg, &scrapedData)
			}
		})
		defer wg.Done()
	}()

	func() {
		wg.Add(1)
		go c.OnHTML("body", func(e *colly.HTMLElement) {
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
		defer wg.Done()
	}()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(baseURL)

	return &scrapedData

}

// some working external links in the app get failed with http.GET method
// witch contains unsupported protocol scheme eg: {//, ://}, those links not handled
func validateLink(url string, linkType string, wg *sync.WaitGroup, scrapedData *dto.WebPageDetail) {

	defer wg.Done()

	if _, err := http.Get(url); err == nil {
		return
	}

	switch linkType {
	case "internalPath":
		scrapedData.InternalDeadPathLink++
	case "external":
		scrapedData.ExternalDeadLink++
	}
}

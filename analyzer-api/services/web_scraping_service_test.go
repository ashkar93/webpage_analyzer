package services

import (
	"example/sample/dto"
	"testing"
)

func TestWebScraper(t *testing.T) {

	got := WebScraper("https://go101.org/article/101.html")

	want := &dto.WebPageDetail{Version: "", Title: "Go 101 -Go 101", InternalLink: 67,
		ExternalLink: 18, InternalDeadIdLink: 0, InternalDeadPathLink: 0,
		ExternalDeadLink: 0, H1: 1, H2: 0, H3: 1, H4: 0, H5: 0, H6: 0}

	if *got != *want {
		t.Errorf("got %+v\n, wanted %+v\n", got, want)
	}

	got_2 := WebScraper("https://en.wikipedia.org/wiki/Web_scraping")

	want_2 := &dto.WebPageDetail{Version: "", Title: "Web scraping - Wikipedia", InternalLink: 284,
		ExternalLink: 84, InternalDeadIdLink: 0, InternalDeadPathLink: 0,
		ExternalDeadLink: 5, H1: 1, H2: 9, H3: 21, H4: 0, H5: 0, H6: 0, IsWithLogin: true}

	if *got_2 != *want_2 {
		t.Errorf("got %+v\n, wanted %+v\n", got_2, want_2)
	}

}

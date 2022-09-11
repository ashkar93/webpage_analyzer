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

}

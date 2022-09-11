package controllers

import (
	"example/sample/common"
	"example/sample/services"
	"net/http"
	"strings"
)

func WebScraper(res http.ResponseWriter, req *http.Request) {

	url := strings.TrimSpace(req.URL.Query().Get("url"))

	if url == "" {
		common.RespondWithError(res, 400, "Not a valid URL")
		return
	}

	_res, err := http.Get(url)

	if err != nil {
		common.RespondWithError(res, 400, err.Error())
		return
	}

	if _res.StatusCode != 200 {
		common.RespondWithError(res, _res.StatusCode, _res.Status)
		return
	}

	r := services.WebScraper(url)

	common.RespondwithJSON(res, 200, *r)

}

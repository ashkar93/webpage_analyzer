package controllers

import (
	"example/sample/common"
	"example/sample/services"
	"net/http"
)

func WebScraper(res http.ResponseWriter, req *http.Request) {

	status, msg := services.ValidateURL(req.URL.Query().Get("url"))

	if status != 200 {
		common.RespondWithError(res, status, msg)
		return
	}

	r := services.WebScraper(req.URL.Query().Get("url"))

	common.RespondwithJSON(res, 200, *r)

}

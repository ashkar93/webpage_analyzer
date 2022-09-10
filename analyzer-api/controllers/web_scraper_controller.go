package controllers

import (
	"example/sample/common"
	"example/sample/services"
	"fmt"
	"net/http"
)

func WebScraper(res http.ResponseWriter, req *http.Request) {

	r := services.WebScraper(req.URL.Query().Get("url"))
	fmt.Printf("%+v\n", *r)
	common.RespondwithJSON(res, 200, *r)

}

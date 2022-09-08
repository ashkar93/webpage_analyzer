package controllers

import (
	"example/sample/common"
	"net/http"
)

func GetAnalyze(res http.ResponseWriter, req *http.Request) {

	common.RespondwithJSON(res, 200, req.URL.Query().Get("url"))

}

package main

import (
	"example/sample/routers"
	"net/http"

	"github.com/sirupsen/logrus"
)

// @title Webpage Analyzer
// @version 1.0
// @description Webpage Analyzer contains APIs for analyze a web page
func main() {

	routes := routers.InitRouter()

	logrus.Info("app started on :8000")
	http.ListenAndServe(":8000", routes)
}

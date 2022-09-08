package controllers

import (
	"example/sample/common"
	"net/http"
)

func SayHelloworld(res http.ResponseWriter, req *http.Request) {

	common.RespondwithJSON(res, 200, "Hello world!")

}

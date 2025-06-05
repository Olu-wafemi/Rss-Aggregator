package main

import (
	"net/http"
)

func handlerError(res http.ResponseWriter, req *http.Request) {
	respondwithError(res, 400, "Something went Wrong")
}

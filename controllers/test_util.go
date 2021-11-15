package controllers

import (
	"net/http"
	"net/http/httptest"
	"github.com/gorilla/mux"
)

func SendRequest(route, url, verb string, function func(http.ResponseWriter, *http.Request)) (*httptest.ResponseRecorder){
	req, _ := http.NewRequest(verb, url, nil)
	res := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(route, function)
	router.ServeHTTP(res, req)
	return res
}
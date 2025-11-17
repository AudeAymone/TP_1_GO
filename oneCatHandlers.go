package main

import "net/http"

func getCat(req *http.Request) (int, any) {
	Logger.Info("Cat not found")
	return http.StatusNotFound, "Cat not found"
}

package main

import (
	"net/http"
)


func getCat(req *http.Request) (int, any) {
	catID := req.PathValue("catId")
	cat, exists := catsDatabase[catID]

	if !exists {
		return http.StatusNotFound, "Cat not found"
	}

	return http.StatusOK, cat
}


func deleteCat(req *http.Request) (int, any) {
	catID := req.PathValue("catId")
	_, exists := catsDatabase[catID]

	if !exists {
		return http.StatusNotFound, "Cat not found"
	}

	delete(catsDatabase, catID)
	return http.StatusNoContent, nil
}
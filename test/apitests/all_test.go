package apitests

import (
	"fmt"
	"net/http"
	"testing"
)

var initCatId string

func init() {
	// Preparation: delete all existing & create a cat
	ids := []string{}
	call("GET", "/cats", nil, nil, &ids)

	for _, id := range ids {
		code := 0
		call("DELETE", "/cats/"+id, nil, &code, nil)
		fmt.Println("DELETE /cats ->", code)
	}

	// Create a single cat into the DB
	call("POST", "/cats", &CatModel{Name: "Toto"}, nil, &initCatId)
}

func TestGetCats(t *testing.T) {

	code := 0
	result := []string{}
	err := call("GET", "/cats", nil, &code, &result)
	if err != nil {
		t.Error("Request error", err)
	}

	fmt.Println("GET /cats ->", code, result)

	if code != http.StatusOK {
		t.Error("We should get code 200, got", code)
	}

	if len(result) != 1 {
		t.Error("We should get one item, got", len(result))
		return
	}

	if result[0] != initCatId {
		t.Error("Listing the IDs, got", result[0])
	}
}

func TestCreateCat(t *testing.T) {
	code := 0
	newID := ""

	err := call("POST", "/cats", &CatModel{Name: "Mimi"}, &code, &newID)
	if err != nil {
		t.Error("Request error", err)
	}

	fmt.Println("POST /cats ->", code, newID)

	if code != http.StatusCreated {
		t.Error("We should get code 201, got", code)
	}

	if newID == "" {
		t.Error("New cat ID should not be empty")
	}
}

func TestGetCatById(t *testing.T) {
	code := 0
	result := CatModel{}

	err := call("GET", "/cats/"+initCatId, nil, &code, &result)
	if err != nil {
		t.Error("Request error", err)
	}

	if code != http.StatusOK {
		t.Error("Expected 200, got", code)
	}

	if result.ID != initCatId {
		t.Error("Expected ID", initCatId, "got", result.ID)
	}
}



func TestDeleteCat(t *testing.T) {
	code := 0
	newID := ""

	// Créer un chat à supprimer
	call("POST", "/cats", &CatModel{Name: "ToDelete"}, &code, &newID)

	// Supprimer
	code = 0
	err := call("DELETE", "/cats/"+newID, nil, &code, nil)
	if err != nil {
		t.Error("Request error", err)
	}

	if code != http.StatusNoContent {
		t.Error("Expected 204, got", code)
	}
}



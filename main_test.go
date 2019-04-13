package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetAOCHandlerHappyPath(t *testing.T) {
	req, err := http.NewRequest("GET", "/aoc/0142e2fa3543cb32bf000100620002", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := executeRequest(req, t)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("%s returned wrong status code: got '%v' want '%v'",
			req.URL.Path, status, http.StatusOK)
	}

	body, _ := ioutil.ReadAll(rr.Body)
	data := sip{}

	json.Unmarshal(body, &data)

	if data.AddressOfRecord != "0142e2fa3543cb32bf000100620002" {
		t.Errorf("%s returned the wrong 'addressOfRecord': got '%s' want '%s'",
			req.URL.Path, data.AddressOfRecord, "0142e2fa3543cb32bf000100620002")
	}
}

func TestGetAOCHandlerUnhappyPath(t *testing.T) {
	req, err := http.NewRequest("GET", "/aoc/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := executeRequest(req, t)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("%s returned wrong status code: got %v want %v",
			req.URL.Path, status, http.StatusOK)
	}

	body, _ := ioutil.ReadAll(rr.Body)

	if !strings.Contains(string(body), "doesn't seem to be valid") {
		t.Errorf("%s didn't return the correct error JSON.", req.URL.Path)
	}
}

func executeRequest(req *http.Request, t *testing.T) *httptest.ResponseRecorder {
	rh := registrationHandler{
		registrations: loadRegs("./data/regs"),
	}

	res := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/aoc/{address}", rh.getSIPHandler)
	r.ServeHTTP(res, req)

	return res
}

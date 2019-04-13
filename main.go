package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type registrationHandler struct {
	registrations regs
}

func main() {
	port := ":80"
	r := registrationHandler{
		registrations: loadRegs("./data/regs"),
	}
	router := mux.NewRouter()

	router.HandleFunc("/aor/{address}", r.getSIPHandler).Methods("GET")

	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":80", router))
}

func (rh registrationHandler) getSIPHandler(w http.ResponseWriter, r *http.Request) {
	found := false
	params := mux.Vars(r)
	for i := range rh.registrations {
		if rh.registrations[i].AddressOfRecord == params["address"] {
			found = true
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(rh.registrations[i])
		}
	}
	if !found {
		notFoundHandler(w, r, params)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request, p map[string]string) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, `{"error": "AOR '%s' doesn't seem to be valid."}`, p["address"])
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type sip struct {
	AddressOfRecord string   `json:"addressOfRecord,omitempty"`
	TenantID        string   `json:"tenantId,omitempty"`
	URI             string   `json:"uri,omitempty"`
	Contact         string   `json:"contact,omitempty"`
	Path            []string `json:"path,omitempty"`
	Source          string   `json:"source,omitempty"`
	Target          string   `json:"target,omitempty"`
	UserAgent       string   `json:"userAgent,omitempty"`
	RawUserAgent    string   `json:"rawUserAgent,omitempty"`
	Created         string   `json:"created,omitempty"`
	LineID          string   `json:"lineId,omitempty"`
}

type regs []sip

func loadRegs(filename string) regs {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	r := regs{}
	s := strings.Split(string(bs), "\n")

	for _, v := range s {
		if v != "" {
			data := sip{}
			json.Unmarshal([]byte(v), &data)
			r = append(r, data)
		}
	}

	return r
}

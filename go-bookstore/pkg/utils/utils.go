package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/**
ParseBody extracts and parses the contents of a HTTP request called `r` and assigns it to `x`
**/
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

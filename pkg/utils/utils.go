package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ParseBody parses the body of a request into a struct
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

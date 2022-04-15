package testutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func AssertMethodDidNotError(err error, response *http.Response, methodName string, t *testing.T) {
	if err != nil {
		t.Errorf("%s response: %v", methodName, response)
		t.Fatalf("%s failed: %v", methodName, err)
	}
}

func AssertPropertiesMatch(actualProperty interface{}, expectedProperty interface{}, t *testing.T, comparedType ...interface{}) {
	if comparedType != nil {
		if !cmp.Equal(actualProperty, expectedProperty, cmp.AllowUnexported(comparedType...)) {
			t.Errorf("does not match expected: got: %+v, want: %+v", spew.Sdump(actualProperty), spew.Sdump(expectedProperty))
		}
	} else {
		if !cmp.Equal(actualProperty, expectedProperty) {
			t.Errorf("does not match expected: got: %+v, want: %+v", spew.Sdump(actualProperty), spew.Sdump(expectedProperty))
		}
	}
}

func AssertPropertiesDoNotMatch(actualProperty interface{}, expectedProperty interface{}, t *testing.T, comparedType ...interface{}) {
	if comparedType != nil {
		if cmp.Equal(actualProperty, expectedProperty, cmp.AllowUnexported(comparedType...)) {
			t.Errorf("does not match expected: got: %+v, want: %+v", spew.Sdump(actualProperty), spew.Sdump(expectedProperty))
		}
	} else {
		if cmp.Equal(actualProperty, expectedProperty) {
			t.Errorf("does not match expected: got: %+v, want: %+v", spew.Sdump(actualProperty), spew.Sdump(expectedProperty))
		}
	}
}

func AssertNotFound(err error, t *testing.T) {
	assert.EqualErrorf(t, err, "404 Not Found", "error should be: %v, got: %v", "404 Not Found", err)
}

func AssertRequestWasMade(httpMethod string, path string, pathParameters []interface{}) []interface{} {
	client := &http.Client{}

	body := make(map[string]interface{})

	body["method"] = httpMethod
	body["path"] = path

	if pathParameters != nil {
		body["pathParameters"] = pathParameters
	}

	jsonBytes, err := json.Marshal(body)

	req, _ := http.NewRequest("PUT", "http://localhost:1080/mockserver/retrieve", bytes.NewReader(jsonBytes))

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("error retrieving requests")
	}

	var matchingRequests []interface{}
	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&matchingRequests); err != nil {
		log.Fatalf("error deserializing response")
	}

	fmt.Printf("matching request count: %v", len(matchingRequests))

	return matchingRequests
}

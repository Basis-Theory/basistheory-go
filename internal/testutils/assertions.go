package testutils

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
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

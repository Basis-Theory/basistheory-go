package testutils

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"testing"
)

func AssertMethodDidNotError(err error, response *http.Response, methodName string, t *testing.T) {
	if err != nil {
		t.Errorf("%s response: %v", methodName, response)
		t.Fatalf("%s failed: %v", methodName, err)
	}
}

func AssertPropertiesMatch(actualProperty interface{}, expectedProperty interface{}, t *testing.T) {
	if !cmp.Equal(actualProperty, expectedProperty) {
		t.Errorf("does not match expected: got: %+v, want: %+v", spew.Sdump(actualProperty), spew.Sdump(expectedProperty))
	}
}

func AssertPropertiesMatchGeneric(actualProperty interface{}, expectedProperty interface{}, t *testing.T, comparedType ...interface{}) {
	if !cmp.Equal(actualProperty, expectedProperty, cmp.AllowUnexported(comparedType...)) {
		t.Errorf("does not match expected: got: %+v, want: %+v", spew.Sdump(actualProperty), spew.Sdump(expectedProperty))
	}
}

func AssertPropertiesDoNotMatchGeneric(actualProperty interface{}, expectedProperty interface{}, t *testing.T, comparedType ...interface{}) {
	if cmp.Equal(actualProperty, expectedProperty, cmp.AllowUnexported(comparedType...)) {
		t.Errorf("matches expected: got: %+v, want: %+v", spew.Sdump(actualProperty), spew.Sdump(expectedProperty))
	}
}

func AssertDeletion[T any](actualDeletedResource T, expectedDeletedResource T, methodName string, t *testing.T) {
	if !cmp.Equal(actualDeletedResource, expectedDeletedResource) {
		t.Fatalf("%s failed to delete", methodName)
	}
}

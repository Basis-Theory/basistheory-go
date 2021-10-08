package test_utils

import (
	"github.com/jarcoal/httpmock"
	"github.com/jaswdr/faker"
)

func SetupErrorPath(httpMethod string, path string) (interface{}, int) {
	fakerInst := faker.New()
	const lowestHttpStatusError, highestHttpStatusError = 400, 500
	status := fakerInst.IntBetween(lowestHttpStatusError, highestHttpStatusError)
	errorPayload := map[string]string{
		fakerInst.Lorem().Word(): fakerInst.Lorem().Word(),
	}
	responder, _ := httpmock.NewJsonResponder(status, errorPayload)
	httpmock.RegisterResponder(httpMethod, path, responder)

	return errorPayload, status
}

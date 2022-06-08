package basistheory_test

import (
	"github.com/Basis-Theory/basistheory-go/v2"
	"github.com/Basis-Theory/basistheory-go/v2/internal/testutils"
	"testing"
)

func TestTokenizeCRUD(t *testing.T) {
	// SETUP
	apiClient, contextWithAPIKey := testutils.CreateApiAndSrvrContext(t)

	// CREATE
	tokenData := map[string]interface{}{
		"myData":      "myValue",
		"myOtherData": "myOtherValue",
	}
	createdToken, response, err := apiClient.TokenizeApi.TokenizeTokenize(contextWithAPIKey).Body(tokenData).Execute()

	testutils.AssertMethodDidNotError(err, response, "TokenizeTokenize", t)

	// GET List
	myDataTokenId := createdToken.(map[string]interface{})["myData"].(string)
	myOtherDataTokenId := createdToken.(map[string]interface{})["myOtherData"].(string)

	var tokens *basistheory.TokenPaginatedList
	tokens, response, err = apiClient.TokensApi.TokensGet(contextWithAPIKey).Id([]string{myDataTokenId, myOtherDataTokenId}).Execute()

	testutils.AssertMethodDidNotError(err, response, "TokensGetById", t)
	testutils.AssertPropertiesMatch(tokens.Data[0].GetData(), "myOtherValue", t)
	testutils.AssertPropertiesMatch(tokens.Data[1].GetData(), "myValue", t)

	// DELETE
	response, err = apiClient.TokensApi.TokensDelete(contextWithAPIKey, myDataTokenId).Execute()

	testutils.AssertMethodDidNotError(err, response, "TokensDelete", t)

	response, err = apiClient.TokensApi.TokensDelete(contextWithAPIKey, myOtherDataTokenId).Execute()

	testutils.AssertMethodDidNotError(err, response, "TokensDelete", t)
}

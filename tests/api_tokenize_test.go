package basistheory_test

import (
	"github.com/Basis-Theory/basistheory-go/v3"
	"github.com/Basis-Theory/basistheory-go/v3/internal/testutils"
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
	createdToken, response, err := apiClient.TokenizeApi.Tokenize(contextWithAPIKey).Body(tokenData).Execute()

	testutils.AssertMethodDidNotError(err, response, "TokenizeTokenize", t)

	// GET List
	myDataTokenId := createdToken.(map[string]interface{})["myData"].(string)
	myOtherDataTokenId := createdToken.(map[string]interface{})["myOtherData"].(string)

	var tokens *basistheory.TokenPaginatedList
	tokens, response, err = apiClient.TokensApi.Get(contextWithAPIKey).Id([]string{myDataTokenId, myOtherDataTokenId}).Execute()

	var myDataToken basistheory.Token
	var myOtherDataToken basistheory.Token
	for _, tokenData := range tokens.Data {
		if tokenData.GetId() == myDataTokenId {
			myDataToken = tokenData
		} else {
			myOtherDataToken = tokenData
		}
	}

	testutils.AssertMethodDidNotError(err, response, "TokensGetById", t)
	testutils.AssertPropertiesMatch(myDataToken.GetData(), "myValue", t)
	testutils.AssertPropertiesMatch(myOtherDataToken.GetData(), "myOtherValue", t)

	// DELETE
	response, err = apiClient.TokensApi.Delete(contextWithAPIKey, myDataTokenId).Execute()

	testutils.AssertMethodDidNotError(err, response, "TokensDelete", t)

	response, err = apiClient.TokensApi.Delete(contextWithAPIKey, myOtherDataTokenId).Execute()

	testutils.AssertMethodDidNotError(err, response, "TokensDelete", t)
}

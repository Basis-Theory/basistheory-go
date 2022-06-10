package basistheory_test

import (
	"github.com/Basis-Theory/basistheory-go/v3"
	"github.com/Basis-Theory/basistheory-go/v3/internal/testutils"
	"testing"
)

func TestTokenCRUD(t *testing.T) {
	// SETUP
	apiClient, contextWithAPIKey := testutils.CreateApiAndSrvrContext(t)

	// CREATE
	tokenData := map[string]interface{}{
		"myData": "myValue",
	}
	tokenType := "token"
	tokenSearchIndexes := []string{"{{ data.myData }}"}
	createTokenRequest := *basistheory.NewCreateTokenRequest(tokenData)
	createTokenRequest.SetType(tokenType)
	createTokenRequest.SetSearchIndexes(tokenSearchIndexes)

	createdToken, response, err := apiClient.TokensApi.Create(contextWithAPIKey).CreateTokenRequest(createTokenRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "TokensCreate", t)
	testutils.AssertPropertiesMatch(createdToken.GetData(), nil, t)
	testutils.AssertPropertiesMatch(createdToken.GetType(), tokenType, t)
	testutils.AssertPropertiesMatch(createdToken.GetSearchIndexes(), tokenSearchIndexes, t)

	// GET BY ID
	var token *basistheory.Token
	token, response, err = apiClient.TokensApi.GetById(contextWithAPIKey, createdToken.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "TokensGetById", t)
	testutils.AssertPropertiesMatch(token.GetData(), tokenData, t)
	testutils.AssertPropertiesMatch(token.GetType(), tokenType, t)
	testutils.AssertPropertiesMatch(token.GetSearchIndexes(), tokenSearchIndexes, t)

	// GET LIST
	var tokens *basistheory.TokenPaginatedList
	tokens, response, err = apiClient.TokensApi.Get(contextWithAPIKey).Execute()

	testutils.AssertMethodDidNotError(err, response, "TokensGet", t)
	testutils.AssertPropertiesMatch(tokens.Data[0].GetData(), tokenData, t)
	testutils.AssertPropertiesMatch(tokens.Data[0].GetType(), tokenType, t)
	testutils.AssertPropertiesMatch(tokens.Data[0].GetSearchIndexes(), tokenSearchIndexes, t)

	// DELETE
	response, err = apiClient.TokensApi.Delete(contextWithAPIKey, createdToken.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "TokensDelete", t)

	_, _, err = apiClient.TokensApi.GetById(contextWithAPIKey, createdToken.GetId()).Execute()

	testutils.AssertNotFound(err, t)
}

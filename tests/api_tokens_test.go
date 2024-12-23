package basistheory_test

import (
	"fmt"
	"testing"

	"github.com/Basis-Theory/basistheory-go/v6"
	"github.com/Basis-Theory/basistheory-go/v6/internal/testutils"
)

func TestTokenCRUD(t *testing.T) {
	// SETUP
	apiClient, contextWithAPIKey := testutils.CreateApiAndSrvrContext(t)

	// CREATE
	tokenData := map[string]interface{}{
		"myData": "4444-4444-4444-4444",
	}
	tokenMask := map[string]interface{}{
		"myData": "{{ data.myData | append: 'suffix' }}",
	}
	tokenType := "token"
	tokenSearchIndexes := []string{"{{ data.myData }}"}
	createTokenRequest := *basistheory.NewCreateTokenRequest()
	createTokenRequest.SetData(tokenData)
	createTokenRequest.SetId("{{ data.myData | alias_preserve_length: 0, 4 }}")
	createTokenRequest.SetType(tokenType)
	createTokenRequest.SetSearchIndexes(tokenSearchIndexes)
	createTokenRequest.SetDeduplicateToken(false)
	createTokenRequest.SetMask(tokenMask)

	createdToken, response, err := apiClient.TokensApi.Create(contextWithAPIKey).CreateTokenRequest(createTokenRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "TokensApi Create", t)
	testutils.AssertPropertiesDoNotMatch(createdToken.GetId(), tokenData["myData"], t)
	testutils.AssertPropertiesMatch(createdToken.GetId()[len(createdToken.GetId())-4:], "4444", t)
	testutils.AssertPropertiesMatch(createdToken.GetData().(map[string]interface{})["myData"], fmt.Sprint(tokenData["myData"], "suffix"), t)
	testutils.AssertPropertiesMatch(createdToken.GetType(), tokenType, t)
	testutils.AssertPropertiesMatch(createdToken.GetSearchIndexes(), tokenSearchIndexes, t)
	testutils.AssertPropertiesMatch(createdToken.GetMask(), tokenMask, t)

	// GET BY ID
	var token *basistheory.Token
	token, response, err = apiClient.TokensApi.GetById(contextWithAPIKey, createdToken.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "TokensApi GetById", t)
	testutils.AssertPropertiesMatch(token.GetData().(map[string]interface{})["myData"], tokenData["myData"], t)
	testutils.AssertPropertiesMatch(token.GetType(), tokenType, t)
	testutils.AssertPropertiesMatch(token.GetSearchIndexes(), tokenSearchIndexes, t)
	testutils.AssertPropertiesMatch(token.GetMask(), tokenMask, t)

	// GET LIST
	var tokens *basistheory.TokenPaginatedList
	tokens, response, err = apiClient.TokensApi.Get(contextWithAPIKey).Execute()

	testutils.AssertMethodDidNotError(err, response, "TokensApi Get", t)
	testutils.AssertPropertiesMatch(tokens.Data[0].GetData().(map[string]interface{})["myData"], tokenData["myData"], t)
	testutils.AssertPropertiesMatch(tokens.Data[0].GetType(), tokenType, t)
	testutils.AssertPropertiesMatch(tokens.Data[0].GetMask(), tokenMask, t)

	// PATCH
	updatedTokenData := map[string]interface{}{
		"myNewData": "myNewValue",
	}
	updatedTokenMask := map[string]interface{}{
		"myNewData": "{{ data.myNewData | prepend: 'prefix' }}",
	}
	updatedTokenSearchIndexes := []string{"{{ data.myNewData }}"}
	updateTokenRequest := *basistheory.NewUpdateTokenRequest()
	updateTokenRequest.SetData(updatedTokenData)
	updateTokenRequest.SetMask(updatedTokenMask)
	updateTokenRequest.SetSearchIndexes(updatedTokenSearchIndexes)
	updateTokenRequest.SetDeduplicateToken(false)

	token, response, err = apiClient.TokensApi.Update(contextWithAPIKey, createdToken.GetId()).UpdateTokenRequest(updateTokenRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "TokensApi Update", t)
	testutils.AssertPropertiesMatch(token.GetData().(map[string]interface{})["myData"], fmt.Sprint(tokenData["myData"], "suffix"), t)
	testutils.AssertPropertiesMatch(token.GetData().(map[string]interface{})["myNewData"], fmt.Sprint("prefix", updatedTokenData["myNewData"]), t)
	testutils.AssertPropertiesMatch(token.GetType(), tokenType, t)
	testutils.AssertPropertiesMatch(token.GetSearchIndexes(), updatedTokenSearchIndexes, t)
	testutils.AssertPropertiesMatch(token.GetMask().(map[string]interface{})["myData"], tokenMask["myData"], t)
	testutils.AssertPropertiesMatch(token.GetMask().(map[string]interface{})["myNewData"], updatedTokenMask["myNewData"], t)

	// DELETE
	response, err = apiClient.TokensApi.Delete(contextWithAPIKey, createdToken.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "TokensApi Delete", t)

	_, _, err = apiClient.TokensApi.GetById(contextWithAPIKey, createdToken.GetId()).Execute()

	testutils.AssertNotFound(err, t)
}

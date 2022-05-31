package basistheory_test

import (
	"github.com/Basis-Theory/basistheory-go"
	"github.com/Basis-Theory/basistheory-go/internal/testutils"
	"testing"
)

func TestReactorFormulaCRUD(t *testing.T) {
	// SETUP
	apiClient, contextWithAPIKey := testutils.CreateApiAndMgmtContext(t)

	// CREATE
	reactorFormulaName := "Go Test Reactor Formula"
	reactorFormulaCode := "module.exports = function (req) {return {raw: \"Goodbye World\"}}"

	createReactorFormulaRequest := *basistheory.NewCreateReactorFormulaRequest("private", reactorFormulaName)
	createReactorFormulaRequest.SetCode(reactorFormulaCode)

	createdReactorFormula, response, err := apiClient.ReactorFormulasApi.ReactorFormulasCreate(contextWithAPIKey).CreateReactorFormulaRequest(createReactorFormulaRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorFormulaCreate", t)
	testutils.AssertPropertiesMatch(createdReactorFormula.GetName(), reactorFormulaName, t)
	testutils.AssertPropertiesMatch(createdReactorFormula.GetCode(), reactorFormulaCode, t)

	// GET BY ID
	var reactorFormula *basistheory.ReactorFormula
	reactorFormula, response, err = apiClient.ReactorFormulasApi.ReactorFormulasGetById(contextWithAPIKey, createdReactorFormula.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorFormulaGetById", t)
	testutils.AssertPropertiesMatch(reactorFormula.GetName(), reactorFormulaName, t)
	testutils.AssertPropertiesMatch(reactorFormula.GetCode(), reactorFormulaCode, t)

	// GET LIST
	var reactorFormulas *basistheory.ReactorFormulaPaginatedList
	reactorFormulas, response, err = apiClient.ReactorFormulasApi.ReactorFormulasGet(contextWithAPIKey).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorFormulasGet", t)
	testutils.AssertPropertiesMatch(reactorFormulas.Data[0].GetName(), reactorFormulaName, t)
	testutils.AssertPropertiesMatch(reactorFormulas.Data[0].GetCode(), reactorFormulaCode, t)

	// UPDATE
	updatedReactorFormulaName := "Go Test Reactor Formula Update"
	updatedReactorFormulaCode := "module.exports = function (req) {return {raw: \"Hello World\"}}"
	updateReactorFormulaRequest := *basistheory.NewUpdateReactorFormulaRequest("private", updatedReactorFormulaName)
	updateReactorFormulaRequest.SetCode(updatedReactorFormulaCode)

	var updatedReactorFormula *basistheory.ReactorFormula
	updatedReactorFormula, response, err = apiClient.ReactorFormulasApi.ReactorFormulasUpdate(contextWithAPIKey, createdReactorFormula.GetId()).UpdateReactorFormulaRequest(updateReactorFormulaRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorFormulaUpdate", t)
	testutils.AssertPropertiesMatch(updatedReactorFormula.GetName(), updatedReactorFormulaName, t)
	testutils.AssertPropertiesMatch(updatedReactorFormula.GetCode(), updatedReactorFormulaCode, t)

	// DELETE
	_, err = apiClient.ReactorFormulasApi.ReactorFormulasDelete(contextWithAPIKey, createdReactorFormula.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorFormulaDelete", t)

	_, _, err = apiClient.ReactorFormulasApi.ReactorFormulasGetById(contextWithAPIKey, createdReactorFormula.GetId()).Execute()

	testutils.AssertNotFound(err, t)
}

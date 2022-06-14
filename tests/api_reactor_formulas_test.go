package basistheory_test

import (
	"github.com/Basis-Theory/basistheory-go/v3"
	"github.com/Basis-Theory/basistheory-go/v3/internal/testutils"
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

	createdReactorFormula, response, err := apiClient.ReactorFormulasApi.Create(contextWithAPIKey).CreateReactorFormulaRequest(createReactorFormulaRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorFormulasApi Create", t)
	testutils.AssertPropertiesMatch(createdReactorFormula.GetName(), reactorFormulaName, t)
	testutils.AssertPropertiesMatch(createdReactorFormula.GetCode(), reactorFormulaCode, t)

	// GET BY ID
	var reactorFormula *basistheory.ReactorFormula
	reactorFormula, response, err = apiClient.ReactorFormulasApi.GetById(contextWithAPIKey, createdReactorFormula.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorFormulasApi GetById", t)
	testutils.AssertPropertiesMatch(reactorFormula.GetName(), reactorFormulaName, t)
	testutils.AssertPropertiesMatch(reactorFormula.GetCode(), reactorFormulaCode, t)

	// GET LIST
	var reactorFormulas *basistheory.ReactorFormulaPaginatedList
	reactorFormulas, response, err = apiClient.ReactorFormulasApi.Get(contextWithAPIKey).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorFormulasApi sGet", t)
	testutils.AssertPropertiesMatch(reactorFormulas.Data[0].GetName(), reactorFormulaName, t)
	testutils.AssertPropertiesMatch(reactorFormulas.Data[0].GetCode(), reactorFormulaCode, t)

	// UPDATE
	updatedReactorFormulaName := "Go Test Reactor Formula Update"
	updatedReactorFormulaCode := "module.exports = function (req) {return {raw: \"Hello World\"}}"
	updateReactorFormulaRequest := *basistheory.NewUpdateReactorFormulaRequest("private", updatedReactorFormulaName)
	updateReactorFormulaRequest.SetCode(updatedReactorFormulaCode)

	var updatedReactorFormula *basistheory.ReactorFormula
	updatedReactorFormula, response, err = apiClient.ReactorFormulasApi.Update(contextWithAPIKey, createdReactorFormula.GetId()).UpdateReactorFormulaRequest(updateReactorFormulaRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorFormulasApi Update", t)
	testutils.AssertPropertiesMatch(updatedReactorFormula.GetName(), updatedReactorFormulaName, t)
	testutils.AssertPropertiesMatch(updatedReactorFormula.GetCode(), updatedReactorFormulaCode, t)

	// DELETE
	_, err = apiClient.ReactorFormulasApi.Delete(contextWithAPIKey, createdReactorFormula.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorFormulasApi Delete", t)

	_, _, err = apiClient.ReactorFormulasApi.GetById(contextWithAPIKey, createdReactorFormula.GetId()).Execute()

	testutils.AssertNotFound(err, t)
}

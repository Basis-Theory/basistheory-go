package basistheory_test

import (
	"github.com/Basis-Theory/basistheory-go"
	"github.com/Basis-Theory/basistheory-go/internal/testutils"
	"testing"
)

func TestReactorFormulaCRUD(t *testing.T) {
	// SETUP
	apiClient, contextWithAPIKey := testutils.CreateLocalAPIAndContext()

	// CREATE
	reactorFormulaName := "Go Test Reactor Formula"
	reactorFormulaCode := "module.exports = function (req) {return {raw: \"Goodbye World\"}}"

	createReactorFormulaModel := basistheory.CreateReactorFormulaModel{}
	createReactorFormulaModel.SetName(reactorFormulaName)
	createReactorFormulaModel.SetCode(reactorFormulaCode)

	createdReactorFormula, response, err := apiClient.ReactorFormulasApi.ReactorFormulaCreate(contextWithAPIKey).CreateReactorFormulaModel(createReactorFormulaModel).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorFormulaCreate", t)
	testutils.AssertPropertiesMatch(createdReactorFormula.GetName(), reactorFormulaName, t)
	testutils.AssertPropertiesMatch(createdReactorFormula.GetCode(), reactorFormulaCode, t)

	// GET BY ID
	var reactorFormula *basistheory.ReactorFormulaModel
	reactorFormula, response, err = apiClient.ReactorFormulasApi.ReactorFormulaGetById(contextWithAPIKey, createdReactorFormula.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorFormulaGetById", t)
	testutils.AssertPropertiesMatch(reactorFormula.GetName(), reactorFormulaName, t)
	testutils.AssertPropertiesMatch(reactorFormula.GetCode(), reactorFormulaCode, t)

	// GET LIST
	var reactorFormulas *basistheory.ReactorFormulaModelPaginatedList
	reactorFormulas, response, err = apiClient.ReactorFormulasApi.ReactorFormulasGet(contextWithAPIKey).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorFormulasGet", t)
	testutils.AssertPropertiesMatch(reactorFormulas.Data[0].GetName(), reactorFormulaName, t)
	testutils.AssertPropertiesMatch(reactorFormulas.Data[0].GetCode(), reactorFormulaCode, t)

	// UPDATE
	updatedReactorFormulaName := "Go Test Reactor Formula Update"
	updatedReactorFormulaCode := "module.exports = function (req) {return {raw: \"Hello World\"}}"
	updateReactorFormulaModel := basistheory.UpdateReactorFormulaModel{}
	updateReactorFormulaModel.SetName(updatedReactorFormulaName)
	updateReactorFormulaModel.SetCode(updatedReactorFormulaCode)

	var updatedReactorFormula *basistheory.ReactorFormulaModel
	updatedReactorFormula, response, err = apiClient.ReactorFormulasApi.ReactorFormulaUpdate(contextWithAPIKey, createdReactorFormula.GetId()).UpdateReactorFormulaModel(updateReactorFormulaModel).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorFormulaUpdate", t)
	testutils.AssertPropertiesMatch(updatedReactorFormula.GetName(), updatedReactorFormulaName, t)
	testutils.AssertPropertiesMatch(updatedReactorFormula.GetCode(), updatedReactorFormulaCode, t)

	// DELETE
	_, response, err = apiClient.ReactorFormulasApi.ReactorFormulaDelete(contextWithAPIKey, createdReactorFormula.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorFormulaDelete", t)

	_, _, err = apiClient.ReactorFormulasApi.ReactorFormulaGetById(contextWithAPIKey, createdReactorFormula.GetId()).Execute()

	testutils.AssertNotFound(err, t)
}

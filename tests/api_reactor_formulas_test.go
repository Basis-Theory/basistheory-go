package basistheory_test

import (
	"github.com/Basis-Theory/basistheory-go"
	"github.com/Basis-Theory/basistheory-go/internal/testutils"
	"testing"
)

func TestReactorFormulaCRUD(t *testing.T) {
	// CREATE
	apiClient, contextWithAPIKey := testutils.CreateLocalAPIAndContext()
	reactorFormulaName := "Go Test Reactor Formula"
	reactorFormulaCode := "module.exports = function (req) {return {raw: \"Goodbye World\"}}"

	createReactorFormulaModel := basistheory.CreateReactorFormulaModel{}
	createReactorFormulaModel.SetName(reactorFormulaName)
	createReactorFormulaModel.SetCode(reactorFormulaCode)

	createdReactorFormula, response, err := apiClient.ReactorFormulasApi.ReactorFormulaCreate(contextWithAPIKey).CreateReactorFormulaModel(createReactorFormulaModel).Execute()
	testutils.AssertMethodDidNotError(err, response, "ReactorFormulaCreate", t)

	expectedReactorFormulaName := basistheory.NullableString{}
	expectedReactorFormulaName.Set(&reactorFormulaName)
	testutils.AssertPropertiesMatch(createdReactorFormula.Name, expectedReactorFormulaName, t, basistheory.NullableString{})

	expectedReactorFormulaCode := basistheory.NullableString{}
	expectedReactorFormulaCode.Set(&reactorFormulaCode)
	testutils.AssertPropertiesMatch(createdReactorFormula.Code, expectedReactorFormulaCode, t, basistheory.NullableString{})

	// GET BY ID
	var reactorFormula *basistheory.ReactorFormulaModel
	reactorFormula, response, err = apiClient.ReactorFormulasApi.ReactorFormulaGetById(contextWithAPIKey, createdReactorFormula.GetId()).Execute()
	testutils.AssertMethodDidNotError(err, response, "ReactorFormulaGetById", t)

	testutils.AssertPropertiesMatch(reactorFormula.Name, expectedReactorFormulaName, t, basistheory.NullableString{})
	testutils.AssertPropertiesMatch(reactorFormula.Code, expectedReactorFormulaCode, t, basistheory.NullableString{})

	// GET LIST
	var reactorFormulas *basistheory.ReactorFormulaModelPaginatedList
	reactorFormulas, response, err = apiClient.ReactorFormulasApi.ReactorFormulasGet(contextWithAPIKey).Execute()
	testutils.AssertMethodDidNotError(err, response, "ReactorFormulasGet", t)

	testutils.AssertPropertiesMatch(reactorFormulas.Data[0].Name, expectedReactorFormulaName, t, basistheory.NullableString{})
	testutils.AssertPropertiesMatch(reactorFormulas.Data[0].Code, expectedReactorFormulaCode, t, basistheory.NullableString{})

	// UPDATE
	updatedReactorFormulaName := "Go Test Reactor Formula Update"
	updatedReactorFormulaCode := "module.exports = function (req) {return {raw: \"Hello World\"}}"
	updateReactorFormulaModel := basistheory.UpdateReactorFormulaModel{}
	updateReactorFormulaModel.SetName(updatedReactorFormulaName)
	updateReactorFormulaModel.SetCode(updatedReactorFormulaCode)

	var updatedReactorFormula *basistheory.ReactorFormulaModel
	updatedReactorFormula, response, err = apiClient.ReactorFormulasApi.ReactorFormulaUpdate(contextWithAPIKey, createdReactorFormula.GetId()).UpdateReactorFormulaModel(updateReactorFormulaModel).Execute()
	testutils.AssertMethodDidNotError(err, response, "ReactorFormulaUpdate", t)

	expectedUpdatedReactorFormulaName := basistheory.NullableString{}
	expectedUpdatedReactorFormulaName.Set(&updatedReactorFormulaName)
	testutils.AssertPropertiesMatch(updatedReactorFormula.Name, expectedUpdatedReactorFormulaName, t, basistheory.NullableString{})

	expectedUpdatedReactorFormulaCode := basistheory.NullableString{}
	expectedUpdatedReactorFormulaCode.Set(&updatedReactorFormulaCode)
	testutils.AssertPropertiesMatch(updatedReactorFormula.Code, expectedUpdatedReactorFormulaCode, t, basistheory.NullableString{})

	// DELETE
	var deletedReactorFormula *basistheory.ReactorFormulaModel
	deletedReactorFormula, response, err = apiClient.ReactorFormulasApi.ReactorFormulaDelete(contextWithAPIKey, createdReactorFormula.GetId()).Execute()
	testutils.AssertMethodDidNotError(err, response, "ReactorFormulaDelete", t)

	testutils.AssertDeletion[*basistheory.ReactorFormulaModel](deletedReactorFormula, nil, "ReactorFormulaDelete", t)
}

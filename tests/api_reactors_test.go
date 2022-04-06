package basistheory_test

import (
	"github.com/Basis-Theory/basistheory-go"
	"github.com/Basis-Theory/basistheory-go/internal/testutils"
	"testing"
)

func TestReactorCRUD(t *testing.T) {
	// SETUP
	apiClient, contextWithAPIKey := testutils.CreateLocalAPIAndContext()
	reactorFormulaName := "Go Test Reactor Formula"
	reactorFormulaCode := "module.exports = function (req) {return {raw: \"Goodbye World\"}}"

	createReactorFormulaModel := basistheory.CreateReactorFormulaModel{}
	createReactorFormulaModel.SetName(reactorFormulaName)
	createReactorFormulaModel.SetCode(reactorFormulaCode)

	createdReactorFormula, response, err := apiClient.ReactorFormulasApi.ReactorFormulaCreate(contextWithAPIKey).CreateReactorFormulaModel(createReactorFormulaModel).Execute()
	testutils.AssertMethodDidNotError(err, response, "ReactorFormulaCreate", t)

	// CREATE
	reactorName := "Go Test Reactor"

	createReactorModel := basistheory.CreateReactorModel{}
	createReactorModel.SetName(reactorName)
	createReactorModel.SetFormula(*createdReactorFormula)
	var createdReactor *basistheory.ReactorModel
	createdReactor, response, err = apiClient.ReactorsApi.ReactorCreate(contextWithAPIKey).CreateReactorModel(createReactorModel).Execute()
	testutils.AssertMethodDidNotError(err, response, "ReactorCreate", t)

	expectedReactorName := basistheory.NullableString{}
	expectedReactorName.Set(&reactorName)
	testutils.AssertPropertiesMatch(createdReactor.Name, expectedReactorName, t, basistheory.NullableString{})

	testutils.AssertPropertiesMatch(createdReactor.Formula, createdReactorFormula, t, basistheory.NullableString{}, basistheory.NullableTime{})

	// GET BY ID
	var reactor *basistheory.ReactorModel
	reactor, response, err = apiClient.ReactorsApi.ReactorGetById(contextWithAPIKey, createdReactor.GetId()).Execute()
	testutils.AssertMethodDidNotError(err, response, "ReactorGetById", t)

	testutils.AssertPropertiesMatch(reactor.Name, expectedReactorName, t, basistheory.NullableString{})
	testutils.AssertPropertiesMatch(reactor.Formula, createdReactorFormula, t, basistheory.NullableString{}, basistheory.NullableTime{})

	// GET LIST
	var reactors *basistheory.ReactorModelPaginatedList
	reactors, response, err = apiClient.ReactorsApi.ReactorsGet(contextWithAPIKey).Execute()
	testutils.AssertMethodDidNotError(err, response, "ReactorsGet", t)

	testutils.AssertPropertiesMatch(reactors.Data[0].Name, expectedReactorName, t, basistheory.NullableString{})
	testutils.AssertPropertiesMatch(reactors.Data[0].Formula, createdReactorFormula, t, basistheory.NullableString{}, basistheory.NullableTime{})

	// UPDATE
	updatedReactorName := "Go Test Reactor Update"
	updateReactorModel := basistheory.UpdateReactorModel{}
	updateReactorModel.SetName(updatedReactorName)
	updateReactorModel.SetConfiguration(map[string]string{})

	var updatedReactor *basistheory.ReactorModel
	updatedReactor, response, err = apiClient.ReactorsApi.ReactorUpdate(contextWithAPIKey, createdReactor.GetId()).UpdateReactorModel(updateReactorModel).Execute()
	testutils.AssertMethodDidNotError(err, response, "ReactorUpdate", t)

	expectedUpdatedReactorName := basistheory.NullableString{}
	expectedUpdatedReactorName.Set(&updatedReactorName)
	testutils.AssertPropertiesMatch(updatedReactor.Name, expectedUpdatedReactorName, t, basistheory.NullableString{})

	// DELETE
	response, err = apiClient.ReactorsApi.ReactorDelete(contextWithAPIKey, createdReactor.GetId()).Execute()
	testutils.AssertMethodDidNotError(err, response, "ReactorDelete", t)

	_, _, err = apiClient.ReactorsApi.ReactorGetById(contextWithAPIKey, createdReactor.GetId()).Execute()
	testutils.AssertNotFound(err, t)
}

func TestReactorReact(t *testing.T) {
	t.Fatal("Implement test")
}

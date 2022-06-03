package basistheory_test

import (
	"github.com/Basis-Theory/basistheory-go"
	"github.com/Basis-Theory/basistheory-go/internal/testutils"
	"testing"
)

func TestReactorCRUD(t *testing.T) {
	// SETUP
	apiClient, contextWithAPIKey := testutils.CreateApiAndMgmtContext(t)
	reactorFormulaName := "Go Test Reactor Formula"
	reactorFormulaCode := "module.exports = function (req) {return {raw: \"Goodbye World\"}}"

	createReactorFormulaRequest := *basistheory.NewCreateReactorFormulaRequest("private", reactorFormulaName)
	createReactorFormulaRequest.SetCode(reactorFormulaCode)

	createdReactorFormula, response, err := apiClient.ReactorFormulasApi.ReactorFormulasCreate(contextWithAPIKey).CreateReactorFormulaRequest(createReactorFormulaRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorFormulaCreate", t)

	createApplicationRequest := *basistheory.NewCreateApplicationRequest("Go Test App", "server_to_server")

	createdApplication, response, err := apiClient.ApplicationsApi.ApplicationsCreate(contextWithAPIKey).CreateApplicationRequest(createApplicationRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationCreate", t)

	// CREATE
	reactorName := "Go Test Reactor"

	createReactorRequest := *basistheory.NewCreateReactorRequest(reactorName)
	createReactorRequest.SetFormula(*createdReactorFormula)
	createReactorRequest.SetApplication(*createdApplication)
	var createdReactor *basistheory.Reactor
	createdReactor, response, err = apiClient.ReactorsApi.ReactorsCreate(contextWithAPIKey).CreateReactorRequest(createReactorRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorCreate", t)
	testutils.AssertPropertiesMatch(createdReactor.GetName(), reactorName, t)
	testutils.AssertPropertiesMatch(createdReactor.Formula, createdReactorFormula, t, basistheory.NullableString{}, basistheory.NullableTime{})

	// GET BY ID
	var reactor *basistheory.Reactor
	reactor, response, err = apiClient.ReactorsApi.ReactorsGetById(contextWithAPIKey, createdReactor.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorGetById", t)
	testutils.AssertPropertiesMatch(reactor.GetName(), reactorName, t)
	testutils.AssertPropertiesMatch(reactor.Formula, createdReactorFormula, t, basistheory.NullableString{}, basistheory.NullableTime{})

	// GET LIST
	var reactors *basistheory.ReactorPaginatedList
	reactors, response, err = apiClient.ReactorsApi.ReactorsGet(contextWithAPIKey).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorsGet", t)
	testutils.AssertPropertiesMatch(reactors.Data[0].GetName(), reactorName, t)
	testutils.AssertPropertiesMatch(reactors.Data[0].Formula, createdReactorFormula, t, basistheory.NullableString{}, basistheory.NullableTime{})

	// UPDATE
	updatedReactorName := "Go Test Reactor Update"
	updateReactorRequest := *basistheory.NewUpdateReactorRequest(updatedReactorName)
	updateReactorRequest.SetConfiguration(map[string]string{})

	var updatedReactor *basistheory.Reactor
	updatedReactor, response, err = apiClient.ReactorsApi.ReactorsUpdate(contextWithAPIKey, createdReactor.GetId()).UpdateReactorRequest(updateReactorRequest).Execute()
	testutils.AssertMethodDidNotError(err, response, "ReactorUpdate", t)

	testutils.AssertPropertiesMatch(updatedReactor.GetName(), updatedReactorName, t)

	// DELETE
	response, err = apiClient.ReactorsApi.ReactorsDelete(contextWithAPIKey, createdReactor.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorDelete", t)

	_, _, err = apiClient.ReactorsApi.ReactorsGetById(contextWithAPIKey, createdReactor.GetId()).Execute()

	testutils.AssertNotFound(err, t)

	_, _ = apiClient.ReactorFormulasApi.ReactorFormulasDelete(contextWithAPIKey, createdReactorFormula.GetId()).Execute()
	_, _ = apiClient.ApplicationsApi.ApplicationsDelete(contextWithAPIKey, createdApplication.GetId()).Execute()
}

func TestReactorReact(t *testing.T) {
	apiClient, contextWithAPIKey := testutils.CreateApiAndMgmtContext(t)
	reactorFormulaName := "Go Test Reactor Formula"
	reactorFormulaCode := "module.exports = function (req) {return {raw: \"Goodbye World\"}}"

	createReactorFormulaRequest := *basistheory.NewCreateReactorFormulaRequest("private", reactorFormulaName)
	createReactorFormulaRequest.SetCode(reactorFormulaCode)

	createdReactorFormula, response, err := apiClient.ReactorFormulasApi.ReactorFormulasCreate(contextWithAPIKey).CreateReactorFormulaRequest(createReactorFormulaRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorFormulaCreate", t)

	reactorName := "Go Test Reactor"

	createReactorRequest := *basistheory.NewCreateReactorRequest(reactorName)
	createReactorRequest.SetFormula(*createdReactorFormula)
	var createdReactor *basistheory.Reactor
	createdReactor, response, err = apiClient.ReactorsApi.ReactorsCreate(contextWithAPIKey).CreateReactorRequest(createReactorRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorCreate", t)

	apiClient, contextWithAPIKey = testutils.CreateApiAndSrvrContext(t)
	reactRequest := basistheory.ReactRequest{}
	reactRequest.SetArgs(map[string]interface{}{
		"property": "value",
	})
	var reactResponse *basistheory.ReactResponse
	reactResponse, response, err = apiClient.ReactorsApi.ReactorsReact(contextWithAPIKey, createdReactor.GetId()).ReactRequest(reactRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorReact", t)
	testutils.AssertPropertiesMatch(reactResponse.Raw, "Goodbye World", t)

	apiClient, contextWithAPIKey = testutils.CreateApiAndMgmtContext(t)
	response, err = apiClient.ReactorsApi.ReactorsDelete(contextWithAPIKey, createdReactor.GetId()).Execute()
	_, _ = apiClient.ReactorFormulasApi.ReactorFormulasDelete(contextWithAPIKey, createdReactorFormula.GetId()).Execute()
}

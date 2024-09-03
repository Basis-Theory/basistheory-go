package basistheory_test

import (
	"github.com/Basis-Theory/basistheory-go/v6"
	"github.com/Basis-Theory/basistheory-go/v6/internal/testutils"
	"testing"
)

func TestReactorCRUD(t *testing.T) {
	// SETUP
	apiClient, contextWithAPIKey := testutils.CreateApiAndMgmtContext(t)

	applicationPermissions := []string{"token:create"}
	createApplicationRequest := *basistheory.NewCreateApplicationRequest("Go Test App", "private")
	createApplicationRequest.SetPermissions(applicationPermissions)

	createdApplication, response, err := apiClient.ApplicationsApi.Create(contextWithAPIKey).CreateApplicationRequest(createApplicationRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationsApi Create", t)

	// CREATE
	reactorName := "Go Test Reactor"
	code := "module.exports = function (req) {return {raw: \"Goodbye World\"}}"

	createReactorRequest := *basistheory.NewCreateReactorRequest(reactorName, code)
	createReactorRequest.SetApplication(*createdApplication)
	var createdReactor *basistheory.Reactor
	createdReactor, response, err = apiClient.ReactorsApi.Create(contextWithAPIKey).CreateReactorRequest(createReactorRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorsApi Create", t)
	testutils.AssertPropertiesMatch(createdReactor.GetName(), reactorName, t)
	testutils.AssertPropertiesMatch(createdReactor.GetCode(), createReactorRequest.GetCode(), t)

	// GET BY ID
	var reactor *basistheory.Reactor
	reactor, response, err = apiClient.ReactorsApi.GetById(contextWithAPIKey, createdReactor.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorsApi GetById", t)
	testutils.AssertPropertiesMatch(reactor.GetName(), reactorName, t)
	testutils.AssertPropertiesMatch(reactor.GetCode(), createReactorRequest.GetCode(), t)

	// GET LIST
	var reactors *basistheory.ReactorPaginatedList
	reactors, response, err = apiClient.ReactorsApi.Get(contextWithAPIKey).Id([]string{createdReactor.GetId()}).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorsApi sGet", t)
	testutils.AssertPropertiesMatch(reactors.Data[0].GetName(), reactorName, t)
	testutils.AssertPropertiesMatch(reactors.Data[0].GetCode(), createReactorRequest.GetCode(), t)

	// UPDATE
	updatedReactorName := "Go Test Reactor Update"
	updateReactorRequest := *basistheory.NewUpdateReactorRequest(updatedReactorName, code)
	updateReactorRequest.SetConfiguration(map[string]string{})

	var updatedReactor *basistheory.Reactor
	updatedReactor, response, err = apiClient.ReactorsApi.Update(contextWithAPIKey, createdReactor.GetId()).UpdateReactorRequest(updateReactorRequest).Execute()
	testutils.AssertMethodDidNotError(err, response, "ReactorsApi Update", t)

	testutils.AssertPropertiesMatch(updatedReactor.GetName(), updatedReactorName, t)
	testutils.AssertPropertiesMatch(updatedReactor.GetCode(), updateReactorRequest.GetCode(), t)

	// DELETE
	response, err = apiClient.ReactorsApi.Delete(contextWithAPIKey, createdReactor.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorsApi Delete", t)

	_, _, err = apiClient.ReactorsApi.GetById(contextWithAPIKey, createdReactor.GetId()).Execute()

	testutils.AssertNotFound(err, t)

	_, _ = apiClient.ApplicationsApi.Delete(contextWithAPIKey, createdApplication.GetId()).Execute()
}

func TestReactorReact(t *testing.T) {
	apiClient, contextWithAPIKey := testutils.CreateApiAndMgmtContext(t)

	reactorName := "Go Test Reactor"

	createReactorRequest := *basistheory.NewCreateReactorRequest(reactorName, "module.exports = function (req) {return {raw: \"Goodbye World\"}}")
	var createdReactor *basistheory.Reactor
	createdReactor, response, err := apiClient.ReactorsApi.Create(contextWithAPIKey).CreateReactorRequest(createReactorRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorsApi Create", t)

	apiClient, contextWithAPIKey = testutils.CreateApiAndSrvrContext(t)
	reactRequest := basistheory.ReactRequest{}
	reactRequest.SetArgs(map[string]interface{}{
		"property": "value",
	})
	var reactResponse *basistheory.ReactResponse
	reactResponse, response, err = apiClient.ReactorsApi.React(contextWithAPIKey, createdReactor.GetId()).ReactRequest(reactRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorsApi React", t)
	testutils.AssertPropertiesMatch(reactResponse.Raw, "Goodbye World", t)

	apiClient, contextWithAPIKey = testutils.CreateApiAndMgmtContext(t)
	response, err = apiClient.ReactorsApi.Delete(contextWithAPIKey, createdReactor.GetId()).Execute()
}

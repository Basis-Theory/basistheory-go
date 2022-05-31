package basistheory_test

import (
	"github.com/Basis-Theory/basistheory-go"
	"github.com/Basis-Theory/basistheory-go/internal/testutils"
	"testing"
)

func TestInboundProxiesCRUD(t *testing.T) {
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

	reactorName := "Go Test Reactor"
	createReactorRequest := *basistheory.NewCreateReactorRequest(reactorName)
	createReactorRequest.SetFormula(*createdReactorFormula)
	createReactorRequest.SetApplication(*createdApplication)
	var createdReactor *basistheory.Reactor
	createdReactor, response, err = apiClient.ReactorsApi.ReactorsCreate(contextWithAPIKey).CreateReactorRequest(createReactorRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorCreate", t)

	// CREATE
	inboundProxyName := "Go Test Inbound Proxy"
	inboundProxyDestinationUrl := "http://httpbin.org/post"

	createInboundProxyRequest := *basistheory.NewCreateInboundProxyRequest(inboundProxyName, inboundProxyDestinationUrl, createdReactor.GetId())

	createdInboundProxy, response, err := apiClient.InboundProxiesApi.InboundProxiesCreate(contextWithAPIKey).CreateInboundProxyRequest(createInboundProxyRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "InboundProxiesCreate", t)
	testutils.AssertPropertiesMatch(createdInboundProxy.GetName(), inboundProxyName, t)
	testutils.AssertPropertiesMatch(createdInboundProxy.GetDestinationUrl(), inboundProxyDestinationUrl, t)
	testutils.AssertPropertiesMatch(createdInboundProxy.GetRequestReactorId(), createdReactor.GetId(), t)

	// GET BY ID
	var inboundProxy *basistheory.InboundProxy
	inboundProxy, response, err = apiClient.InboundProxiesApi.InboundProxiesGetById(contextWithAPIKey, createdInboundProxy.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "InboundProxiesGetById", t)
	testutils.AssertPropertiesMatch(inboundProxy.GetName(), inboundProxyName, t)
	testutils.AssertPropertiesMatch(inboundProxy.GetDestinationUrl(), inboundProxyDestinationUrl, t)
	testutils.AssertPropertiesMatch(inboundProxy.GetRequestReactorId(), createdReactor.GetId(), t)

	// GET LIST
	var inboundProxies *basistheory.InboundProxyPaginatedList
	inboundProxies, response, err = apiClient.InboundProxiesApi.InboundProxiesGet(contextWithAPIKey).Execute()

	testutils.AssertMethodDidNotError(err, response, "InboundProxiesGet", t)
	testutils.AssertPropertiesMatch(inboundProxies.Data[0].GetName(), inboundProxyName, t)
	testutils.AssertPropertiesMatch(inboundProxies.Data[0].GetDestinationUrl(), inboundProxyDestinationUrl, t)
	testutils.AssertPropertiesMatch(inboundProxies.Data[0].GetRequestReactorId(), createdReactor.GetId(), t)

	// UPDATE
	updatedInboundProxyName := "Go Test Inbound Proxy Update"
	updatedInboundProxyDestinationUrl := "https://httpbin.org/post"
	updateInboundProxyRequest := *basistheory.NewUpdateInboundProxyRequest(updatedInboundProxyName, updatedInboundProxyDestinationUrl, createdReactor.GetId())

	var updatedInboundProxy *basistheory.InboundProxy
	updatedInboundProxy, response, err = apiClient.InboundProxiesApi.InboundProxiesUpdate(contextWithAPIKey, createdInboundProxy.GetId()).UpdateInboundProxyRequest(updateInboundProxyRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "InboundProxiesUpdate", t)
	testutils.AssertPropertiesMatch(updatedInboundProxy.GetName(), updatedInboundProxyName, t)
	testutils.AssertPropertiesMatch(updatedInboundProxy.GetDestinationUrl(), updatedInboundProxyDestinationUrl, t)
	testutils.AssertPropertiesMatch(updatedInboundProxy.GetRequestReactorId(), createdReactor.GetId(), t)

	// DELETE
	_, err = apiClient.InboundProxiesApi.InboundProxiesDelete(contextWithAPIKey, createdInboundProxy.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "InboundProxiesDelete", t)

	_, _, err = apiClient.InboundProxiesApi.InboundProxiesGetById(contextWithAPIKey, createdInboundProxy.GetId()).Execute()

	testutils.AssertNotFound(err, t)

	_, _ = apiClient.ReactorsApi.ReactorsDelete(contextWithAPIKey, createdReactor.GetId()).Execute()
	_, _ = apiClient.ReactorFormulasApi.ReactorFormulasDelete(contextWithAPIKey, createdReactorFormula.GetId()).Execute()
	_, _ = apiClient.ApplicationsApi.ApplicationsDelete(contextWithAPIKey, createdApplication.GetId()).Execute()
}

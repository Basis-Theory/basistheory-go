package basistheory_test

import (
	"github.com/Basis-Theory/basistheory-go"
	"github.com/Basis-Theory/basistheory-go/internal/testutils"
	"testing"
)

func TestProxiesCRUD(t *testing.T) {
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
	proxyName := "Go Test  Proxy"
	proxyDestinationUrl := "http://httpbin.org/post"

	createProxyRequest := *basistheory.NewCreateProxyRequest(proxyName, proxyDestinationUrl, createdReactor.GetId())

	createdProxy, response, err := apiClient.ProxiesApi.ProxiesCreate(contextWithAPIKey).CreateProxyRequest(createProxyRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ProxiesCreate", t)
	testutils.AssertPropertiesMatch(createdProxy.GetName(), proxyName, t)
	testutils.AssertPropertiesMatch(createdProxy.GetDestinationUrl(), proxyDestinationUrl, t)
	testutils.AssertPropertiesMatch(createdProxy.GetRequestReactorId(), createdReactor.GetId(), t)

	// GET BY ID
	var proxy *basistheory.Proxy
	proxy, response, err = apiClient.ProxiesApi.ProxiesGetById(contextWithAPIKey, createdProxy.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ProxiesGetById", t)
	testutils.AssertPropertiesMatch(proxy.GetName(), proxyName, t)
	testutils.AssertPropertiesMatch(proxy.GetDestinationUrl(), proxyDestinationUrl, t)
	testutils.AssertPropertiesMatch(proxy.GetRequestReactorId(), createdReactor.GetId(), t)

	// GET LIST
	var proxies *basistheory.ProxyPaginatedList
	proxies, response, err = apiClient.ProxiesApi.ProxiesGet(contextWithAPIKey).Execute()

	testutils.AssertMethodDidNotError(err, response, "ProxiesGet", t)
	testutils.AssertPropertiesMatch(proxies.Data[0].GetName(), proxyName, t)
	testutils.AssertPropertiesMatch(proxies.Data[0].GetDestinationUrl(), proxyDestinationUrl, t)
	testutils.AssertPropertiesMatch(proxies.Data[0].GetRequestReactorId(), createdReactor.GetId(), t)

	// UPDATE
	updatedProxyName := "Go Test  Proxy Update"
	updatedProxyDestinationUrl := "https://httpbin.org/post"
	updateProxyRequest := *basistheory.NewUpdateProxyRequest(updatedProxyName, updatedProxyDestinationUrl, createdReactor.GetId())

	var updatedProxy *basistheory.Proxy
	updatedProxy, response, err = apiClient.ProxiesApi.ProxiesUpdate(contextWithAPIKey, createdProxy.GetId()).UpdateProxyRequest(updateProxyRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ProxiesUpdate", t)
	testutils.AssertPropertiesMatch(updatedProxy.GetName(), updatedProxyName, t)
	testutils.AssertPropertiesMatch(updatedProxy.GetDestinationUrl(), updatedProxyDestinationUrl, t)
	testutils.AssertPropertiesMatch(updatedProxy.GetRequestReactorId(), createdReactor.GetId(), t)

	// DELETE
	_, err = apiClient.ProxiesApi.ProxiesDelete(contextWithAPIKey, createdProxy.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ProxiesDelete", t)

	_, _, err = apiClient.ProxiesApi.ProxiesGetById(contextWithAPIKey, createdProxy.GetId()).Execute()

	testutils.AssertNotFound(err, t)

	_, _ = apiClient.ReactorsApi.ReactorsDelete(contextWithAPIKey, createdReactor.GetId()).Execute()
	_, _ = apiClient.ReactorFormulasApi.ReactorFormulasDelete(contextWithAPIKey, createdReactorFormula.GetId()).Execute()
	_, _ = apiClient.ApplicationsApi.ApplicationsDelete(contextWithAPIKey, createdApplication.GetId()).Execute()
}

package basistheory_test

import (
	"github.com/Basis-Theory/basistheory-go/v6"
	"github.com/Basis-Theory/basistheory-go/v6/internal/testutils"
	"testing"
)

func TestProxiesCRUD(t *testing.T) {
	// SETUP
	apiClient, contextWithAPIKey := testutils.CreateApiAndMgmtContext(t)

	applicationPermissions := []string{"token:create"}
	createApplicationRequest := *basistheory.NewCreateApplicationRequest("Go Test App", "private")
	createApplicationRequest.SetPermissions(applicationPermissions)

	createdApplication, response, err := apiClient.ApplicationsApi.Create(contextWithAPIKey).CreateApplicationRequest(createApplicationRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationsApi Create", t)

	reactorName := "Go Test Reactor"
	createReactorRequest := *basistheory.NewCreateReactorRequest(reactorName, "module.exports = function (req) {return {raw: \"Goodbye World\"}}")
	createReactorRequest.SetApplication(*createdApplication)
	var createdReactor *basistheory.Reactor
	createdReactor, response, err = apiClient.ReactorsApi.Create(contextWithAPIKey).CreateReactorRequest(createReactorRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ReactorsApi Create", t)

	// CREATE
	proxyName := "Go Test  Proxy"
	proxyDestinationUrl := "https://httpbin.org/post"

	createProxyRequest := *basistheory.NewCreateProxyRequest(proxyName, proxyDestinationUrl)
	createProxyRequest.SetRequestReactorId(createdReactor.GetId())
	createProxyRequest.SetRequireAuth(true)

	createdProxy, response, err := apiClient.ProxiesApi.Create(contextWithAPIKey).CreateProxyRequest(createProxyRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ProxiesApi Create", t)
	testutils.AssertPropertiesMatch(createdProxy.GetName(), proxyName, t)
	testutils.AssertPropertiesMatch(createdProxy.GetDestinationUrl(), proxyDestinationUrl, t)
	testutils.AssertPropertiesMatch(createdProxy.GetRequestReactorId(), createdReactor.GetId(), t)
	testutils.AssertPropertiesMatch(createdProxy.GetRequireAuth(), true, t)

	// GET BY ID
	var proxy *basistheory.Proxy
	proxy, response, err = apiClient.ProxiesApi.GetById(contextWithAPIKey, createdProxy.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ProxiesApi GetById", t)
	testutils.AssertPropertiesMatch(proxy.GetName(), proxyName, t)
	testutils.AssertPropertiesMatch(proxy.GetDestinationUrl(), proxyDestinationUrl, t)
	testutils.AssertPropertiesMatch(proxy.GetRequestReactorId(), createdReactor.GetId(), t)
	testutils.AssertPropertiesMatch(proxy.GetRequireAuth(), true, t)

	// GET LIST
	var proxies *basistheory.ProxyPaginatedList
	proxies, response, err = apiClient.ProxiesApi.Get(contextWithAPIKey).Id([]string{createdProxy.GetId()}).Execute()

	testutils.AssertMethodDidNotError(err, response, "ProxiesApi Get", t)
	testutils.AssertPropertiesMatch(proxies.Data[0].GetName(), proxyName, t)
	testutils.AssertPropertiesMatch(proxies.Data[0].GetDestinationUrl(), proxyDestinationUrl, t)
	testutils.AssertPropertiesMatch(proxies.Data[0].GetRequestReactorId(), createdReactor.GetId(), t)
	testutils.AssertPropertiesMatch(proxies.Data[0].GetRequireAuth(), true, t)

	// UPDATE
	updatedProxyName := "Go Test  Proxy Update"
	updatedProxyDestinationUrl := "https://httpbin.org/post"
	updateProxyRequest := *basistheory.NewUpdateProxyRequest(updatedProxyName, updatedProxyDestinationUrl)
	updateProxyRequest.SetRequestReactorId(createdReactor.GetId())
	updateProxyRequest.SetRequireAuth(false)

	var updatedProxy *basistheory.Proxy
	updatedProxy, response, err = apiClient.ProxiesApi.Update(contextWithAPIKey, createdProxy.GetId()).UpdateProxyRequest(updateProxyRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ProxiesApi Update", t)
	testutils.AssertPropertiesMatch(updatedProxy.GetName(), updatedProxyName, t)
	testutils.AssertPropertiesMatch(updatedProxy.GetDestinationUrl(), updatedProxyDestinationUrl, t)
	testutils.AssertPropertiesMatch(updatedProxy.GetRequestReactorId(), createdReactor.GetId(), t)
	testutils.AssertPropertiesMatch(updatedProxy.GetRequireAuth(), false, t)

	// DELETE
	_, err = apiClient.ProxiesApi.Delete(contextWithAPIKey, createdProxy.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ProxiesApi Delete", t)

	_, _, err = apiClient.ProxiesApi.GetById(contextWithAPIKey, createdProxy.GetId()).Execute()

	testutils.AssertNotFound(err, t)

	_, _ = apiClient.ReactorsApi.Delete(contextWithAPIKey, createdReactor.GetId()).Execute()
	_, _ = apiClient.ApplicationsApi.Delete(contextWithAPIKey, createdApplication.GetId()).Execute()
}

package basistheory_test

import (
	"github.com/Basis-Theory/basistheory-go/v6"
	"github.com/Basis-Theory/basistheory-go/v6/internal/testutils"
	"testing"
)

func TestTenantCRUD(t *testing.T) {
	// SETUP
	apiClient, contextWithAPIKey := testutils.CreateApiAndMgmtContext(t)

	// GET BY SELF
	var tenant *basistheory.Tenant
	tenant, response, err := apiClient.TenantsApi.Get(contextWithAPIKey).Execute()

	testutils.AssertMethodDidNotError(err, response, "TenantsApi Get", t)

	// UPDATE
	updatedTenantSettings := map[string]string{
		"deduplicate_tokens": "true",
	}
	updateTenantRequest := *basistheory.NewUpdateTenantRequest(tenant.GetName())
	updateTenantRequest.SetSettings(updatedTenantSettings)

	var updatedTenant *basistheory.Tenant
	updatedTenant, response, err = apiClient.TenantsApi.Update(contextWithAPIKey).UpdateTenantRequest(updateTenantRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "TenantsApi Update", t)
	testutils.AssertPropertiesMatch(updatedTenant.GetName(), tenant.GetName(), t)
	testutils.AssertPropertiesMatch(updatedTenant.GetSettings()["deduplicate_tokens"], updatedTenantSettings["deduplicate_tokens"], t)
}

package basistheory_test

import (
	"github.com/Basis-Theory/basistheory-go/v6/internal/testutils"
	"testing"
)

func TestPermissionsCRUD(t *testing.T) {
	// SETUP
	apiClient, contextWithAPIKey := testutils.CreateApiAndSrvrContext(t)

	// GET Permissions
	_, response, err := apiClient.PermissionsApi.Get(contextWithAPIKey).Execute()

	testutils.AssertMethodDidNotError(err, response, "PermissionsApi Get", t)
}

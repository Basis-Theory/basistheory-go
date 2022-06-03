package basistheory_test

import (
	"context"
	"github.com/Basis-Theory/basistheory-go/v2"
	"github.com/Basis-Theory/basistheory-go/v2/internal/testutils"
	"testing"
)

func getApplicationNameAndType() (string, string) {
	return "Go Test App", "server_to_server"
}

func TestApplicationCRUD(t *testing.T) {
	// SETUP
	apiClient, contextWithAPIKey := testutils.CreateApiAndMgmtContext(t)

	// CREATE
	applicationName, applicationType := getApplicationNameAndType()

	applicationPermissions := []string{"token:pci:create"}
	createApplicationRequest := *basistheory.NewCreateApplicationRequest(applicationName, applicationType)
	createApplicationRequest.SetPermissions(applicationPermissions)

	createdApplication, response, err := apiClient.ApplicationsApi.ApplicationsCreate(contextWithAPIKey).CreateApplicationRequest(createApplicationRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationCreate", t)

	// GET BY ID
	var application *basistheory.Application
	application, response, err = apiClient.ApplicationsApi.ApplicationsGetById(contextWithAPIKey, createdApplication.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationGetById", t)
	testutils.AssertPropertiesMatch(application.GetName(), applicationName, t)
	testutils.AssertPropertiesMatch(application.GetType(), applicationType, t)
	testutils.AssertPropertiesMatch(application.Permissions, applicationPermissions, t)

	// GET LIST
	var applications *basistheory.ApplicationPaginatedList
	applications, response, err = apiClient.ApplicationsApi.ApplicationsGet(contextWithAPIKey).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationsGet", t)
	testutils.AssertPropertiesMatch(applications.Data[0].GetName(), applicationName, t)
	testutils.AssertPropertiesMatch(applications.Data[0].GetType(), applicationType, t)
	testutils.AssertPropertiesMatch(applications.Data[0].GetPermissions(), applicationPermissions, t)

	// UPDATE
	updatedApplicationName := "Updated Name"
	updatedApplicationPermissions := []string{"token:general:read:low"}
	updateApplicationRequest := basistheory.UpdateApplicationRequest{}
	updateApplicationRequest.SetName(updatedApplicationName)
	updateApplicationRequest.SetPermissions(updatedApplicationPermissions)

	var updatedApplication *basistheory.Application
	updatedApplication, response, err = apiClient.ApplicationsApi.ApplicationsUpdate(contextWithAPIKey, createdApplication.GetId()).UpdateApplicationRequest(updateApplicationRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationUpdate", t)
	testutils.AssertPropertiesMatch(updatedApplication.GetName(), updatedApplicationName, t)
	testutils.AssertPropertiesMatch(updatedApplication.GetPermissions(), updatedApplicationPermissions, t)

	// DELETE
	response, err = apiClient.ApplicationsApi.ApplicationsDelete(contextWithAPIKey, createdApplication.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationDelete", t)

	_, _, err = apiClient.ApplicationsApi.ApplicationsGetById(contextWithAPIKey, createdApplication.GetId()).Execute()

	testutils.AssertNotFound(err, t)
}

func TestApplicationRegenerate(t *testing.T) {
	apiClient, contextWithAPIKey := testutils.CreateApiAndMgmtContext(t)
	applicationName, applicationType := getApplicationNameAndType()

	applicationPermissions := []string{"token:pci:create"}
	createApplicationRequest := *basistheory.NewCreateApplicationRequest(applicationName, applicationType)
	createApplicationRequest.SetPermissions(applicationPermissions)

	createdApplication, response, err := apiClient.ApplicationsApi.ApplicationsCreate(contextWithAPIKey).CreateApplicationRequest(createApplicationRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationCreate", t)

	var regeneratedApplication *basistheory.Application
	regeneratedApplication, response, err = apiClient.ApplicationsApi.ApplicationsRegenerateKey(contextWithAPIKey, createdApplication.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationRegenerate", t)
	testutils.AssertPropertiesDoNotMatch(regeneratedApplication.GetKey(), createdApplication.GetKey(), t)

	_, _ = apiClient.ApplicationsApi.ApplicationsDelete(contextWithAPIKey, createdApplication.GetId()).Execute()
}

func TestApplicationKey(t *testing.T) {
	apiClient, contextWithAPIKey := testutils.CreateApiAndMgmtContext(t)
	applicationName, applicationType := getApplicationNameAndType()

	applicationPermissions := []string{"token:pci:create"}
	createApplicationRequest := *basistheory.NewCreateApplicationRequest(applicationName, applicationType)
	createApplicationRequest.SetPermissions(applicationPermissions)

	createdApplication, response, err := apiClient.ApplicationsApi.ApplicationsCreate(contextWithAPIKey).CreateApplicationRequest(createApplicationRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationCreate", t)

	contextWithCreatedAppAPIKey := context.WithValue(context.Background(), basistheory.ContextAPIKeys, map[string]basistheory.APIKey{
		"ApiKey": {Key: *createdApplication.Key.Get()},
	})

	var applicationFromApplicationKey *basistheory.Application
	applicationFromApplicationKey, response, err = apiClient.ApplicationsApi.ApplicationsGetByKey(contextWithCreatedAppAPIKey).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationKey", t)

	applicationFromApplicationKey.SetKey(*createdApplication.Key.Get())

	testutils.AssertPropertiesMatch(applicationFromApplicationKey, createdApplication, t, basistheory.NullableString{}, basistheory.NullableTime{})

	_, _ = apiClient.ApplicationsApi.ApplicationsDelete(contextWithAPIKey, createdApplication.GetId()).Execute()
}

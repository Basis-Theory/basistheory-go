package basistheory_test

import (
	"context"
	"github.com/Basis-Theory/basistheory-go/v3/internal/testutils"
	"testing"
)

func getApplicationNameAndType() (string, string) {
	return "Go Test App", "private"
}

func TestApplicationCRUD(t *testing.T) {
	// SETUP
	apiClient, contextWithAPIKey := testutils.CreateApiAndMgmtContext(t)

	// CREATE
	applicationName, applicationType := getApplicationNameAndType()

	applicationPermissions := []string{"token:create"}
	createApplicationRequest := *basistheory.NewCreateApplicationRequest(applicationName, applicationType)
	createApplicationRequest.SetPermissions(applicationPermissions)

	createdApplication, response, err := apiClient.ApplicationsApi.Create(contextWithAPIKey).CreateApplicationRequest(createApplicationRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationsApi Create", t)

	// GET BY ID
	var application *basistheory.Application
	application, response, err = apiClient.ApplicationsApi.GetById(contextWithAPIKey, createdApplication.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationsApi GetById", t)
	testutils.AssertPropertiesMatch(application.GetName(), applicationName, t)
	testutils.AssertPropertiesMatch(application.GetType(), applicationType, t)
	testutils.AssertPropertiesMatch(application.Permissions, applicationPermissions, t)

	// GET LIST
	var applications *basistheory.ApplicationPaginatedList
	applications, response, err = apiClient.ApplicationsApi.Get(contextWithAPIKey).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationsApi sGet", t)
	testutils.AssertPropertiesMatch(applications.Data[0].GetName(), applicationName, t)
	testutils.AssertPropertiesMatch(applications.Data[0].GetType(), applicationType, t)
	testutils.AssertPropertiesMatch(applications.Data[0].GetPermissions(), applicationPermissions, t)

	// UPDATE
	updatedApplicationName := "Updated Name"
	updatedApplicationPermissions := []string{"token:read"}
	updateApplicationRequest := basistheory.UpdateApplicationRequest{}
	updateApplicationRequest.SetName(updatedApplicationName)
	updateApplicationRequest.SetPermissions(updatedApplicationPermissions)

	var updatedApplication *basistheory.Application
	updatedApplication, response, err = apiClient.ApplicationsApi.Update(contextWithAPIKey, createdApplication.GetId()).UpdateApplicationRequest(updateApplicationRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationsApi Update", t)
	testutils.AssertPropertiesMatch(updatedApplication.GetName(), updatedApplicationName, t)
	testutils.AssertPropertiesMatch(updatedApplication.GetPermissions(), updatedApplicationPermissions, t)

	// DELETE
	response, err = apiClient.ApplicationsApi.Delete(contextWithAPIKey, createdApplication.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationsApi Delete", t)

	_, _, err = apiClient.ApplicationsApi.GetById(contextWithAPIKey, createdApplication.GetId()).Execute()

	testutils.AssertNotFound(err, t)
}

func TestApplicationRegenerate(t *testing.T) {
	apiClient, contextWithAPIKey := testutils.CreateApiAndMgmtContext(t)
	applicationName, applicationType := getApplicationNameAndType()

	applicationPermissions := []string{"token:create"}
	createApplicationRequest := *basistheory.NewCreateApplicationRequest(applicationName, applicationType)
	createApplicationRequest.SetPermissions(applicationPermissions)

	createdApplication, response, err := apiClient.ApplicationsApi.Create(contextWithAPIKey).CreateApplicationRequest(createApplicationRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationsApi Create", t)

	var regeneratedApplication *basistheory.Application
	regeneratedApplication, response, err = apiClient.ApplicationsApi.RegenerateKey(contextWithAPIKey, createdApplication.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationsApi Regenerate", t)
	testutils.AssertPropertiesDoNotMatch(regeneratedApplication.GetKey(), createdApplication.GetKey(), t)

	_, _ = apiClient.ApplicationsApi.Delete(contextWithAPIKey, createdApplication.GetId()).Execute()
}

func TestApplicationKey(t *testing.T) {
	apiClient, contextWithAPIKey := testutils.CreateApiAndMgmtContext(t)
	applicationName, applicationType := getApplicationNameAndType()

	applicationPermissions := []string{"token:create"}
	createApplicationRequest := *basistheory.NewCreateApplicationRequest(applicationName, applicationType)
	createApplicationRequest.SetPermissions(applicationPermissions)

	createdApplication, response, err := apiClient.ApplicationsApi.Create(contextWithAPIKey).CreateApplicationRequest(createApplicationRequest).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationsApi Create", t)

	contextWithCreatedAppAPIKey := context.WithValue(context.Background(), basistheory.ContextAPIKeys, map[string]basistheory.APIKey{
		"ApiKey": {Key: *createdApplication.Key.Get()},
	})

	var applicationFromApplicationKey *basistheory.Application
	applicationFromApplicationKey, response, err = apiClient.ApplicationsApi.GetByKey(contextWithCreatedAppAPIKey).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationsApi Key", t)

	applicationFromApplicationKey.SetKey(*createdApplication.Key.Get())

	testutils.AssertPropertiesMatch(applicationFromApplicationKey, createdApplication, t, basistheory.NullableString{}, basistheory.NullableTime{})

	_, _ = apiClient.ApplicationsApi.Delete(contextWithAPIKey, createdApplication.GetId()).Execute()
}

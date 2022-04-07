package basistheory_test

import (
	"context"
	"github.com/Basis-Theory/basistheory-go"
	"github.com/Basis-Theory/basistheory-go/internal/testutils"
	"testing"
)

func getApplicationNameAndType() (string, string) {
	return "Go Test App", "server_to_server"
}

func TestApplicationCRUD(t *testing.T) {
	// SETUP
	apiClient, contextWithAPIKey := testutils.CreateLocalAPIAndContext()

	// CREATE
	applicationName, applicationType := getApplicationNameAndType()

	applicationPermissions := []string{"token:pci:create"}
	createApplicationModel := *basistheory.NewCreateApplicationModel()
	createApplicationModel.SetName(applicationName)
	createApplicationModel.SetType(applicationType)
	createApplicationModel.SetPermissions(applicationPermissions)

	createdApplication, response, err := apiClient.ApplicationsApi.ApplicationCreate(contextWithAPIKey).CreateApplicationModel(createApplicationModel).Execute()
	testutils.AssertMethodDidNotError(err, response, "ApplicationCreate", t)

	// GET BY ID
	var application *basistheory.ApplicationModel
	application, response, err = apiClient.ApplicationsApi.ApplicationGetById(contextWithAPIKey, createdApplication.GetId()).Execute()
	testutils.AssertMethodDidNotError(err, response, "ApplicationGetById", t)

	expectedApplicationName := basistheory.NullableString{}
	expectedApplicationName.Set(&applicationName)
	testutils.AssertPropertiesMatch(application.Name, expectedApplicationName, t, basistheory.NullableString{})

	expectedApplicationType := basistheory.NullableString{}
	expectedApplicationType.Set(&applicationType)
	testutils.AssertPropertiesMatch(application.Type, expectedApplicationType, t, basistheory.NullableString{})

	testutils.AssertPropertiesMatch(application.Permissions, applicationPermissions, t)

	// GET LIST
	var applications *basistheory.ApplicationModelPaginatedList
	applications, response, err = apiClient.ApplicationsApi.ApplicationsGet(contextWithAPIKey).Execute()
	testutils.AssertMethodDidNotError(err, response, "ApplicationsGet", t)

	testutils.AssertPropertiesMatch(applications.Data[0].Name, expectedApplicationName, t, basistheory.NullableString{})
	testutils.AssertPropertiesMatch(applications.Data[0].Type, expectedApplicationType, t, basistheory.NullableString{})
	testutils.AssertPropertiesMatch(applications.Data[0].Permissions, applicationPermissions, t)

	// UPDATE
	updatedApplicationName := "Updated Name"
	updatedApplicationPermissions := []string{"token:general:read:low"}
	updateApplicationModel := basistheory.UpdateApplicationModel{}
	updateApplicationModel.SetName(updatedApplicationName)
	updateApplicationModel.SetPermissions(updatedApplicationPermissions)

	var updatedApplication *basistheory.ApplicationModel
	updatedApplication, response, err = apiClient.ApplicationsApi.ApplicationUpdate(contextWithAPIKey, createdApplication.GetId()).UpdateApplicationModel(updateApplicationModel).Execute()
	testutils.AssertMethodDidNotError(err, response, "ApplicationUpdate", t)

	updatedApplicationNameNullableString := basistheory.NullableString{}
	updatedApplicationNameNullableString.Set(&updatedApplicationName)
	testutils.AssertPropertiesMatch(updatedApplication.Name, updatedApplicationNameNullableString, t, basistheory.NullableString{})
	testutils.AssertPropertiesMatch(updatedApplication.Permissions, updatedApplicationPermissions, t)

	// DELETE
	response, err = apiClient.ApplicationsApi.ApplicationDelete(contextWithAPIKey, createdApplication.GetId()).Execute()
	testutils.AssertMethodDidNotError(err, response, "ApplicationDelete", t)

	_, _, err = apiClient.ApplicationsApi.ApplicationGetById(contextWithAPIKey, createdApplication.GetId()).Execute()
	testutils.AssertNotFound(err, t)
}

func TestApplicationRegenerate(t *testing.T) {
	apiClient, contextWithAPIKey := testutils.CreateLocalAPIAndContext()
	applicationName, applicationType := getApplicationNameAndType()

	applicationPermissions := []string{"token:pci:create"}
	createApplicationModel := *basistheory.NewCreateApplicationModel()
	createApplicationModel.SetName(applicationName)
	createApplicationModel.SetType(applicationType)
	createApplicationModel.SetPermissions(applicationPermissions)

	createdApplication, response, err := apiClient.ApplicationsApi.ApplicationCreate(contextWithAPIKey).CreateApplicationModel(createApplicationModel).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationCreate", t)

	var regeneratedApplication *basistheory.ApplicationModel
	regeneratedApplication, response, err = apiClient.ApplicationsApi.ApplicationRegenerate(contextWithAPIKey, createdApplication.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationRegenerate", t)

	testutils.AssertPropertiesDoNotMatch(regeneratedApplication.Key, createdApplication.Key, t, basistheory.NullableString{})
}

func TestApplicationKey(t *testing.T) {
	apiClient, contextWithAPIKey := testutils.CreateLocalAPIAndContext()
	applicationName, applicationType := getApplicationNameAndType()

	applicationPermissions := []string{"token:pci:create"}
	createApplicationModel := *basistheory.NewCreateApplicationModel()
	createApplicationModel.SetName(applicationName)
	createApplicationModel.SetType(applicationType)
	createApplicationModel.SetPermissions(applicationPermissions)

	createdApplication, response, err := apiClient.ApplicationsApi.ApplicationCreate(contextWithAPIKey).CreateApplicationModel(createApplicationModel).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationCreate", t)

	contextWithCreatedAppAPIKey := context.WithValue(context.Background(), basistheory.ContextAPIKeys, map[string]basistheory.APIKey{
		"ApiKey": {Key: *createdApplication.Key.Get()},
	})

	var applicationFromApplicationKey *basistheory.ApplicationModel
	applicationFromApplicationKey, response, err = apiClient.ApplicationsApi.ApplicationKey(contextWithCreatedAppAPIKey).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationKey", t)

	applicationFromApplicationKey.SetKey(*createdApplication.Key.Get())
	testutils.AssertPropertiesMatch(applicationFromApplicationKey, createdApplication, t, basistheory.NullableString{}, basistheory.NullableTime{})
}

package basistheory_test

import (
	"context"
	"github.com/Basis-Theory/basistheory-go"
	"github.com/Basis-Theory/basistheory-go/internal/testutils"
	"testing"
)

const _applicationName = "Go Test App"
const _applicationType = "server_to_server"

func TestApplicationCRUD(t *testing.T) {
	apiClient, contextWithAPIKey := testutils.CreateLocalAPIAndContext()

	applicationName := _applicationName
	applicationType := _applicationType
	applicationPermissions := []string{"token:pci:create"}
	createApplicationModel := *basistheory.NewCreateApplicationModel()
	createApplicationModel.SetName(applicationName)
	createApplicationModel.SetType(applicationType)
	createApplicationModel.SetPermissions(applicationPermissions)

	createdApplication, response, err := apiClient.ApplicationsApi.ApplicationCreate(contextWithAPIKey).CreateApplicationModel(createApplicationModel).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationCreate", t)

	var application *basistheory.ApplicationModel
	application, response, err = apiClient.ApplicationsApi.ApplicationGetById(contextWithAPIKey, createdApplication.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationGetById", t)

	expectedApplicationName := basistheory.NullableString{}
	expectedApplicationName.Set(&applicationName)

	testutils.AssertPropertiesMatchGeneric(application.Name, expectedApplicationName, t, basistheory.NullableString{})
	expectedApplicationType := basistheory.NullableString{}
	expectedApplicationType.Set(&applicationType)
	testutils.AssertPropertiesMatchGeneric(application.Type, expectedApplicationType, t, basistheory.NullableString{})
	testutils.AssertPropertiesMatch(application.Permissions, applicationPermissions, t)

	var applications *basistheory.ApplicationModelPaginatedList
	applications, response, err = apiClient.ApplicationsApi.ApplicationsGet(contextWithAPIKey).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationsGet", t)

	testutils.AssertPropertiesMatchGeneric(applications.Data[0].Name, expectedApplicationName, t, basistheory.NullableString{})
	testutils.AssertPropertiesMatchGeneric(applications.Data[0].Type, expectedApplicationType, t, basistheory.NullableString{})
	testutils.AssertPropertiesMatch(applications.Data[0].Permissions, applicationPermissions, t)

	updatedApplicationName := "Updated Name"
	updatedApplicationNameNullableString := basistheory.NullableString{}
	updatedApplicationNameNullableString.Set(&updatedApplicationName)
	updatedApplicationPermissions := []string{"token:general:read:low"}
	var updatedApplication *basistheory.ApplicationModel
	updatedApplication, response, err = apiClient.ApplicationsApi.ApplicationUpdate(contextWithAPIKey, createdApplication.GetId()).UpdateApplicationModel(basistheory.UpdateApplicationModel{
		Name:        updatedApplicationNameNullableString,
		Permissions: updatedApplicationPermissions,
	}).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationUpdate", t)
	testutils.AssertPropertiesMatchGeneric(updatedApplication.Name, updatedApplicationNameNullableString, t, basistheory.NullableString{})
	testutils.AssertPropertiesMatch(updatedApplication.Permissions, updatedApplicationPermissions, t)
	response, err = apiClient.ApplicationsApi.ApplicationDelete(contextWithAPIKey, createdApplication.GetId()).Execute()
	testutils.AssertMethodDidNotError(err, response, "ApplicationDelete", t)
	var deletedApplication *basistheory.ApplicationModel
	deletedApplication, _, _ = apiClient.ApplicationsApi.ApplicationGetById(contextWithAPIKey, createdApplication.GetId()).Execute()
	testutils.AssertDeletion[*basistheory.ApplicationModel](deletedApplication, nil, "ApplicationDelete", t)
}

func TestApplicationRegenerate(t *testing.T) {
	apiClient, contextWithAPIKey := testutils.CreateLocalAPIAndContext()

	applicationName := _applicationName
	applicationType := _applicationType
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

	testutils.AssertPropertiesDoNotMatchGeneric(regeneratedApplication.Key, createdApplication.Key, t, basistheory.NullableString{})
}

func TestApplicationKey(t *testing.T) {
	apiClient, contextWithAPIKey := testutils.CreateLocalAPIAndContext()

	applicationName := _applicationName
	applicationType := _applicationType
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
	testutils.AssertPropertiesMatchGeneric(applicationFromApplicationKey, createdApplication, t, basistheory.NullableString{}, basistheory.NullableTime{})
}

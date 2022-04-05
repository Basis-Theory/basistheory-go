package basistheory

import (
	"context"
	"github.com/Basis-Theory/basistheory-go/internal/testutils"
	"testing"
)

const _applicationName = "Go Test App"
const _applicationType = "server_to_server"

func createLocalAPIAndContext() (*APIClient, context.Context) {
	configuration := NewConfiguration()
	configuration.Scheme = "http"
	configuration.Host = "localhost:5090"

	apiClient := NewAPIClient(configuration)
	contextWithAPIKey := context.WithValue(context.Background(), ContextAPIKeys, map[string]APIKey{
		"ApiKey": {Key: "key_acc_testkey"},
	})

	return apiClient, contextWithAPIKey
}

func TestApplicationCRUD(t *testing.T) {
	apiClient, contextWithAPIKey := createLocalAPIAndContext()

	applicationName := _applicationName
	applicationType := _applicationType
	applicationPermissions := []string{"token:pci:create"}
	createApplicationModel := *NewCreateApplicationModel()
	createApplicationModel.SetName(applicationName)
	createApplicationModel.SetType(applicationType)
	createApplicationModel.SetPermissions(applicationPermissions)

	createdApplication, response, err := apiClient.ApplicationsApi.ApplicationCreate(contextWithAPIKey).CreateApplicationModel(createApplicationModel).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationCreate", t)

	var application *ApplicationModel
	application, response, err = apiClient.ApplicationsApi.ApplicationGetById(contextWithAPIKey, createdApplication.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationGetById", t)

	testutils.AssertPropertiesMatchGeneric(application.Name, NullableString{
		value: &applicationName,
		isSet: true,
	}, t, NullableString{})
	testutils.AssertPropertiesMatchGeneric(application.Type, NullableString{
		value: &applicationType,
		isSet: true,
	}, t, NullableString{})
	testutils.AssertPropertiesMatch(application.Permissions, applicationPermissions, t)

	var applications *ApplicationModelPaginatedList
	applications, response, err = apiClient.ApplicationsApi.ApplicationsGet(contextWithAPIKey).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationsGet", t)

	testutils.AssertPropertiesMatchGeneric(applications.Data[0].Name, NullableString{
		value: &applicationName,
		isSet: true,
	}, t, NullableString{})
	testutils.AssertPropertiesMatchGeneric(applications.Data[0].Type, NullableString{
		value: &applicationType,
		isSet: true,
	}, t, NullableString{})
	testutils.AssertPropertiesMatch(applications.Data[0].Permissions, applicationPermissions, t)

	updatedApplicationName := "Updated Name"
	updatedApplicationPermissions := []string{"token:general:read:low"}
	var updatedApplication *ApplicationModel
	updatedApplication, response, err = apiClient.ApplicationsApi.ApplicationUpdate(contextWithAPIKey, createdApplication.GetId()).UpdateApplicationModel(UpdateApplicationModel{
		Name: NullableString{
			value: &updatedApplicationName,
			isSet: true,
		},
		Permissions: updatedApplicationPermissions,
	}).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationUpdate", t)
	testutils.AssertPropertiesMatchGeneric(updatedApplication.Name, NullableString{
		value: &updatedApplicationName,
		isSet: true,
	}, t, NullableString{})
	testutils.AssertPropertiesMatch(updatedApplication.Permissions, updatedApplicationPermissions, t)
	response, err = apiClient.ApplicationsApi.ApplicationDelete(contextWithAPIKey, createdApplication.GetId()).Execute()
	testutils.AssertMethodDidNotError(err, response, "ApplicationDelete", t)
	var deletedApplication *ApplicationModel
	deletedApplication, _, _ = apiClient.ApplicationsApi.ApplicationGetById(contextWithAPIKey, createdApplication.GetId()).Execute()
	testutils.AssertDeletion[*ApplicationModel](deletedApplication, nil, "ApplicationDelete", t)
}

func TestApplicationRegenerate(t *testing.T) {
	apiClient, contextWithAPIKey := createLocalAPIAndContext()

	applicationName := _applicationName
	applicationType := _applicationType
	applicationPermissions := []string{"token:pci:create"}
	createApplicationModel := *NewCreateApplicationModel()
	createApplicationModel.SetName(applicationName)
	createApplicationModel.SetType(applicationType)
	createApplicationModel.SetPermissions(applicationPermissions)

	createdApplication, response, err := apiClient.ApplicationsApi.ApplicationCreate(contextWithAPIKey).CreateApplicationModel(createApplicationModel).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationCreate", t)

	var regeneratedApplication *ApplicationModel
	regeneratedApplication, response, err = apiClient.ApplicationsApi.ApplicationRegenerate(contextWithAPIKey, createdApplication.GetId()).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationRegenerate", t)

	testutils.AssertPropertiesDoNotMatchGeneric(regeneratedApplication.Key, createdApplication.Key, t, NullableString{})
}

func TestApplicationKey(t *testing.T) {
	apiClient, contextWithAPIKey := createLocalAPIAndContext()

	applicationName := _applicationName
	applicationType := _applicationType
	applicationPermissions := []string{"token:pci:create"}
	createApplicationModel := *NewCreateApplicationModel()
	createApplicationModel.SetName(applicationName)
	createApplicationModel.SetType(applicationType)
	createApplicationModel.SetPermissions(applicationPermissions)

	createdApplication, response, err := apiClient.ApplicationsApi.ApplicationCreate(contextWithAPIKey).CreateApplicationModel(createApplicationModel).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationCreate", t)

	contextWithCreatedAppAPIKey := context.WithValue(context.Background(), ContextAPIKeys, map[string]APIKey{
		"ApiKey": {Key: *createdApplication.Key.value},
	})

	var applicationFromApplicationKey *ApplicationModel
	applicationFromApplicationKey, response, err = apiClient.ApplicationsApi.ApplicationKey(contextWithCreatedAppAPIKey).Execute()

	testutils.AssertMethodDidNotError(err, response, "ApplicationKey", t)

	applicationFromApplicationKey.SetKey(*createdApplication.Key.value)
	testutils.AssertPropertiesMatchGeneric(applicationFromApplicationKey, createdApplication, t, NullableString{}, NullableTime{})
}

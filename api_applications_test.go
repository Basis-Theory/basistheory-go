package basistheory

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"net/http"
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

func assertMethodDidNotError(err error, response *http.Response, methodName string, t *testing.T) {
	if err != nil {
		t.Errorf("%s response: %v", methodName, response)
		t.Fatalf("%s failed: %v", methodName, err)
	}
}

func assertPropertiesMatch(actualProperty interface{}, expectedProperty interface{}, t *testing.T) {
	if !cmp.Equal(actualProperty, expectedProperty) {
		t.Errorf("does not match expected: got: %+v, want: %+v", spew.Sdump(actualProperty), spew.Sdump(expectedProperty))
	}
}

func assertPropertiesMatchGeneric(actualProperty interface{}, expectedProperty interface{}, t *testing.T, comparedType ...interface{}) {
	if !cmp.Equal(actualProperty, expectedProperty, cmp.AllowUnexported(comparedType...)) {
		t.Errorf("does not match expected: got: %+v, want: %+v", spew.Sdump(actualProperty), spew.Sdump(expectedProperty))
	}
}

func assertPropertiesDoNotMatchGeneric(actualProperty interface{}, expectedProperty interface{}, t *testing.T, comparedType ...interface{}) {
	if cmp.Equal(actualProperty, expectedProperty, cmp.AllowUnexported(comparedType...)) {
		t.Errorf("matches expected: got: %+v, want: %+v", spew.Sdump(actualProperty), spew.Sdump(expectedProperty))
	}
}

func assertDeletion[T any](actualDeletedResource T, expectedDeletedResource T, methodName string, t *testing.T) {
	if !cmp.Equal(actualDeletedResource, expectedDeletedResource) {
		t.Fatalf("%s failed to delete", methodName)
	}
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

	assertMethodDidNotError(err, response, "ApplicationCreate", t)

	var application *ApplicationModel
	application, response, err = apiClient.ApplicationsApi.ApplicationGetById(contextWithAPIKey, createdApplication.GetId()).Execute()

	assertMethodDidNotError(err, response, "ApplicationGetById", t)

	assertPropertiesMatchGeneric(application.Name, NullableString{
		value: &applicationName,
		isSet: true,
	}, t, NullableString{})
	assertPropertiesMatchGeneric(application.Type, NullableString{
		value: &applicationType,
		isSet: true,
	}, t, NullableString{})
	assertPropertiesMatch(application.Permissions, applicationPermissions, t)

	var applications *ApplicationModelPaginatedList
	applications, response, err = apiClient.ApplicationsApi.ApplicationsGet(contextWithAPIKey).Execute()

	assertMethodDidNotError(err, response, "ApplicationsGet", t)

	assertPropertiesMatchGeneric(applications.Data[0].Name, NullableString{
		value: &applicationName,
		isSet: true,
	}, t, NullableString{})
	assertPropertiesMatchGeneric(applications.Data[0].Type, NullableString{
		value: &applicationType,
		isSet: true,
	}, t, NullableString{})
	assertPropertiesMatch(applications.Data[0].Permissions, applicationPermissions, t)

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

	assertMethodDidNotError(err, response, "ApplicationUpdate", t)
	assertPropertiesMatchGeneric(updatedApplication.Name, NullableString{
		value: &updatedApplicationName,
		isSet: true,
	}, t, NullableString{})
	assertPropertiesMatch(updatedApplication.Permissions, updatedApplicationPermissions, t)
	response, err = apiClient.ApplicationsApi.ApplicationDelete(contextWithAPIKey, createdApplication.GetId()).Execute()
	assertMethodDidNotError(err, response, "ApplicationDelete", t)
	var deletedApplication *ApplicationModel
	deletedApplication, _, _ = apiClient.ApplicationsApi.ApplicationGetById(contextWithAPIKey, createdApplication.GetId()).Execute()
	assertDeletion[*ApplicationModel](deletedApplication, nil, "ApplicationDelete", t)
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

	assertMethodDidNotError(err, response, "ApplicationCreate", t)

	var regeneratedApplication *ApplicationModel
	regeneratedApplication, response, err = apiClient.ApplicationsApi.ApplicationRegenerate(contextWithAPIKey, createdApplication.GetId()).Execute()

	assertMethodDidNotError(err, response, "ApplicationRegenerate", t)

	assertPropertiesDoNotMatchGeneric(regeneratedApplication.Key, createdApplication.Key, t, NullableString{})
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

	assertMethodDidNotError(err, response, "ApplicationCreate", t)

	contextWithCreatedAppAPIKey := context.WithValue(context.Background(), ContextAPIKeys, map[string]APIKey{
		"ApiKey": {Key: *createdApplication.Key.value},
	})

	var applicationFromApplicationKey *ApplicationModel
	applicationFromApplicationKey, response, err = apiClient.ApplicationsApi.ApplicationKey(contextWithCreatedAppAPIKey).Execute()

	assertMethodDidNotError(err, response, "ApplicationKey", t)

	applicationFromApplicationKey.SetKey(*createdApplication.Key.value)
	assertPropertiesMatchGeneric(applicationFromApplicationKey, createdApplication, t, NullableString{}, NullableTime{})
}

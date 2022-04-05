package basistheory

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"testing"
)

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

func assertPropertiesMatchGeneric[T any](actualProperty interface{}, expectedProperty interface{}, t *testing.T, comparedType T) {
	if !cmp.Equal(actualProperty, expectedProperty, cmp.AllowUnexported(comparedType)) {
		t.Errorf("does not match expected: got: %+v, want: %+v", spew.Sdump(actualProperty), spew.Sdump(expectedProperty))
	}
}

func TestApplicationCRUD(t *testing.T) {
	apiClient, contextWithAPIKey := createLocalAPIAndContext()

	applicationName := "Go Test App"
	applicationType := "server_to_server"
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

	assertPropertiesMatchGeneric[NullableString](application.Name, NullableString{
		value: &applicationName,
		isSet: true,
	}, t, NullableString{})
	assertPropertiesMatchGeneric[NullableString](application.Type, NullableString{
		value: &applicationType,
		isSet: true,
	}, t, NullableString{})
	assertPropertiesMatch(application.Permissions, applicationPermissions, t)

	// _, response, err = apiClient.ApplicationsApi.ApplicationsGet(contextWithAPIKey).Execute()
}

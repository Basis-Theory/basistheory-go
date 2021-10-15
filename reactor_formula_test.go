package basistheory_go

import (
	"encoding/json"
	"fmt"
	"github.com/Basis-Theory/basistheory-go/internal/test_utils"
	"github.com/google/go-cmp/cmp"
	"github.com/jarcoal/httpmock"
	"github.com/jaswdr/faker"
	"testing"
)

func getBtClientAndBaseUrl() (*BasisTheoryClient, string) {
	fakerInst := faker.New()
	baseUrl := fmt.Sprintf("http://%s", fakerInst.Lorem().Word())
	apiKey := fakerInst.Lorem().Word()

	btClient, _ := NewBasisTheoryClient(baseUrl, apiKey)
	httpmock.ActivateNonDefault(btClient.httpClient.GetClient())
	httpmock.Reset()

	return btClient, baseUrl
}

func TestGetReactorFormulaReturnsAReactorFormulaWhenItGetsA200(t *testing.T) {
	defer httpmock.DeactivateAndReset()

	fakerInst := faker.New()
	expectedReactorFormula := ReactorFormula{}
	fakerInst.Struct().Fill(&expectedReactorFormula)
	expectedReactorFormulaId := fakerInst.UUID().V4()
	btClient, baseUrl := getBtClientAndBaseUrl()
	path := fmt.Sprintf("%s/reactor-formulas/%s", baseUrl, expectedReactorFormulaId)

	responder, _ := httpmock.NewJsonResponder(200, expectedReactorFormula)
	httpmock.RegisterResponder("GET", path, responder)

	actualReactorFormula, err := btClient.GetReactorFormula(expectedReactorFormulaId)

	if err != nil {
		t.Fatalf("GetReactorFormula failed: %v", err)
	}

	if !cmp.Equal(*actualReactorFormula, expectedReactorFormula) {
		t.Errorf(
			"actualReactorFormula does not match expectedReactorFormula, got: %s, want: %s",
			test_utils.PrettyPrint(*actualReactorFormula), test_utils.PrettyPrint(expectedReactorFormula))
	}
}

func TestGetReactorFormulaReturnsDetailsInErrorWhenItGetsAnErrorHttpStatus(t *testing.T) {
	defer httpmock.DeactivateAndReset()

	fakerInst := faker.New()
	expectedReactorFormulaId := fakerInst.UUID().V4()
	btClient, baseUrl := getBtClientAndBaseUrl()
	path := fmt.Sprintf("%s/reactor-formulas/%s", baseUrl, expectedReactorFormulaId)

	expectedErrorPayload, expectedStatus := test_utils.SetupErrorPath("GET", path)

	_, err := btClient.GetReactorFormula(expectedReactorFormulaId)

	marshalledExpectedError, _ := json.Marshal(expectedErrorPayload)
	actualError := *err.(*ApiError)
	expectedError := ApiError{
		Code: expectedStatus,
		Message: fmt.Sprintf(
			"Sending GET request to %s: %d. Response body: %v",
			path, expectedStatus, string(marshalledExpectedError)),
	}

	if !cmp.Equal(actualError, expectedError) {
		t.Errorf(
			"actualError does not match expectedError, got: %s, want: %s",
			test_utils.PrettyPrint(actualError), test_utils.PrettyPrint(expectedError))
	}
}

func TestGetReactorFormulasWithQueryReturnsPaginatedReactorFormulasWhenItGetsA200(t *testing.T) {
	defer httpmock.DeactivateAndReset()

	fakerInst := faker.New()
	expectedPaginatedReactorFormulas := PaginatedReactorFormulas{}
	fakerInst.Struct().Fill(&expectedPaginatedReactorFormulas)
	expectedReactorFormulaQuery := ReactorFormulaQuery{}
	fakerInst.Struct().Fill(&expectedReactorFormulaQuery)
	btClient, baseUrl := getBtClientAndBaseUrl()

	responder, _ := httpmock.NewJsonResponder(200, expectedPaginatedReactorFormulas)
	params := map[string]string{
		"name":              expectedReactorFormulaQuery.Name,
		"source_token_type": expectedReactorFormulaQuery.SourceTokenType,
		"page":              expectedReactorFormulaQuery.Page,
		"size":              expectedReactorFormulaQuery.Size,
	}
	path := fmt.Sprintf("%s/reactor-formulas?name=%s&page=%s&size=%s&source_token_type=%s",
		baseUrl, params["name"], params["page"], params["size"], params["source_token_type"])
	httpmock.RegisterResponder("GET", path, responder)

	actualPaginatedReactorFormulas, err := btClient.GetReactorFormulasWithQuery(expectedReactorFormulaQuery)

	if err != nil {
		t.Fatalf("GetReactorFormulasWithQuery failed: %v", err)
	}

	if !cmp.Equal(*actualPaginatedReactorFormulas, expectedPaginatedReactorFormulas) {
		t.Errorf(
			"actualPaginatedReactorFormulas does not match expectedPaginatedReactorFormulas, got: %s, want: %s",
			test_utils.PrettyPrint(*actualPaginatedReactorFormulas), test_utils.PrettyPrint(expectedPaginatedReactorFormulas))
	}
}

func TestGetReactorFormulasWithQueryReturnsDetailsInErrorWhenItGetsAnErrorHttpStatus(t *testing.T) {
	defer httpmock.DeactivateAndReset()

	fakerInst := faker.New()
	expectedPaginatedReactorFormulas := PaginatedReactorFormulas{}
	fakerInst.Struct().Fill(&expectedPaginatedReactorFormulas)
	expectedReactorFormulaQuery := ReactorFormulaQuery{}
	fakerInst.Struct().Fill(&expectedReactorFormulaQuery)
	btClient, baseUrl := getBtClientAndBaseUrl()
	params := map[string]string{
		"name":              expectedReactorFormulaQuery.Name,
		"source_token_type": expectedReactorFormulaQuery.SourceTokenType,
		"page":              expectedReactorFormulaQuery.Page,
		"size":              expectedReactorFormulaQuery.Size,
	}
	path := fmt.Sprintf("%s/reactor-formulas?name=%s&page=%s&size=%s&source_token_type=%s",
		baseUrl, params["name"], params["page"], params["size"], params["source_token_type"])

	expectedErrorPayload, expectedStatus := test_utils.SetupErrorPath("GET", path)

	_, err := btClient.GetReactorFormulasWithQuery(expectedReactorFormulaQuery)
	marshalledExpectedError, _ := json.Marshal(expectedErrorPayload)
	actualError := *err.(*ApiError)
	expectedError := ApiError{
		Code: expectedStatus,
		Message: fmt.Sprintf(
			"Sending GET request to %s: %d. Response body: %v",
			path, expectedStatus, string(marshalledExpectedError)),
	}

	if !cmp.Equal(actualError, expectedError) {
		t.Errorf(
			"actualError does not match expectedError, got: %s, want: %s",
			test_utils.PrettyPrint(actualError), test_utils.PrettyPrint(expectedError))
	}
}

func TestGetReactorFormulasReturnsPaginatedReactorFormulasWhenItGetsA200(t *testing.T) {
	defer httpmock.DeactivateAndReset()

	fakerInst := faker.New()
	expectedPaginatedReactorFormulas := PaginatedReactorFormulas{}
	fakerInst.Struct().Fill(&expectedPaginatedReactorFormulas)
	btClient, baseUrl := getBtClientAndBaseUrl()
	path := fmt.Sprintf("%s/reactor-formulas", baseUrl)

	responder, _ := httpmock.NewJsonResponder(200, expectedPaginatedReactorFormulas)
	httpmock.RegisterResponder("GET", path, responder)

	actualPaginatedReactorFormulas, err := btClient.GetReactorFormulas()

	if err != nil {
		t.Fatalf("GetReactorFormulas failed: %v", err)
	}

	if !cmp.Equal(*actualPaginatedReactorFormulas, expectedPaginatedReactorFormulas) {
		t.Errorf(
			"actualPaginatedReactorFormulas does not match expectedPaginatedReactorFormulas, got: %s, want: %s",
			test_utils.PrettyPrint(*actualPaginatedReactorFormulas), test_utils.PrettyPrint(expectedPaginatedReactorFormulas))
	}
}

func TestGetReactorFormulasReturnsDetailsInErrorWhenItGetsAnErrorHttpStatus(t *testing.T) {
	defer httpmock.DeactivateAndReset()

	fakerInst := faker.New()
	expectedPaginatedReactorFormulas := PaginatedReactorFormulas{}
	fakerInst.Struct().Fill(&expectedPaginatedReactorFormulas)
	btClient, baseUrl := getBtClientAndBaseUrl()
	path := fmt.Sprintf("%s/reactor-formulas", baseUrl)

	expectedErrorPayload, expectedStatus := test_utils.SetupErrorPath("GET", path)

	_, err := btClient.GetReactorFormulas()

	marshalledExpectedError, _ := json.Marshal(expectedErrorPayload)
	actualError := *err.(*ApiError)
	expectedError := ApiError{
		Code: expectedStatus,
		Message: fmt.Sprintf(
			"Sending GET request to %s: %d. Response body: %v",
			path, expectedStatus, string(marshalledExpectedError)),
	}

	if !cmp.Equal(actualError, expectedError) {
		t.Errorf(
			"actualError does not match expectedError, got: %s, want: %s",
			test_utils.PrettyPrint(actualError), test_utils.PrettyPrint(expectedError))
	}
}

func TestCreateReactorFormulaReturnsCreatedReactorFormulaWhenItGetsA201(t *testing.T) {
	defer httpmock.DeactivateAndReset()

	fakerInst := faker.New()
	expectedCreateReactorFormula := MutableReactorFormula{}
	fakerInst.Struct().Fill(&expectedCreateReactorFormula)
	expectedCreatedReactorFormula := ReactorFormula{}
	fakerInst.Struct().Fill(&expectedCreatedReactorFormula)
	btClient, baseUrl := getBtClientAndBaseUrl()

	responder, _ := httpmock.NewJsonResponder(201, expectedCreatedReactorFormula)
	path := fmt.Sprintf("%s/reactor-formulas", baseUrl)
	httpmock.RegisterResponder("POST", path, responder)

	actualCreatedReactorFormula, err := btClient.CreateReactorFormula(expectedCreateReactorFormula)

	if err != nil {
		t.Fatalf("CreateReactorFormula failed: %v", err)
	}

	if !cmp.Equal(*actualCreatedReactorFormula, expectedCreatedReactorFormula) {
		t.Errorf(
			"actualCreatedReactorFormula does not match expectedCreatedReactorFormula, got: %s, want: %s",
			test_utils.PrettyPrint(*actualCreatedReactorFormula), test_utils.PrettyPrint(expectedCreatedReactorFormula))
	}
}

func TestCreateReactorFormulaReturnsDetailsInErrorWhenItGetsAnErrorHttpStatus(t *testing.T) {
	defer httpmock.DeactivateAndReset()

	fakerInst := faker.New()
	expectedCreateReactorFormula := MutableReactorFormula{}
	fakerInst.Struct().Fill(&expectedCreateReactorFormula)
	btClient, baseUrl := getBtClientAndBaseUrl()
	path := fmt.Sprintf("%s/reactor-formulas", baseUrl)

	expectedErrorPayload, expectedStatus := test_utils.SetupErrorPath("POST", path)

	_, err := btClient.CreateReactorFormula(expectedCreateReactorFormula)

	marshalledExpectedError, _ := json.Marshal(expectedErrorPayload)
	actualError := *err.(*ApiError)
	expectedError := ApiError{
		Code: expectedStatus,
		Message: fmt.Sprintf(
			"Sending POST request to %s: %d. Response body: %v",
			path, expectedStatus, string(marshalledExpectedError)),
	}

	if !cmp.Equal(actualError, expectedError) {
		t.Errorf(
			"actualError does not match expectedError, got: %s, want: %s",
			test_utils.PrettyPrint(actualError), test_utils.PrettyPrint(expectedError))
	}
}

func TestDeleteReactorFormulaDoesNotReturnErrorWhenItGetsA204(t *testing.T) {
	defer httpmock.DeactivateAndReset()

	fakerInst := faker.New()
	expectedReactorFormulaId := fakerInst.Lorem().Word()
	btClient, baseUrl := getBtClientAndBaseUrl()
	path := fmt.Sprintf("%s/reactor-formulas/%s", baseUrl, expectedReactorFormulaId)

	responder, _ := httpmock.NewJsonResponder(204, nil)
	httpmock.RegisterResponder("DELETE", path, responder)

	err := btClient.DeleteReactorFormula(expectedReactorFormulaId)

	if err != nil {
		t.Fatalf("DeleteReactorFormula failed: %v", err)
	}
}

func TestDeleteReactorFormulaReturnsDetailsInErrorWhenItGetsAnErrorHttpStatus(t *testing.T) {
	defer httpmock.DeactivateAndReset()

	fakerInst := faker.New()
	expectedReactorFormulaId := fakerInst.Lorem().Word()
	btClient, baseUrl := getBtClientAndBaseUrl()
	path := fmt.Sprintf("%s/reactor-formulas/%s", baseUrl, expectedReactorFormulaId)

	expectedErrorPayload, expectedStatus := test_utils.SetupErrorPath("DELETE", path)

	err := btClient.DeleteReactorFormula(expectedReactorFormulaId)

	marshalledExpectedError, _ := json.Marshal(expectedErrorPayload)
	actualError := *err.(*ApiError)
	expectedError := ApiError{
		Code: expectedStatus,
		Message: fmt.Sprintf(
			"Sending DELETE request to %s: %d. Response body: %v",
			path, expectedStatus, string(marshalledExpectedError)),
	}

	if !cmp.Equal(actualError, expectedError) {
		t.Errorf(
			"actualError does not match expectedError, got: %s, want: %s",
			test_utils.PrettyPrint(actualError), test_utils.PrettyPrint(expectedError))
	}
}

func TestUpdateReactorFormulaReturnsUpdatedReactorFormulaWhenItGetsA200(t *testing.T) {
	defer httpmock.DeactivateAndReset()

	fakerInst := faker.New()
	expectedReactorFormulaId := fakerInst.Lorem().Word()
	expectedUpdateReactorFormula := MutableReactorFormula{}
	fakerInst.Struct().Fill(&expectedUpdateReactorFormula)
	expectedUpdatedReactorFormula := ReactorFormula{}
	fakerInst.Struct().Fill(&expectedUpdatedReactorFormula)
	btClient, baseUrl := getBtClientAndBaseUrl()
	path := fmt.Sprintf("%s/reactor-formulas/%s", baseUrl, expectedReactorFormulaId)

	responder, _ := httpmock.NewJsonResponder(200, expectedUpdatedReactorFormula)
	httpmock.RegisterResponder("PUT", path, responder)

	actualUpdatedReactorFormula, err := btClient.UpdateReactorFormula(expectedReactorFormulaId, expectedUpdateReactorFormula)

	if err != nil {
		t.Fatalf("UpdateReactorFormula failed: %v", err)
	}

	if !cmp.Equal(*actualUpdatedReactorFormula, expectedUpdatedReactorFormula) {
		t.Errorf(
			"actualUpdatedReactorFormula does not match expectedUpdatedReactorFormula, got: %s, want: %s",
			test_utils.PrettyPrint(*actualUpdatedReactorFormula), test_utils.PrettyPrint(expectedUpdatedReactorFormula))
	}
}

func TestUpdateReactorFormulaReturnsDetailsInErrorWhenItGetsAnErrorHttpStatus(t *testing.T) {
	defer httpmock.DeactivateAndReset()

	fakerInst := faker.New()
	expectedReactorFormulaId := fakerInst.Lorem().Word()
	expectedUpdateReactorFormula := MutableReactorFormula{}
	fakerInst.Struct().Fill(&expectedUpdateReactorFormula)
	expectedUpdatedReactorFormula := ReactorFormula{}
	fakerInst.Struct().Fill(&expectedUpdatedReactorFormula)
	btClient, baseUrl := getBtClientAndBaseUrl()
	path := fmt.Sprintf("%s/reactor-formulas/%s", baseUrl, expectedReactorFormulaId)

	expectedErrorPayload, expectedStatus := test_utils.SetupErrorPath("PUT", path)

	_, err := btClient.UpdateReactorFormula(expectedReactorFormulaId, expectedUpdateReactorFormula)

	marshalledExpectedError, _ := json.Marshal(expectedErrorPayload)
	actualError := *err.(*ApiError)
	expectedError := ApiError{
		Code: expectedStatus,
		Message: fmt.Sprintf(
			"Sending PUT request to %s: %d. Response body: %v",
			path, expectedStatus, string(marshalledExpectedError)),
	}

	if !cmp.Equal(actualError, expectedError) {
		t.Errorf(
			"actualError does not match expectedError, got: %s, want: %s",
			test_utils.PrettyPrint(actualError), test_utils.PrettyPrint(expectedError))
	}
}

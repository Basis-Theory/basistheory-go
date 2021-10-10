package basistheory_go

import (
	"encoding/json"
	"fmt"
	"github.com/Basis-Theory/basistheory-go/internal/test_utils"
	"github.com/jarcoal/httpmock"
	"github.com/jaswdr/faker"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReactorFormula client", func() {
	var (
		fakerInst faker.Faker
		baseUrl   string
		apiKey    string
		btClient  *BasisTheoryClient
	)

	BeforeEach(func() {
		fakerInst = faker.New()
		baseUrl = fmt.Sprintf("http://%s", fakerInst.Lorem().Word())
		apiKey = fakerInst.Lorem().Word()

		btClient, _ = NewBasisTheoryClient(baseUrl, apiKey)
		httpmock.ActivateNonDefault(btClient.httpClient.GetClient())
		httpmock.Reset()
	})

	Context("GetReactorFormula", func() {
		var (
			expectedReactorFormulaId string
			expectedReactorFormula   ReactorFormula
		)

		BeforeEach(func() {
			expectedReactorFormulaId = fakerInst.UUID().V4()
			expectedReactorFormula = ReactorFormula{}
			fakerInst.Struct().Fill(&expectedReactorFormula)
		})

		Context("request to /reactor-formulas/id is successful", func() {
			It("returns a ReactorFormula", func() {
				responder, _ := httpmock.NewJsonResponder(200, expectedReactorFormula)
				path := fmt.Sprintf("%s/reactor-formulas/%s", baseUrl, expectedReactorFormulaId)
				httpmock.RegisterResponder("GET", path, responder)

				actualReactorFormula, err := btClient.GetReactorFormula(expectedReactorFormulaId)

				if err != nil {
					Fail(fmt.Sprintf("GetReactorFormula failed: %v", err))
				}

				Expect(*actualReactorFormula).To(Equal(expectedReactorFormula))
			})
		})

		Context("request to /reactor-formulas/id is not successful", func() {
			It("returns an error containing details of the failed request", func() {
				path := fmt.Sprintf("%s/reactor-formulas/%s", baseUrl, expectedReactorFormulaId)
				expectedErrorPayload, expectedStatus := test_utils.SetupErrorPath("GET", path)

				_, err := btClient.GetReactorFormula(expectedReactorFormulaId)

				marshalledExpectedError, _ := json.Marshal(expectedErrorPayload)

				Expect(*err.(*ApiError)).To(Equal(ApiError{
					Code: expectedStatus,
					Message: fmt.Sprintf(
						"Sending GET request to %s: %d. Response body: %v",
						path, expectedStatus, string(marshalledExpectedError)),
				}))
			})
		})
	})

	Context("GetReactorFormulas", func() {
		var (
			expectedPaginatedReactorFormulas PaginatedReactorFormulas
		)

		BeforeEach(func() {
			expectedPaginatedReactorFormulas = PaginatedReactorFormulas{}
			fakerInst.Struct().Fill(&expectedPaginatedReactorFormulas)
		})

		Context("query params are defined", func() {
			var (
				expectedReactorFormulaQuery ReactorFormulaQuery
			)

			BeforeEach(func() {
				expectedReactorFormulaQuery = ReactorFormulaQuery{}
				fakerInst.Struct().Fill(&expectedReactorFormulaQuery)
			})

			Context("request to /reactor-formulas with query params is successful", func() {
				It("returns PaginatedReactorFormulas", func() {
					responder, _ := httpmock.NewJsonResponder(200, expectedPaginatedReactorFormulas)
					params := map[string]string {
						"name": expectedReactorFormulaQuery.Name,
						"source_token_type": expectedReactorFormulaQuery.SourceTokenType,
						"page": expectedReactorFormulaQuery.Page,
						"size": expectedReactorFormulaQuery.Size,
					}
					path := fmt.Sprintf("%s/reactor-formulas?name=%s&page=%s&size=%s&source_token_type=%s",
						baseUrl, params["name"], params["page"], params["size"], params["source_token_type"])
					httpmock.RegisterResponder("GET", path, responder)

					actualReactorFormulas, err := btClient.GetReactorFormulasWithQuery(expectedReactorFormulaQuery)

					if err != nil {
						Fail(fmt.Sprintf("GetReactorFormulasWithQuery failed: %v", err))
					}

					Expect(*actualReactorFormulas).To(Equal(expectedPaginatedReactorFormulas))
				})
			})

			Context("request to /reactor-formulas with query params is not successful", func() {
				It("returns an error containing details of the failed request", func() {
					params := map[string]string {
						"name": expectedReactorFormulaQuery.Name,
						"source_token_type": expectedReactorFormulaQuery.SourceTokenType,
						"page": expectedReactorFormulaQuery.Page,
						"size": expectedReactorFormulaQuery.Size,
					}
					path := fmt.Sprintf("%s/reactor-formulas?name=%s&page=%s&size=%s&source_token_type=%s",
						baseUrl, params["name"], params["page"], params["size"], params["source_token_type"])
					expectedErrorPayload, expectedStatus := test_utils.SetupErrorPath("GET", path)

					_, err := btClient.GetReactorFormulasWithQuery(expectedReactorFormulaQuery)

					marshalledExpectedError, _ := json.Marshal(expectedErrorPayload)

					Expect(*err.(*ApiError)).To(Equal(ApiError{
						Code: expectedStatus,
						Message: fmt.Sprintf(
							"Sending GET request to %s: %d. Response body: %v",
							path, expectedStatus, string(marshalledExpectedError)),
					}))
				})
			})
		})

		Context("query params are not defined", func() {
			Context("request to /reactor-formulas is successful", func() {
				It("returns PaginatedReactorFormulas", func() {
					responder, _ := httpmock.NewJsonResponder(200, expectedPaginatedReactorFormulas)
					path := fmt.Sprintf("%s/reactor-formulas", baseUrl)
					httpmock.RegisterResponder("GET", path, responder)

					actualReactorFormulas, err := btClient.GetReactorFormulas()

					if err != nil {
						Fail(fmt.Sprintf("GetReactorFormulas failed: %v", err))
					}

					Expect(*actualReactorFormulas).To(Equal(expectedPaginatedReactorFormulas))
				})
			})

			Context("request to /reactor-formulas is not successful", func() {
				It("returns an error containing details of the failed request", func() {
					path := fmt.Sprintf("%s/reactor-formulas", baseUrl)
					expectedErrorPayload, expectedStatus := test_utils.SetupErrorPath("GET", path)

					_, err := btClient.GetReactorFormulas()

					marshalledExpectedError, _ := json.Marshal(expectedErrorPayload)

					Expect(*err.(*ApiError)).To(Equal(ApiError{
						Code: expectedStatus,
						Message: fmt.Sprintf(
							"Sending GET request to %s: %d. Response body: %v",
							path, expectedStatus, string(marshalledExpectedError)),
					}))
				})
			})
		})
	})
})

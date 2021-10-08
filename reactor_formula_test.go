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
		fakerInst                faker.Faker
		baseUrl                  string
		apiKey                   string
		expectedReactorFormulaId string
		expectedReactorFormula   ReactorFormula
		btClient                 *BasisTheoryClient
	)

	BeforeSuite(func() {
		fakerInst = faker.New()
		baseUrl = fmt.Sprintf("http://%s", fakerInst.Lorem().Word())
		apiKey = fakerInst.Lorem().Word()
		expectedReactorFormulaId = fakerInst.UUID().V4()
		expectedReactorFormula = ReactorFormula{}
		fakerInst.Struct().Fill(&expectedReactorFormula)

		btClient, _ = NewBasisTheoryClient(baseUrl, apiKey)

		httpmock.ActivateNonDefault(btClient.httpClient.GetClient())
	})

	BeforeEach(func() {
		httpmock.Reset()
	})

	AfterSuite(func() {
		httpmock.DeactivateAndReset()
	})

	Context("GetReactorFormula", func() {
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
		Context("request to /reactor-formulas is successful", func() {

		})

		Context("request to /reactor-formulas is not successful", func() {

		})
	})
})

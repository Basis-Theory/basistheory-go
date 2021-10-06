package client

import (
	"encoding/json"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/jaswdr/faker"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func setupBtClient(baseUrl string, apiKey string) *BasisTheoryClient {
	btClient, err := NewBasisTheoryClient(baseUrl, apiKey, "reactor_formula_testing", nil, 10)

	if err != nil {
		Fail(fmt.Sprintf("Could not create BasisTheory client %v", err))
	}

	return btClient
}

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

		btClient = setupBtClient(baseUrl, apiKey)

		httpmock.ActivateNonDefault(btClient.httpClient.GetClient())
	})

	BeforeEach(func() {
		httpmock.Reset()
	})

	AfterSuite(func() {
		httpmock.DeactivateAndReset()
	})

	Context("GetReactorFormula", func() {
		Context("request to /reactor-formulas is successful", func() {
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

		Context("request to /reactor-formulas is not successful", func() {
			It("returns an error containing details of the failed request", func() {
				path := fmt.Sprintf("%s/reactor-formulas/%s", baseUrl, expectedReactorFormulaId)
				expectedStatus := fakerInst.IntBetween(400, 500)
				expectedErrorPayload := map[string]string{
					fakerInst.Lorem().Word(): fakerInst.Lorem().Word(),
				}
				responder, _ := httpmock.NewJsonResponder(expectedStatus, expectedErrorPayload)
				httpmock.RegisterResponder("GET", path, responder)

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
})

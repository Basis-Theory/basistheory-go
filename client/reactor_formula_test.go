package client

import (
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
		fakerInst faker.Faker
		baseUrl   string
		apiKey    string
		btClient  *BasisTheoryClient
	)

	BeforeSuite(func() {
		fakerInst = faker.New()
		baseUrl = fmt.Sprintf("http://%s", fakerInst.Lorem().Word())
		apiKey = fakerInst.Lorem().Word()

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
		It("returns a ReactorFormula", func() {
			expectedReactorFormulaId := fakerInst.UUID().V4()

			expectedReactorFormula := ReactorFormula{}
			fakerInst.Struct().Fill(&expectedReactorFormula)

			responder, _ := httpmock.NewJsonResponder(200, expectedReactorFormula)
			httpmock.RegisterResponder("GET", fmt.Sprintf("%s/reactor-formulas/%s", baseUrl, expectedReactorFormulaId), responder)

			actualReactorFormula, err := btClient.GetReactorFormula(fmt.Sprintf("%v", expectedReactorFormulaId))

			if err != nil {
				Fail(fmt.Sprintf("GetReactorFormula failed: %v", err))
			}

			Expect(actualReactorFormula).To(Equal(&expectedReactorFormula))
		})
	})
})

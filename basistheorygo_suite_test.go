package basistheory_go

import (
	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BasistheoryGo Suite")
}

var _ = AfterSuite(func() {
	httpmock.DeactivateAndReset()
})

package testutils

import (
	"context"
	"github.com/Basis-Theory/basistheory-go/v2"
	"github.com/joho/godotenv"
	"os"
	"strings"
	"testing"
)

func getDevConfiguration() *basistheory.Configuration {
	configuration := basistheory.NewConfiguration()

	apiUrl := os.Getenv("BASISTHEORY_API_URL")

	if apiUrl != "" {
		urlArray := strings.Split(apiUrl, "://")
		configuration.Scheme = urlArray[0]
		configuration.Host = urlArray[1]
	} else {
		configuration.Scheme = "https"
		configuration.Host = "api-dev.basistheory.com"
	}

	return configuration
}

func getApiAndContext(t *testing.T, apiKeyName string) (*basistheory.APIClient, context.Context) {
	_ = godotenv.Load("../.env.local")
	apiKey := os.Getenv(apiKeyName)

	if apiKey == "" {
		t.Fatalf("%s must be set before running acceptance tests.", apiKeyName)
	}

	contextWithAPIKey := context.WithValue(context.Background(), basistheory.ContextAPIKeys, map[string]basistheory.APIKey{
		"ApiKey": {Key: apiKey},
	})

	apiClient := basistheory.NewAPIClient(getDevConfiguration())

	return apiClient, contextWithAPIKey
}

func CreateApiAndMgmtContext(t *testing.T) (*basistheory.APIClient, context.Context) {
	return getApiAndContext(t, "BASISTHEORY_MGMT_API_KEY")
}

func CreateApiAndSrvrContext(t *testing.T) (*basistheory.APIClient, context.Context) {
	return getApiAndContext(t, "BASISTHEORY_SRVR_API_KEY")
}

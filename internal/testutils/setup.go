package testutils

import (
	"context"
	"github.com/Basis-Theory/basistheory-go"
)

func CreateLocalAPIAndContext() (*basistheory.APIClient, context.Context) {
	configuration := basistheory.NewConfiguration()
	configuration.Scheme = "http"
	configuration.Host = "localhost:1080"

	apiClient := basistheory.NewAPIClient(configuration)
	contextWithAPIKey := context.WithValue(context.Background(), basistheory.ContextAPIKeys, map[string]basistheory.APIKey{
		"ApiKey": {Key: "key_acc_testkey"},
	})

	return apiClient, contextWithAPIKey
}

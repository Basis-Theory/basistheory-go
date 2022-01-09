package basistheory_go

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"time"
)

type BasisTheoryClient struct {
	baseUrl    string
	httpClient *resty.Client
}

const defaultMaxTimeout = 10

func NewBasisTheoryClient(baseUrl string, apiKey string) (*BasisTheoryClient, error) {
	return NewBasisTheoryClientWithAddtnlHeadersAndTimeout(baseUrl, apiKey, nil, defaultMaxTimeout)
}

func NewBasisTheoryClientWithAddtnlHeaders(baseUrl string, apiKey string, additionalHeaders map[string]string) (*BasisTheoryClient, error) {
	return NewBasisTheoryClientWithAddtnlHeadersAndTimeout(baseUrl, apiKey, additionalHeaders, defaultMaxTimeout)
}

func NewBasisTheoryClientWithAddtnlHeadersAndTimeout(baseUrl string, apiKey string, additionalHeaders map[string]string, clientTimeout int) (*BasisTheoryClient, error) {
	client := resty.New()

	setupClientMiddleWareAndHeaders(client, apiKey, additionalHeaders, clientTimeout)

	basisTheoryClient := BasisTheoryClient{
		baseUrl:    baseUrl,
		httpClient: client,
	}

	return &basisTheoryClient, nil
}

func setupClientMiddleWareAndHeaders(httpClient *resty.Client, apiKey string, additionalHeaders map[string]string, clientTimeout int) {
	headers := map[string]string{
		"BT-API-KEY": apiKey,
		"Accept":    "application/json",
	}
	for additonalHeaderName, additionalHeaderValue := range additionalHeaders {
		headers[additonalHeaderName] = additionalHeaderValue
	}

	httpClient.SetHeaders(headers)

	httpClient.SetTimeout(time.Second * time.Duration(clientTimeout))

	httpClient.OnBeforeRequest(func(_ *resty.Client, req *resty.Request) error {
		log.Debug().Msg(fmt.Sprintf("Sending %s to %s", req.Method, req.URL))

		if req.Body != nil {
			log.Debug().Msg(fmt.Sprintf("Request body: %s", req.Body))
		}

		return nil
	})

	httpClient.OnAfterResponse(func(_ *resty.Client, resp *resty.Response) error {
		const lowestHttpStatusError = 400
		if resp.StatusCode() >= lowestHttpStatusError {
			errorMessage := fmt.Sprintf("Sending %s request to %s: %d.", resp.Request.Method, resp.Request.URL, resp.StatusCode())

			if len(resp.Body()) != 0 {
				errorMessage = fmt.Sprintf("%s Response body: %s", errorMessage, resp.Body())
			}

			log.Error().Msg(errorMessage)

			return &ApiError{
				Code:    resp.StatusCode(),
				Message: errorMessage,
			}
		}

		return nil
	})
}

func (basisTheoryClient *BasisTheoryClient) get(result interface{}, path string, params map[string]string) error {
	resourceUrl := basisTheoryClient.baseUrl + path

	_, err := basisTheoryClient.httpClient.R().
		SetQueryParams(params).
		SetResult(result).
		Get(resourceUrl)

	if err != nil {
		return err
	}

	return nil
}

func (basisTheoryClient *BasisTheoryClient) post(result interface{}, path string, requestBody interface{}) error {
	resourceUrl := basisTheoryClient.baseUrl + path

	_, err := basisTheoryClient.httpClient.R().
		SetBody(requestBody).
		SetResult(result).
		Post(resourceUrl)

	if err != nil {
		return err
	}

	return nil
}

func (basisTheoryClient *BasisTheoryClient) put(result interface{}, path string, requestBody interface{}) error {
	resourceUrl := basisTheoryClient.baseUrl + path

	_, err := basisTheoryClient.httpClient.R().
		SetBody(requestBody).
		SetResult(result).
		Put(resourceUrl)

	if err != nil {
		return err
	}

	return nil
}

func (basisTheoryClient *BasisTheoryClient) delete(path string) error {
	resourceUrl := basisTheoryClient.baseUrl + path

	_, err := basisTheoryClient.httpClient.R().
		Delete(resourceUrl)

	return err
}

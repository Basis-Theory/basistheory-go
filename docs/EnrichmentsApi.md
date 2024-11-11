# \EnrichmentsApi

All URIs are relative to *https://api.basistheory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BankAccountVerify**](EnrichmentsApi.md#BankAccountVerify) | **Post** /enrichments/bank-account-verify | 



## BankAccountVerify

> BankAccountVerify(ctx).BankVerificationRequest(bankVerificationRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bankVerificationRequest := *openapiclient.NewBankVerificationRequest("TokenId_example") // BankVerificationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EnrichmentsApi.BankAccountVerify(context.Background()).BankVerificationRequest(bankVerificationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrichmentsApi.BankAccountVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankAccountVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankVerificationRequest** | [**BankVerificationRequest**](BankVerificationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


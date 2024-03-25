# \ThreeDSApi

All URIs are relative to *https://api.basistheory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Call3dsSessionsIdGet**](ThreeDSApi.md#Call3dsSessionsIdGet) | **Get** /3ds/sessions/{id} | 
[**ThreeDSAuthenticateSession**](ThreeDSApi.md#ThreeDSAuthenticateSession) | **Post** /3ds/sessions/{sessionId}/authenticate | 
[**ThreeDSGetChallengeResult**](ThreeDSApi.md#ThreeDSGetChallengeResult) | **Get** /3ds/sessions/{sessionId}/challenge-result | 



## Call3dsSessionsIdGet

> ThreeDSSession Call3dsSessionsIdGet(ctx, id).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThreeDSApi.Call3dsSessionsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThreeDSApi.Call3dsSessionsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Call3dsSessionsIdGet`: ThreeDSSession
    fmt.Fprintf(os.Stdout, "Response from `ThreeDSApi.Call3dsSessionsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCall3dsSessionsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ThreeDSSession**](ThreeDSSession.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreeDSAuthenticateSession

> ThreeDSAuthentication ThreeDSAuthenticateSession(ctx, sessionId).AuthenticateThreeDSSessionRequest(authenticateThreeDSSessionRequest).Execute()



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
    sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    authenticateThreeDSSessionRequest := *openapiclient.NewAuthenticateThreeDSSessionRequest("AuthenticationCategory_example", "AuthenticationType_example", *openapiclient.NewThreeDSRequestorInfo()) // AuthenticateThreeDSSessionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThreeDSApi.ThreeDSAuthenticateSession(context.Background(), sessionId).AuthenticateThreeDSSessionRequest(authenticateThreeDSSessionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThreeDSApi.ThreeDSAuthenticateSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThreeDSAuthenticateSession`: ThreeDSAuthentication
    fmt.Fprintf(os.Stdout, "Response from `ThreeDSApi.ThreeDSAuthenticateSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiThreeDSAuthenticateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticateThreeDSSessionRequest** | [**AuthenticateThreeDSSessionRequest**](AuthenticateThreeDSSessionRequest.md) |  | 

### Return type

[**ThreeDSAuthentication**](ThreeDSAuthentication.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreeDSGetChallengeResult

> ThreeDSAuthentication ThreeDSGetChallengeResult(ctx, sessionId).Execute()



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
    sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThreeDSApi.ThreeDSGetChallengeResult(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThreeDSApi.ThreeDSGetChallengeResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThreeDSGetChallengeResult`: ThreeDSAuthentication
    fmt.Fprintf(os.Stdout, "Response from `ThreeDSApi.ThreeDSGetChallengeResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiThreeDSGetChallengeResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ThreeDSAuthentication**](ThreeDSAuthentication.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


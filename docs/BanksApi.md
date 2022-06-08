# \BanksApi

All URIs are relative to *https://api.basistheory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AtomicBanksCreate**](BanksApi.md#AtomicBanksCreate) | **Post** /atomic/banks | 
[**AtomicBanksDelete**](BanksApi.md#AtomicBanksDelete) | **Delete** /atomic/banks/{id} | 
[**AtomicBanksGet**](BanksApi.md#AtomicBanksGet) | **Get** /atomic/banks | 
[**AtomicBanksGetById**](BanksApi.md#AtomicBanksGetById) | **Get** /atomic/banks/{id} | 
[**AtomicBanksReact**](BanksApi.md#AtomicBanksReact) | **Post** /atomic/banks/{bankId}/react | 
[**AtomicBanksUpdate**](BanksApi.md#AtomicBanksUpdate) | **Patch** /atomic/banks/{id} | 



## AtomicBanksCreate

> AtomicBank AtomicBanksCreate(ctx).CreateAtomicBankRequest(createAtomicBankRequest).Execute()



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
    createAtomicBankRequest := *openapiclient.NewCreateAtomicBankRequest(*openapiclient.NewBank("RoutingNumber_example", "AccountNumber_example")) // CreateAtomicBankRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BanksApi.AtomicBanksCreate(context.Background()).CreateAtomicBankRequest(createAtomicBankRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BanksApi.AtomicBanksCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AtomicBanksCreate`: AtomicBank
    fmt.Fprintf(os.Stdout, "Response from `BanksApi.AtomicBanksCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAtomicBanksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAtomicBankRequest** | [**CreateAtomicBankRequest**](CreateAtomicBankRequest.md) |  | 

### Return type

[**AtomicBank**](AtomicBank.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AtomicBanksDelete

> AtomicBanksDelete(ctx, id).Execute()



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
    resp, r, err := apiClient.BanksApi.AtomicBanksDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BanksApi.AtomicBanksDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAtomicBanksDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AtomicBanksGet

> AtomicBankPaginatedList AtomicBanksGet(ctx).Page(page).Size(size).Execute()



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
    page := int32(56) // int32 |  (optional)
    size := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BanksApi.AtomicBanksGet(context.Background()).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BanksApi.AtomicBanksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AtomicBanksGet`: AtomicBankPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `BanksApi.AtomicBanksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAtomicBanksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **size** | **int32** |  | 

### Return type

[**AtomicBankPaginatedList**](AtomicBankPaginatedList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AtomicBanksGetById

> AtomicBank AtomicBanksGetById(ctx, id).Execute()



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
    resp, r, err := apiClient.BanksApi.AtomicBanksGetById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BanksApi.AtomicBanksGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AtomicBanksGetById`: AtomicBank
    fmt.Fprintf(os.Stdout, "Response from `BanksApi.AtomicBanksGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAtomicBanksGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AtomicBank**](AtomicBank.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AtomicBanksReact

> ReactResponse AtomicBanksReact(ctx, bankId).AtomicReactRequest(atomicReactRequest).Execute()



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
    bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    atomicReactRequest := *openapiclient.NewAtomicReactRequest("ReactorId_example") // AtomicReactRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BanksApi.AtomicBanksReact(context.Background(), bankId).AtomicReactRequest(atomicReactRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BanksApi.AtomicBanksReact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AtomicBanksReact`: ReactResponse
    fmt.Fprintf(os.Stdout, "Response from `BanksApi.AtomicBanksReact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAtomicBanksReactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **atomicReactRequest** | [**AtomicReactRequest**](AtomicReactRequest.md) |  | 

### Return type

[**ReactResponse**](ReactResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AtomicBanksUpdate

> AtomicBank AtomicBanksUpdate(ctx, id).UpdateAtomicBankRequest(updateAtomicBankRequest).Execute()



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
    updateAtomicBankRequest := *openapiclient.NewUpdateAtomicBankRequest(*openapiclient.NewBank("RoutingNumber_example", "AccountNumber_example")) // UpdateAtomicBankRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BanksApi.AtomicBanksUpdate(context.Background(), id).UpdateAtomicBankRequest(updateAtomicBankRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BanksApi.AtomicBanksUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AtomicBanksUpdate`: AtomicBank
    fmt.Fprintf(os.Stdout, "Response from `BanksApi.AtomicBanksUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAtomicBanksUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAtomicBankRequest** | [**UpdateAtomicBankRequest**](UpdateAtomicBankRequest.md) |  | 

### Return type

[**AtomicBank**](AtomicBank.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


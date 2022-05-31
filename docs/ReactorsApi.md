# \ReactorsApi

All URIs are relative to *https://api.basistheory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReactorsCreate**](ReactorsApi.md#ReactorsCreate) | **Post** /reactors | 
[**ReactorsDelete**](ReactorsApi.md#ReactorsDelete) | **Delete** /reactors/{id} | 
[**ReactorsGet**](ReactorsApi.md#ReactorsGet) | **Get** /reactors | 
[**ReactorsGetById**](ReactorsApi.md#ReactorsGetById) | **Get** /reactors/{id} | 
[**ReactorsReact**](ReactorsApi.md#ReactorsReact) | **Post** /reactors/{id}/react | 
[**ReactorsUpdate**](ReactorsApi.md#ReactorsUpdate) | **Put** /reactors/{id} | 



## ReactorsCreate

> Reactor ReactorsCreate(ctx).CreateReactorRequest(createReactorRequest).Execute()



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
    createReactorRequest := *openapiclient.NewCreateReactorRequest("Name_example") // CreateReactorRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactorsApi.ReactorsCreate(context.Background()).CreateReactorRequest(createReactorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactorsApi.ReactorsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactorsCreate`: Reactor
    fmt.Fprintf(os.Stdout, "Response from `ReactorsApi.ReactorsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReactorsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createReactorRequest** | [**CreateReactorRequest**](CreateReactorRequest.md) |  | 

### Return type

[**Reactor**](Reactor.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactorsDelete

> ReactorsDelete(ctx, id).Execute()



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
    resp, r, err := apiClient.ReactorsApi.ReactorsDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactorsApi.ReactorsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReactorsDeleteRequest struct via the builder pattern


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


## ReactorsGet

> ReactorPaginatedList ReactorsGet(ctx).Id(id).Name(name).Page(page).Size(size).Execute()



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
    id := []string{"Inner_example"} // []string |  (optional)
    name := "name_example" // string |  (optional)
    page := int32(56) // int32 |  (optional)
    size := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactorsApi.ReactorsGet(context.Background()).Id(id).Name(name).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactorsApi.ReactorsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactorsGet`: ReactorPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `ReactorsApi.ReactorsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReactorsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** |  | 
 **name** | **string** |  | 
 **page** | **int32** |  | 
 **size** | **int32** |  | 

### Return type

[**ReactorPaginatedList**](ReactorPaginatedList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactorsGetById

> Reactor ReactorsGetById(ctx, id).Execute()



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
    resp, r, err := apiClient.ReactorsApi.ReactorsGetById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactorsApi.ReactorsGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactorsGetById`: Reactor
    fmt.Fprintf(os.Stdout, "Response from `ReactorsApi.ReactorsGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactorsGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Reactor**](Reactor.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactorsReact

> ReactResponse ReactorsReact(ctx, id).ReactRequest(reactRequest).Execute()



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
    reactRequest := *openapiclient.NewReactRequest() // ReactRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactorsApi.ReactorsReact(context.Background(), id).ReactRequest(reactRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactorsApi.ReactorsReact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactorsReact`: ReactResponse
    fmt.Fprintf(os.Stdout, "Response from `ReactorsApi.ReactorsReact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactorsReactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reactRequest** | [**ReactRequest**](ReactRequest.md) |  | 

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


## ReactorsUpdate

> Reactor ReactorsUpdate(ctx, id).UpdateReactorRequest(updateReactorRequest).Execute()



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
    updateReactorRequest := *openapiclient.NewUpdateReactorRequest("Name_example") // UpdateReactorRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactorsApi.ReactorsUpdate(context.Background(), id).UpdateReactorRequest(updateReactorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactorsApi.ReactorsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactorsUpdate`: Reactor
    fmt.Fprintf(os.Stdout, "Response from `ReactorsApi.ReactorsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactorsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateReactorRequest** | [**UpdateReactorRequest**](UpdateReactorRequest.md) |  | 

### Return type

[**Reactor**](Reactor.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


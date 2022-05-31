# \InboundProxiesApi

All URIs are relative to *https://api.basistheory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InboundProxiesCreate**](InboundProxiesApi.md#InboundProxiesCreate) | **Post** /inbound-proxies | 
[**InboundProxiesDelete**](InboundProxiesApi.md#InboundProxiesDelete) | **Delete** /inbound-proxies/{id} | 
[**InboundProxiesGet**](InboundProxiesApi.md#InboundProxiesGet) | **Get** /inbound-proxies | 
[**InboundProxiesGetById**](InboundProxiesApi.md#InboundProxiesGetById) | **Get** /inbound-proxies/{id} | 
[**InboundProxiesUpdate**](InboundProxiesApi.md#InboundProxiesUpdate) | **Put** /inbound-proxies/{id} | 



## InboundProxiesCreate

> InboundProxy InboundProxiesCreate(ctx).CreateInboundProxyRequest(createInboundProxyRequest).Execute()



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
    createInboundProxyRequest := *openapiclient.NewCreateInboundProxyRequest("Name_example", "DestinationUrl_example", "RequestReactorId_example") // CreateInboundProxyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InboundProxiesApi.InboundProxiesCreate(context.Background()).CreateInboundProxyRequest(createInboundProxyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboundProxiesApi.InboundProxiesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InboundProxiesCreate`: InboundProxy
    fmt.Fprintf(os.Stdout, "Response from `InboundProxiesApi.InboundProxiesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInboundProxiesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInboundProxyRequest** | [**CreateInboundProxyRequest**](CreateInboundProxyRequest.md) |  | 

### Return type

[**InboundProxy**](InboundProxy.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InboundProxiesDelete

> InboundProxiesDelete(ctx, id).Execute()



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
    resp, r, err := apiClient.InboundProxiesApi.InboundProxiesDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboundProxiesApi.InboundProxiesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInboundProxiesDeleteRequest struct via the builder pattern


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


## InboundProxiesGet

> InboundProxyPaginatedList InboundProxiesGet(ctx).Id(id).Name(name).Page(page).Size(size).Execute()



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
    resp, r, err := apiClient.InboundProxiesApi.InboundProxiesGet(context.Background()).Id(id).Name(name).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboundProxiesApi.InboundProxiesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InboundProxiesGet`: InboundProxyPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `InboundProxiesApi.InboundProxiesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInboundProxiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** |  | 
 **name** | **string** |  | 
 **page** | **int32** |  | 
 **size** | **int32** |  | 

### Return type

[**InboundProxyPaginatedList**](InboundProxyPaginatedList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InboundProxiesGetById

> InboundProxy InboundProxiesGetById(ctx, id).Execute()



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
    resp, r, err := apiClient.InboundProxiesApi.InboundProxiesGetById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboundProxiesApi.InboundProxiesGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InboundProxiesGetById`: InboundProxy
    fmt.Fprintf(os.Stdout, "Response from `InboundProxiesApi.InboundProxiesGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInboundProxiesGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InboundProxy**](InboundProxy.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InboundProxiesUpdate

> InboundProxy InboundProxiesUpdate(ctx, id).UpdateInboundProxyRequest(updateInboundProxyRequest).Execute()



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
    updateInboundProxyRequest := *openapiclient.NewUpdateInboundProxyRequest("Name_example", "DestinationUrl_example", "RequestReactorId_example") // UpdateInboundProxyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InboundProxiesApi.InboundProxiesUpdate(context.Background(), id).UpdateInboundProxyRequest(updateInboundProxyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboundProxiesApi.InboundProxiesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InboundProxiesUpdate`: InboundProxy
    fmt.Fprintf(os.Stdout, "Response from `InboundProxiesApi.InboundProxiesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInboundProxiesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInboundProxyRequest** | [**UpdateInboundProxyRequest**](UpdateInboundProxyRequest.md) |  | 

### Return type

[**InboundProxy**](InboundProxy.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


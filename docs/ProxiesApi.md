# \ProxiesApi

All URIs are relative to *https://api.basistheory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProxiesCreate**](ProxiesApi.md#ProxiesCreate) | **Post** /proxies | 
[**ProxiesDelete**](ProxiesApi.md#ProxiesDelete) | **Delete** /proxies/{id} | 
[**ProxiesGet**](ProxiesApi.md#ProxiesGet) | **Get** /proxies | 
[**ProxiesGetById**](ProxiesApi.md#ProxiesGetById) | **Get** /proxies/{id} | 
[**ProxiesUpdate**](ProxiesApi.md#ProxiesUpdate) | **Put** /proxies/{id} | 



## ProxiesCreate

> Proxy ProxiesCreate(ctx).CreateProxyRequest(createProxyRequest).Execute()



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
    createProxyRequest := *openapiclient.NewCreateProxyRequest("Name_example", "DestinationUrl_example", "RequestReactorId_example") // CreateProxyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProxiesApi.ProxiesCreate(context.Background()).CreateProxyRequest(createProxyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxiesApi.ProxiesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProxiesCreate`: Proxy
    fmt.Fprintf(os.Stdout, "Response from `ProxiesApi.ProxiesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProxiesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProxyRequest** | [**CreateProxyRequest**](CreateProxyRequest.md) |  | 

### Return type

[**Proxy**](Proxy.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxiesDelete

> ProxiesDelete(ctx, id).Execute()



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
    resp, r, err := apiClient.ProxiesApi.ProxiesDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxiesApi.ProxiesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProxiesDeleteRequest struct via the builder pattern


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


## ProxiesGet

> ProxyPaginatedList ProxiesGet(ctx).Id(id).Name(name).Page(page).Size(size).Execute()



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
    resp, r, err := apiClient.ProxiesApi.ProxiesGet(context.Background()).Id(id).Name(name).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxiesApi.ProxiesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProxiesGet`: ProxyPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `ProxiesApi.ProxiesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProxiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** |  | 
 **name** | **string** |  | 
 **page** | **int32** |  | 
 **size** | **int32** |  | 

### Return type

[**ProxyPaginatedList**](ProxyPaginatedList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxiesGetById

> Proxy ProxiesGetById(ctx, id).Execute()



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
    resp, r, err := apiClient.ProxiesApi.ProxiesGetById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxiesApi.ProxiesGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProxiesGetById`: Proxy
    fmt.Fprintf(os.Stdout, "Response from `ProxiesApi.ProxiesGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxiesGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Proxy**](Proxy.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxiesUpdate

> Proxy ProxiesUpdate(ctx, id).UpdateProxyRequest(updateProxyRequest).Execute()



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
    updateProxyRequest := *openapiclient.NewUpdateProxyRequest("Name_example", "DestinationUrl_example", "RequestReactorId_example") // UpdateProxyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProxiesApi.ProxiesUpdate(context.Background(), id).UpdateProxyRequest(updateProxyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxiesApi.ProxiesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProxiesUpdate`: Proxy
    fmt.Fprintf(os.Stdout, "Response from `ProxiesApi.ProxiesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxiesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProxyRequest** | [**UpdateProxyRequest**](UpdateProxyRequest.md) |  | 

### Return type

[**Proxy**](Proxy.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \CardsApi

All URIs are relative to *https://api.basistheory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AtomicCardsCreate**](CardsApi.md#AtomicCardsCreate) | **Post** /atomic/cards | 
[**AtomicCardsDelete**](CardsApi.md#AtomicCardsDelete) | **Delete** /atomic/cards/{id} | 
[**AtomicCardsGet**](CardsApi.md#AtomicCardsGet) | **Get** /atomic/cards | 
[**AtomicCardsGetById**](CardsApi.md#AtomicCardsGetById) | **Get** /atomic/cards/{id} | 
[**AtomicCardsReact**](CardsApi.md#AtomicCardsReact) | **Post** /atomic/cards/{cardId}/react | 
[**AtomicCardsUpdate**](CardsApi.md#AtomicCardsUpdate) | **Patch** /atomic/cards/{id} | 



## AtomicCardsCreate

> AtomicCard AtomicCardsCreate(ctx).CreateAtomicCardRequest(createAtomicCardRequest).Execute()



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
    createAtomicCardRequest := *openapiclient.NewCreateAtomicCardRequest(*openapiclient.NewCard()) // CreateAtomicCardRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.AtomicCardsCreate(context.Background()).CreateAtomicCardRequest(createAtomicCardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.AtomicCardsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AtomicCardsCreate`: AtomicCard
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.AtomicCardsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAtomicCardsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAtomicCardRequest** | [**CreateAtomicCardRequest**](CreateAtomicCardRequest.md) |  | 

### Return type

[**AtomicCard**](AtomicCard.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AtomicCardsDelete

> AtomicCardsDelete(ctx, id).Execute()



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
    resp, r, err := apiClient.CardsApi.AtomicCardsDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.AtomicCardsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAtomicCardsDeleteRequest struct via the builder pattern


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


## AtomicCardsGet

> AtomicCardPaginatedList AtomicCardsGet(ctx).Page(page).Size(size).Execute()



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
    resp, r, err := apiClient.CardsApi.AtomicCardsGet(context.Background()).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.AtomicCardsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AtomicCardsGet`: AtomicCardPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.AtomicCardsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAtomicCardsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **size** | **int32** |  | 

### Return type

[**AtomicCardPaginatedList**](AtomicCardPaginatedList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AtomicCardsGetById

> AtomicCard AtomicCardsGetById(ctx, id).Execute()



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
    resp, r, err := apiClient.CardsApi.AtomicCardsGetById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.AtomicCardsGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AtomicCardsGetById`: AtomicCard
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.AtomicCardsGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAtomicCardsGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AtomicCard**](AtomicCard.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AtomicCardsReact

> ReactResponse AtomicCardsReact(ctx, cardId).AtomicReactRequest(atomicReactRequest).Execute()



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
    cardId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    atomicReactRequest := *openapiclient.NewAtomicReactRequest("ReactorId_example") // AtomicReactRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.AtomicCardsReact(context.Background(), cardId).AtomicReactRequest(atomicReactRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.AtomicCardsReact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AtomicCardsReact`: ReactResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.AtomicCardsReact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAtomicCardsReactRequest struct via the builder pattern


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


## AtomicCardsUpdate

> AtomicCard AtomicCardsUpdate(ctx, id).UpdateAtomicCardRequest(updateAtomicCardRequest).Execute()



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
    updateAtomicCardRequest := *openapiclient.NewUpdateAtomicCardRequest(*openapiclient.NewCard()) // UpdateAtomicCardRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.AtomicCardsUpdate(context.Background(), id).UpdateAtomicCardRequest(updateAtomicCardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.AtomicCardsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AtomicCardsUpdate`: AtomicCard
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.AtomicCardsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAtomicCardsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAtomicCardRequest** | [**UpdateAtomicCardRequest**](UpdateAtomicCardRequest.md) |  | 

### Return type

[**AtomicCard**](AtomicCard.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


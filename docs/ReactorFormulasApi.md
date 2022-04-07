# \ReactorFormulasApi

All URIs are relative to *https://api.basistheory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReactorFormulaCreate**](ReactorFormulasApi.md#ReactorFormulaCreate) | **Post** /reactor-formulas | 
[**ReactorFormulaDelete**](ReactorFormulasApi.md#ReactorFormulaDelete) | **Delete** /reactor-formulas/{id} | 
[**ReactorFormulaGetById**](ReactorFormulasApi.md#ReactorFormulaGetById) | **Get** /reactor-formulas/{id} | 
[**ReactorFormulaUpdate**](ReactorFormulasApi.md#ReactorFormulaUpdate) | **Put** /reactor-formulas/{id} | 
[**ReactorFormulasGet**](ReactorFormulasApi.md#ReactorFormulasGet) | **Get** /reactor-formulas | 



## ReactorFormulaCreate

> ReactorFormulaModel ReactorFormulaCreate(ctx).CreateReactorFormulaModel(createReactorFormulaModel).Execute()



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
    createReactorFormulaModel := *openapiclient.NewCreateReactorFormulaModel() // CreateReactorFormulaModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactorFormulasApi.ReactorFormulaCreate(context.Background()).CreateReactorFormulaModel(createReactorFormulaModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactorFormulasApi.ReactorFormulaCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactorFormulaCreate`: ReactorFormulaModel
    fmt.Fprintf(os.Stdout, "Response from `ReactorFormulasApi.ReactorFormulaCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReactorFormulaCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createReactorFormulaModel** | [**CreateReactorFormulaModel**](CreateReactorFormulaModel.md) |  | 

### Return type

[**ReactorFormulaModel**](ReactorFormulaModel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactorFormulaDelete

> ReactorFormulaModel ReactorFormulaDelete(ctx, id).Execute()



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
    resp, r, err := apiClient.ReactorFormulasApi.ReactorFormulaDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactorFormulasApi.ReactorFormulaDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactorFormulaDelete`: ReactorFormulaModel
    fmt.Fprintf(os.Stdout, "Response from `ReactorFormulasApi.ReactorFormulaDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactorFormulaDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReactorFormulaModel**](ReactorFormulaModel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactorFormulaGetById

> ReactorFormulaModel ReactorFormulaGetById(ctx, id).Execute()



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
    resp, r, err := apiClient.ReactorFormulasApi.ReactorFormulaGetById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactorFormulasApi.ReactorFormulaGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactorFormulaGetById`: ReactorFormulaModel
    fmt.Fprintf(os.Stdout, "Response from `ReactorFormulasApi.ReactorFormulaGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactorFormulaGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReactorFormulaModel**](ReactorFormulaModel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactorFormulaUpdate

> ReactorFormulaModel ReactorFormulaUpdate(ctx, id).UpdateReactorFormulaModel(updateReactorFormulaModel).Execute()



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
    updateReactorFormulaModel := *openapiclient.NewUpdateReactorFormulaModel() // UpdateReactorFormulaModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactorFormulasApi.ReactorFormulaUpdate(context.Background(), id).UpdateReactorFormulaModel(updateReactorFormulaModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactorFormulasApi.ReactorFormulaUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactorFormulaUpdate`: ReactorFormulaModel
    fmt.Fprintf(os.Stdout, "Response from `ReactorFormulasApi.ReactorFormulaUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactorFormulaUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateReactorFormulaModel** | [**UpdateReactorFormulaModel**](UpdateReactorFormulaModel.md) |  | 

### Return type

[**ReactorFormulaModel**](ReactorFormulaModel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactorFormulasGet

> ReactorFormulaModelPaginatedList ReactorFormulasGet(ctx).Name(name).Page(page).Size(size).Execute()



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
    name := "name_example" // string |  (optional)
    page := int32(56) // int32 |  (optional)
    size := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactorFormulasApi.ReactorFormulasGet(context.Background()).Name(name).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactorFormulasApi.ReactorFormulasGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactorFormulasGet`: ReactorFormulaModelPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `ReactorFormulasApi.ReactorFormulasGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReactorFormulasGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **page** | **int32** |  | 
 **size** | **int32** |  | 

### Return type

[**ReactorFormulaModelPaginatedList**](ReactorFormulaModelPaginatedList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


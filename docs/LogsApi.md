# \LogsApi

All URIs are relative to *https://api.basistheory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LogsGet**](LogsApi.md#LogsGet) | **Get** /logs | 
[**LogsGetEntityTypes**](LogsApi.md#LogsGetEntityTypes) | **Get** /logs/entity-types | 



## LogsGet

> LogPaginatedList LogsGet(ctx).EntityType(entityType).EntityId(entityId).StartDate(startDate).EndDate(endDate).Page(page).Size(size).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    entityType := "entityType_example" // string |  (optional)
    entityId := "entityId_example" // string |  (optional)
    startDate := time.Now() // time.Time |  (optional)
    endDate := time.Now() // time.Time |  (optional)
    page := int32(56) // int32 |  (optional)
    size := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.LogsGet(context.Background()).EntityType(entityType).EntityId(entityId).StartDate(startDate).EndDate(endDate).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.LogsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogsGet`: LogPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.LogsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityType** | **string** |  | 
 **entityId** | **string** |  | 
 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 
 **page** | **int32** |  | 
 **size** | **int32** |  | 

### Return type

[**LogPaginatedList**](LogPaginatedList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogsGetEntityTypes

> []LogEntityType LogsGetEntityTypes(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.LogsGetEntityTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.LogsGetEntityTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogsGetEntityTypes`: []LogEntityType
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.LogsGetEntityTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogsGetEntityTypesRequest struct via the builder pattern


### Return type

[**[]LogEntityType**](LogEntityType.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


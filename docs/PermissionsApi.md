# \PermissionsApi

All URIs are relative to *https://api.basistheory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](PermissionsApi.md#Get) | **Get** /permissions | 



## Get

> []Permission Get(ctx).ApplicationType(applicationType).Version(version).Execute()



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
    applicationType := "applicationType_example" // string |  (optional)
    version := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.Get(context.Background()).ApplicationType(applicationType).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: []Permission
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationType** | **string** |  | 
 **version** | **int32** |  | 

### Return type

[**[]Permission**](Permission.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


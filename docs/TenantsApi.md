# \TenantsApi

All URIs are relative to *https://api.basistheory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TenantInvitationsCreateInvitation**](TenantsApi.md#TenantInvitationsCreateInvitation) | **Post** /tenants/self/invitations | 
[**TenantInvitationsDeleteInvitation**](TenantsApi.md#TenantInvitationsDeleteInvitation) | **Delete** /tenants/self/invitations/{invitationId} | 
[**TenantInvitationsGetInvitations**](TenantsApi.md#TenantInvitationsGetInvitations) | **Get** /tenants/self/invitations | 
[**TenantInvitationsResendInvitation**](TenantsApi.md#TenantInvitationsResendInvitation) | **Post** /tenants/self/invitations/{invitationId}/resend | 
[**TenantMembersDeleteMember**](TenantsApi.md#TenantMembersDeleteMember) | **Delete** /tenants/self/members/{memberId} | 
[**TenantMembersGetMembers**](TenantsApi.md#TenantMembersGetMembers) | **Get** /tenants/self/members | 
[**TenantReportsGetTenantOperationReport**](TenantsApi.md#TenantReportsGetTenantOperationReport) | **Get** /tenants/self/reports/operations | 
[**TenantReportsGetTenantUsageReport**](TenantsApi.md#TenantReportsGetTenantUsageReport) | **Get** /tenants/self/reports/usage | 
[**TenantsDelete**](TenantsApi.md#TenantsDelete) | **Delete** /tenants/self | 
[**TenantsGet**](TenantsApi.md#TenantsGet) | **Get** /tenants/self | 
[**TenantsUpdate**](TenantsApi.md#TenantsUpdate) | **Put** /tenants/self | 



## TenantInvitationsCreateInvitation

> TenantInvitationResponse TenantInvitationsCreateInvitation(ctx).CreateTenantInvitationRequest(createTenantInvitationRequest).Execute()



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
    createTenantInvitationRequest := *openapiclient.NewCreateTenantInvitationRequest("Email_example") // CreateTenantInvitationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantInvitationsCreateInvitation(context.Background()).CreateTenantInvitationRequest(createTenantInvitationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantInvitationsCreateInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantInvitationsCreateInvitation`: TenantInvitationResponse
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantInvitationsCreateInvitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantInvitationsCreateInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTenantInvitationRequest** | [**CreateTenantInvitationRequest**](CreateTenantInvitationRequest.md) |  | 

### Return type

[**TenantInvitationResponse**](TenantInvitationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantInvitationsDeleteInvitation

> TenantInvitationsDeleteInvitation(ctx, invitationId).Execute()



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
    invitationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantInvitationsDeleteInvitation(context.Background(), invitationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantInvitationsDeleteInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantInvitationsDeleteInvitationRequest struct via the builder pattern


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


## TenantInvitationsGetInvitations

> TenantInvitationResponsePaginatedList TenantInvitationsGetInvitations(ctx).Status(status).Page(page).Size(size).Execute()



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
    status := openapiclient.TenantInvitationStatus("PENDING") // TenantInvitationStatus |  (optional)
    page := int32(56) // int32 |  (optional)
    size := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantInvitationsGetInvitations(context.Background()).Status(status).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantInvitationsGetInvitations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantInvitationsGetInvitations`: TenantInvitationResponsePaginatedList
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantInvitationsGetInvitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantInvitationsGetInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**TenantInvitationStatus**](TenantInvitationStatus.md) |  | 
 **page** | **int32** |  | 
 **size** | **int32** |  | 

### Return type

[**TenantInvitationResponsePaginatedList**](TenantInvitationResponsePaginatedList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantInvitationsResendInvitation

> TenantInvitationResponse TenantInvitationsResendInvitation(ctx, invitationId).Execute()



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
    invitationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantInvitationsResendInvitation(context.Background(), invitationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantInvitationsResendInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantInvitationsResendInvitation`: TenantInvitationResponse
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantInvitationsResendInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantInvitationsResendInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TenantInvitationResponse**](TenantInvitationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantMembersDeleteMember

> TenantMembersDeleteMember(ctx, memberId).Execute()



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
    memberId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantMembersDeleteMember(context.Background(), memberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantMembersDeleteMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantMembersDeleteMemberRequest struct via the builder pattern


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


## TenantMembersGetMembers

> TenantMemberResponsePaginatedList TenantMembersGetMembers(ctx).UserId(userId).Page(page).Size(size).Execute()



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
    userId := []string{"Inner_example"} // []string |  (optional)
    page := int32(56) // int32 |  (optional)
    size := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantMembersGetMembers(context.Background()).UserId(userId).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantMembersGetMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantMembersGetMembers`: TenantMemberResponsePaginatedList
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantMembersGetMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantMembersGetMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **[]string** |  | 
 **page** | **int32** |  | 
 **size** | **int32** |  | 

### Return type

[**TenantMemberResponsePaginatedList**](TenantMemberResponsePaginatedList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantReportsGetTenantOperationReport

> TenantUsageReport TenantReportsGetTenantOperationReport(ctx).Execute()



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
    resp, r, err := apiClient.TenantsApi.TenantReportsGetTenantOperationReport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantReportsGetTenantOperationReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantReportsGetTenantOperationReport`: TenantUsageReport
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantReportsGetTenantOperationReport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantReportsGetTenantOperationReportRequest struct via the builder pattern


### Return type

[**TenantUsageReport**](TenantUsageReport.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantReportsGetTenantUsageReport

> TenantUsageReport TenantReportsGetTenantUsageReport(ctx).Execute()



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
    resp, r, err := apiClient.TenantsApi.TenantReportsGetTenantUsageReport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantReportsGetTenantUsageReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantReportsGetTenantUsageReport`: TenantUsageReport
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantReportsGetTenantUsageReport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantReportsGetTenantUsageReportRequest struct via the builder pattern


### Return type

[**TenantUsageReport**](TenantUsageReport.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsDelete

> TenantsDelete(ctx).Execute()



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
    resp, r, err := apiClient.TenantsApi.TenantsDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantsDeleteRequest struct via the builder pattern


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


## TenantsGet

> Tenant TenantsGet(ctx).Execute()



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
    resp, r, err := apiClient.TenantsApi.TenantsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantsGet`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantsGetRequest struct via the builder pattern


### Return type

[**Tenant**](Tenant.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsUpdate

> Tenant TenantsUpdate(ctx).UpdateTenantRequest(updateTenantRequest).Execute()



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
    updateTenantRequest := *openapiclient.NewUpdateTenantRequest("Name_example") // UpdateTenantRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantsUpdate(context.Background()).UpdateTenantRequest(updateTenantRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantsUpdate`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTenantRequest** | [**UpdateTenantRequest**](UpdateTenantRequest.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# Basis Theory Go SDK

[![Release](https://github.com/Basis-Theory/basistheory-go/actions/workflows/release.yml/badge.svg)](https://github.com/Basis-Theory/basistheory-go/actions/workflows/release.yml)

## Getting Started
* Sign-in to [Basis Theory](https://basistheory.com) and go to [Applications](https://portal.basistheory.com/applications)
* Create a Basis Theory Private Application
* All permissions should be selected
* Paste the API Key into the `BT-API-KEY` variable

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import basistheory "github.com/Basis-Theory/basistheory-go/v3"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Running tests locally

To run tests locally, you'll need to create a Basis Theory tenant with private and management applications 
that are granted all permissions, then add the API keys for these applications to a `.env.local` file. 
You'll want to follow the same format as defined in the `.env.example` file. Then you can run `make verify` from
the root of this repository to run all tests.

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), basistheory.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), basistheory.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), basistheory.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), basistheory.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.basistheory.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApplicationTemplatesApi* | [**Get**](docs/ApplicationTemplatesApi.md#get) | **Get** /application-templates | 
*ApplicationTemplatesApi* | [**GetById**](docs/ApplicationTemplatesApi.md#getbyid) | **Get** /application-templates/{id} | 
*ApplicationsApi* | [**Create**](docs/ApplicationsApi.md#create) | **Post** /applications | 
*ApplicationsApi* | [**Delete**](docs/ApplicationsApi.md#delete) | **Delete** /applications/{id} | 
*ApplicationsApi* | [**Get**](docs/ApplicationsApi.md#get) | **Get** /applications | 
*ApplicationsApi* | [**GetById**](docs/ApplicationsApi.md#getbyid) | **Get** /applications/{id} | 
*ApplicationsApi* | [**GetByKey**](docs/ApplicationsApi.md#getbykey) | **Get** /applications/key | 
*ApplicationsApi* | [**RegenerateKey**](docs/ApplicationsApi.md#regeneratekey) | **Post** /applications/{id}/regenerate | 
*ApplicationsApi* | [**Update**](docs/ApplicationsApi.md#update) | **Put** /applications/{id} | 
*LogsApi* | [**Get**](docs/LogsApi.md#get) | **Get** /logs | 
*LogsApi* | [**GetEntityTypes**](docs/LogsApi.md#getentitytypes) | **Get** /logs/entity-types | 
*PermissionsApi* | [**Get**](docs/PermissionsApi.md#get) | **Get** /permissions | 
*ProxiesApi* | [**Create**](docs/ProxiesApi.md#create) | **Post** /proxies | 
*ProxiesApi* | [**Delete**](docs/ProxiesApi.md#delete) | **Delete** /proxies/{id} | 
*ProxiesApi* | [**Get**](docs/ProxiesApi.md#get) | **Get** /proxies | 
*ProxiesApi* | [**GetById**](docs/ProxiesApi.md#getbyid) | **Get** /proxies/{id} | 
*ProxiesApi* | [**Patch**](docs/ProxiesApi.md#patch) | **Patch** /proxies/{id} | 
*ProxiesApi* | [**Update**](docs/ProxiesApi.md#update) | **Put** /proxies/{id} | 
*ReactorFormulasApi* | [**Create**](docs/ReactorFormulasApi.md#create) | **Post** /reactor-formulas | 
*ReactorFormulasApi* | [**Delete**](docs/ReactorFormulasApi.md#delete) | **Delete** /reactor-formulas/{id} | 
*ReactorFormulasApi* | [**Get**](docs/ReactorFormulasApi.md#get) | **Get** /reactor-formulas | 
*ReactorFormulasApi* | [**GetById**](docs/ReactorFormulasApi.md#getbyid) | **Get** /reactor-formulas/{id} | 
*ReactorFormulasApi* | [**Update**](docs/ReactorFormulasApi.md#update) | **Put** /reactor-formulas/{id} | 
*ReactorsApi* | [**Create**](docs/ReactorsApi.md#create) | **Post** /reactors | 
*ReactorsApi* | [**Delete**](docs/ReactorsApi.md#delete) | **Delete** /reactors/{id} | 
*ReactorsApi* | [**Get**](docs/ReactorsApi.md#get) | **Get** /reactors | 
*ReactorsApi* | [**GetById**](docs/ReactorsApi.md#getbyid) | **Get** /reactors/{id} | 
*ReactorsApi* | [**Patch**](docs/ReactorsApi.md#patch) | **Patch** /reactors/{id} | 
*ReactorsApi* | [**React**](docs/ReactorsApi.md#react) | **Post** /reactors/{id}/react | 
*ReactorsApi* | [**Update**](docs/ReactorsApi.md#update) | **Put** /reactors/{id} | 
*SessionsApi* | [**Authorize**](docs/SessionsApi.md#authorize) | **Post** /sessions/authorize | 
*SessionsApi* | [**Create**](docs/SessionsApi.md#create) | **Post** /sessions | 
*TenantsApi* | [**CreateInvitation**](docs/TenantsApi.md#createinvitation) | **Post** /tenants/self/invitations | 
*TenantsApi* | [**Delete**](docs/TenantsApi.md#delete) | **Delete** /tenants/self | 
*TenantsApi* | [**DeleteInvitation**](docs/TenantsApi.md#deleteinvitation) | **Delete** /tenants/self/invitations/{invitationId} | 
*TenantsApi* | [**DeleteMember**](docs/TenantsApi.md#deletemember) | **Delete** /tenants/self/members/{memberId} | 
*TenantsApi* | [**Get**](docs/TenantsApi.md#get) | **Get** /tenants/self | 
*TenantsApi* | [**GetInvitations**](docs/TenantsApi.md#getinvitations) | **Get** /tenants/self/invitations | 
*TenantsApi* | [**GetMembers**](docs/TenantsApi.md#getmembers) | **Get** /tenants/self/members | 
*TenantsApi* | [**GetTenantUsageReport**](docs/TenantsApi.md#gettenantusagereport) | **Get** /tenants/self/reports/usage | 
*TenantsApi* | [**ResendInvitation**](docs/TenantsApi.md#resendinvitation) | **Post** /tenants/self/invitations/{invitationId}/resend | 
*TenantsApi* | [**Update**](docs/TenantsApi.md#update) | **Put** /tenants/self | 
*ThreeDSApi* | [**ThreeDSAuthenticateSession**](docs/ThreeDSApi.md#threedsauthenticatesession) | **Post** /3ds/sessions/{sessionId}/authenticate | 
*ThreeDSApi* | [**ThreeDSGetChallengeResult**](docs/ThreeDSApi.md#threedsgetchallengeresult) | **Get** /3ds/sessions/{sessionId}/challenge-result | 
*TokenizeApi* | [**Tokenize**](docs/TokenizeApi.md#tokenize) | **Post** /tokenize | 
*TokensApi* | [**Create**](docs/TokensApi.md#create) | **Post** /tokens | 
*TokensApi* | [**Delete**](docs/TokensApi.md#delete) | **Delete** /tokens/{id} | 
*TokensApi* | [**Get**](docs/TokensApi.md#get) | **Get** /tokens | 
*TokensApi* | [**GetById**](docs/TokensApi.md#getbyid) | **Get** /tokens/{id} | 
*TokensApi* | [**Search**](docs/TokensApi.md#search) | **Post** /tokens/search | 
*TokensApi* | [**Update**](docs/TokensApi.md#update) | **Patch** /tokens/{id} | 


## Documentation For Models

 - [AccessRule](docs/AccessRule.md)
 - [Application](docs/Application.md)
 - [ApplicationPaginatedList](docs/ApplicationPaginatedList.md)
 - [ApplicationTemplate](docs/ApplicationTemplate.md)
 - [AuthenticateThreeDSSessionRequest](docs/AuthenticateThreeDSSessionRequest.md)
 - [AuthorizeSessionRequest](docs/AuthorizeSessionRequest.md)
 - [BinDetails](docs/BinDetails.md)
 - [BinDetailsBank](docs/BinDetailsBank.md)
 - [BinDetailsCountry](docs/BinDetailsCountry.md)
 - [BinDetailsProduct](docs/BinDetailsProduct.md)
 - [Condition](docs/Condition.md)
 - [CreateApplicationRequest](docs/CreateApplicationRequest.md)
 - [CreateProxyRequest](docs/CreateProxyRequest.md)
 - [CreateReactorFormulaRequest](docs/CreateReactorFormulaRequest.md)
 - [CreateReactorRequest](docs/CreateReactorRequest.md)
 - [CreateSessionResponse](docs/CreateSessionResponse.md)
 - [CreateTenantInvitationRequest](docs/CreateTenantInvitationRequest.md)
 - [CreateTokenRequest](docs/CreateTokenRequest.md)
 - [EncryptionKey](docs/EncryptionKey.md)
 - [EncryptionMetadata](docs/EncryptionMetadata.md)
 - [GetApplications](docs/GetApplications.md)
 - [GetLogs](docs/GetLogs.md)
 - [GetPermissions](docs/GetPermissions.md)
 - [GetProxies](docs/GetProxies.md)
 - [GetReactorFormulas](docs/GetReactorFormulas.md)
 - [GetReactors](docs/GetReactors.md)
 - [GetTenantInvitations](docs/GetTenantInvitations.md)
 - [GetTenantMembers](docs/GetTenantMembers.md)
 - [GetTokens](docs/GetTokens.md)
 - [Log](docs/Log.md)
 - [LogEntityType](docs/LogEntityType.md)
 - [LogPaginatedList](docs/LogPaginatedList.md)
 - [Pagination](docs/Pagination.md)
 - [PatchProxyRequest](docs/PatchProxyRequest.md)
 - [PatchReactorRequest](docs/PatchReactorRequest.md)
 - [Permission](docs/Permission.md)
 - [Privacy](docs/Privacy.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [Proxy](docs/Proxy.md)
 - [ProxyPaginatedList](docs/ProxyPaginatedList.md)
 - [ProxyTransform](docs/ProxyTransform.md)
 - [ReactRequest](docs/ReactRequest.md)
 - [ReactResponse](docs/ReactResponse.md)
 - [Reactor](docs/Reactor.md)
 - [ReactorFormula](docs/ReactorFormula.md)
 - [ReactorFormulaConfiguration](docs/ReactorFormulaConfiguration.md)
 - [ReactorFormulaPaginatedList](docs/ReactorFormulaPaginatedList.md)
 - [ReactorFormulaRequestParameter](docs/ReactorFormulaRequestParameter.md)
 - [ReactorPaginatedList](docs/ReactorPaginatedList.md)
 - [SearchTokensRequest](docs/SearchTokensRequest.md)
 - [StringStringKeyValuePair](docs/StringStringKeyValuePair.md)
 - [Tenant](docs/Tenant.md)
 - [TenantInvitationResponse](docs/TenantInvitationResponse.md)
 - [TenantInvitationResponsePaginatedList](docs/TenantInvitationResponsePaginatedList.md)
 - [TenantInvitationStatus](docs/TenantInvitationStatus.md)
 - [TenantMemberResponse](docs/TenantMemberResponse.md)
 - [TenantMemberResponsePaginatedList](docs/TenantMemberResponsePaginatedList.md)
 - [TenantUsageReport](docs/TenantUsageReport.md)
 - [ThreeDSAcsRenderingType](docs/ThreeDSAcsRenderingType.md)
 - [ThreeDSAddress](docs/ThreeDSAddress.md)
 - [ThreeDSAuthentication](docs/ThreeDSAuthentication.md)
 - [ThreeDSCardholderAccountInfo](docs/ThreeDSCardholderAccountInfo.md)
 - [ThreeDSCardholderAuthenticationInfo](docs/ThreeDSCardholderAuthenticationInfo.md)
 - [ThreeDSCardholderInfo](docs/ThreeDSCardholderInfo.md)
 - [ThreeDSCardholderPhoneNumber](docs/ThreeDSCardholderPhoneNumber.md)
 - [ThreeDSMerchantInfo](docs/ThreeDSMerchantInfo.md)
 - [ThreeDSMerchantRiskInfo](docs/ThreeDSMerchantRiskInfo.md)
 - [ThreeDSMessageExtension](docs/ThreeDSMessageExtension.md)
 - [ThreeDSPriorAuthenticationInfo](docs/ThreeDSPriorAuthenticationInfo.md)
 - [ThreeDSPurchaseInfo](docs/ThreeDSPurchaseInfo.md)
 - [ThreeDSRequestorInfo](docs/ThreeDSRequestorInfo.md)
 - [Token](docs/Token.md)
 - [TokenEnrichments](docs/TokenEnrichments.md)
 - [TokenMetrics](docs/TokenMetrics.md)
 - [TokenPaginatedList](docs/TokenPaginatedList.md)
 - [TokenReport](docs/TokenReport.md)
 - [UpdateApplicationRequest](docs/UpdateApplicationRequest.md)
 - [UpdatePrivacy](docs/UpdatePrivacy.md)
 - [UpdateProxyRequest](docs/UpdateProxyRequest.md)
 - [UpdateReactorFormulaRequest](docs/UpdateReactorFormulaRequest.md)
 - [UpdateReactorRequest](docs/UpdateReactorRequest.md)
 - [UpdateTenantRequest](docs/UpdateTenantRequest.md)
 - [UpdateTokenRequest](docs/UpdateTokenRequest.md)
 - [User](docs/User.md)
 - [ValidationProblemDetails](docs/ValidationProblemDetails.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### ApiKey

- **Type**: API key
- **API key parameter name**: BT-API-KEY
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: BT-API-KEY and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author




# Go API client for basistheory

## Getting Started
* Sign-in to [Basis Theory](https://basistheory.com) and go to [Applications](https://portal.basistheory.com/applications)
* Create a Basis Theory Server to Server Application
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
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import basistheory "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

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

```
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
*ApplicationsApi* | [**ApplicationsCreate**](docs/ApplicationsApi.md#applicationscreate) | **Post** /applications | 
*ApplicationsApi* | [**ApplicationsDelete**](docs/ApplicationsApi.md#applicationsdelete) | **Delete** /applications/{id} | 
*ApplicationsApi* | [**ApplicationsGet**](docs/ApplicationsApi.md#applicationsget) | **Get** /applications | 
*ApplicationsApi* | [**ApplicationsGetById**](docs/ApplicationsApi.md#applicationsgetbyid) | **Get** /applications/{id} | 
*ApplicationsApi* | [**ApplicationsGetByKey**](docs/ApplicationsApi.md#applicationsgetbykey) | **Get** /applications/key | 
*ApplicationsApi* | [**ApplicationsRegenerateKey**](docs/ApplicationsApi.md#applicationsregeneratekey) | **Post** /applications/{id}/regenerate | 
*ApplicationsApi* | [**ApplicationsUpdate**](docs/ApplicationsApi.md#applicationsupdate) | **Put** /applications/{id} | 
*InboundProxiesApi* | [**InboundProxiesCreate**](docs/InboundProxiesApi.md#inboundproxiescreate) | **Post** /inbound-proxies | 
*InboundProxiesApi* | [**InboundProxiesDelete**](docs/InboundProxiesApi.md#inboundproxiesdelete) | **Delete** /inbound-proxies/{id} | 
*InboundProxiesApi* | [**InboundProxiesGet**](docs/InboundProxiesApi.md#inboundproxiesget) | **Get** /inbound-proxies | 
*InboundProxiesApi* | [**InboundProxiesGetById**](docs/InboundProxiesApi.md#inboundproxiesgetbyid) | **Get** /inbound-proxies/{id} | 
*InboundProxiesApi* | [**InboundProxiesUpdate**](docs/InboundProxiesApi.md#inboundproxiesupdate) | **Put** /inbound-proxies/{id} | 
*ReactorFormulasApi* | [**ReactorFormulasCreate**](docs/ReactorFormulasApi.md#reactorformulascreate) | **Post** /reactor-formulas | 
*ReactorFormulasApi* | [**ReactorFormulasDelete**](docs/ReactorFormulasApi.md#reactorformulasdelete) | **Delete** /reactor-formulas/{id} | 
*ReactorFormulasApi* | [**ReactorFormulasGet**](docs/ReactorFormulasApi.md#reactorformulasget) | **Get** /reactor-formulas | 
*ReactorFormulasApi* | [**ReactorFormulasGetById**](docs/ReactorFormulasApi.md#reactorformulasgetbyid) | **Get** /reactor-formulas/{id} | 
*ReactorFormulasApi* | [**ReactorFormulasUpdate**](docs/ReactorFormulasApi.md#reactorformulasupdate) | **Put** /reactor-formulas/{id} | 
*ReactorsApi* | [**ReactorsCreate**](docs/ReactorsApi.md#reactorscreate) | **Post** /reactors | 
*ReactorsApi* | [**ReactorsDelete**](docs/ReactorsApi.md#reactorsdelete) | **Delete** /reactors/{id} | 
*ReactorsApi* | [**ReactorsGet**](docs/ReactorsApi.md#reactorsget) | **Get** /reactors | 
*ReactorsApi* | [**ReactorsGetById**](docs/ReactorsApi.md#reactorsgetbyid) | **Get** /reactors/{id} | 
*ReactorsApi* | [**ReactorsReact**](docs/ReactorsApi.md#reactorsreact) | **Post** /reactors/{id}/react | 
*ReactorsApi* | [**ReactorsUpdate**](docs/ReactorsApi.md#reactorsupdate) | **Put** /reactors/{id} | 


## Documentation For Models

 - [Application](docs/Application.md)
 - [ApplicationPaginatedList](docs/ApplicationPaginatedList.md)
 - [CreateApplicationRequest](docs/CreateApplicationRequest.md)
 - [CreateInboundProxyRequest](docs/CreateInboundProxyRequest.md)
 - [CreateReactorFormulaRequest](docs/CreateReactorFormulaRequest.md)
 - [CreateReactorRequest](docs/CreateReactorRequest.md)
 - [GetInboundProxies](docs/GetInboundProxies.md)
 - [InboundProxy](docs/InboundProxy.md)
 - [InboundProxyPaginatedList](docs/InboundProxyPaginatedList.md)
 - [Pagination](docs/Pagination.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [ReactRequest](docs/ReactRequest.md)
 - [ReactResponse](docs/ReactResponse.md)
 - [Reactor](docs/Reactor.md)
 - [ReactorFormula](docs/ReactorFormula.md)
 - [ReactorFormulaConfiguration](docs/ReactorFormulaConfiguration.md)
 - [ReactorFormulaPaginatedList](docs/ReactorFormulaPaginatedList.md)
 - [ReactorFormulaRequestParameter](docs/ReactorFormulaRequestParameter.md)
 - [ReactorPaginatedList](docs/ReactorPaginatedList.md)
 - [UpdateApplicationRequest](docs/UpdateApplicationRequest.md)
 - [UpdateInboundProxyRequest](docs/UpdateInboundProxyRequest.md)
 - [UpdateReactorFormulaRequest](docs/UpdateReactorFormulaRequest.md)
 - [UpdateReactorRequest](docs/UpdateReactorRequest.md)
 - [ValidationProblemDetails](docs/ValidationProblemDetails.md)


## Documentation For Authorization



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




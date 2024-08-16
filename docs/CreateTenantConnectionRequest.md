# CreateTenantConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strategy** | **string** |  | 
**Options** | [**TenantConnectionOptions**](TenantConnectionOptions.md) |  | 

## Methods

### NewCreateTenantConnectionRequest

`func NewCreateTenantConnectionRequest(strategy string, options TenantConnectionOptions, ) *CreateTenantConnectionRequest`

NewCreateTenantConnectionRequest instantiates a new CreateTenantConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTenantConnectionRequestWithDefaults

`func NewCreateTenantConnectionRequestWithDefaults() *CreateTenantConnectionRequest`

NewCreateTenantConnectionRequestWithDefaults instantiates a new CreateTenantConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrategy

`func (o *CreateTenantConnectionRequest) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *CreateTenantConnectionRequest) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *CreateTenantConnectionRequest) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.


### GetOptions

`func (o *CreateTenantConnectionRequest) GetOptions() TenantConnectionOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateTenantConnectionRequest) GetOptionsOk() (*TenantConnectionOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateTenantConnectionRequest) SetOptions(v TenantConnectionOptions)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



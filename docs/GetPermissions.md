# GetPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationType** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewGetPermissions

`func NewGetPermissions() *GetPermissions`

NewGetPermissions instantiates a new GetPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPermissionsWithDefaults

`func NewGetPermissionsWithDefaults() *GetPermissions`

NewGetPermissionsWithDefaults instantiates a new GetPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationType

`func (o *GetPermissions) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *GetPermissions) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *GetPermissions) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *GetPermissions) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### SetApplicationTypeNil

`func (o *GetPermissions) SetApplicationTypeNil(b bool)`

 SetApplicationTypeNil sets the value for ApplicationType to be an explicit nil

### UnsetApplicationType
`func (o *GetPermissions) UnsetApplicationType()`

UnsetApplicationType ensures that no value is present for ApplicationType, not even an explicit nil
### GetVersion

`func (o *GetPermissions) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetPermissions) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetPermissions) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetPermissions) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *GetPermissions) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *GetPermissions) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ThreeDSVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecommendedVersion** | Pointer to **NullableString** |  | [optional] 
**AvailableVersions** | Pointer to **[]string** |  | [optional] 
**EarliestAcsSupportedVersion** | Pointer to **NullableString** |  | [optional] 
**EarliestDsSupportedVersion** | Pointer to **NullableString** |  | [optional] 
**LatestAcsSupportedVersion** | Pointer to **NullableString** |  | [optional] 
**LatestDsSupportedVersion** | Pointer to **NullableString** |  | [optional] 
**AcsInformation** | Pointer to **[]string** |  | [optional] 

## Methods

### NewThreeDSVersion

`func NewThreeDSVersion() *ThreeDSVersion`

NewThreeDSVersion instantiates a new ThreeDSVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSVersionWithDefaults

`func NewThreeDSVersionWithDefaults() *ThreeDSVersion`

NewThreeDSVersionWithDefaults instantiates a new ThreeDSVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecommendedVersion

`func (o *ThreeDSVersion) GetRecommendedVersion() string`

GetRecommendedVersion returns the RecommendedVersion field if non-nil, zero value otherwise.

### GetRecommendedVersionOk

`func (o *ThreeDSVersion) GetRecommendedVersionOk() (*string, bool)`

GetRecommendedVersionOk returns a tuple with the RecommendedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedVersion

`func (o *ThreeDSVersion) SetRecommendedVersion(v string)`

SetRecommendedVersion sets RecommendedVersion field to given value.

### HasRecommendedVersion

`func (o *ThreeDSVersion) HasRecommendedVersion() bool`

HasRecommendedVersion returns a boolean if a field has been set.

### SetRecommendedVersionNil

`func (o *ThreeDSVersion) SetRecommendedVersionNil(b bool)`

 SetRecommendedVersionNil sets the value for RecommendedVersion to be an explicit nil

### UnsetRecommendedVersion
`func (o *ThreeDSVersion) UnsetRecommendedVersion()`

UnsetRecommendedVersion ensures that no value is present for RecommendedVersion, not even an explicit nil
### GetAvailableVersions

`func (o *ThreeDSVersion) GetAvailableVersions() []string`

GetAvailableVersions returns the AvailableVersions field if non-nil, zero value otherwise.

### GetAvailableVersionsOk

`func (o *ThreeDSVersion) GetAvailableVersionsOk() (*[]string, bool)`

GetAvailableVersionsOk returns a tuple with the AvailableVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableVersions

`func (o *ThreeDSVersion) SetAvailableVersions(v []string)`

SetAvailableVersions sets AvailableVersions field to given value.

### HasAvailableVersions

`func (o *ThreeDSVersion) HasAvailableVersions() bool`

HasAvailableVersions returns a boolean if a field has been set.

### SetAvailableVersionsNil

`func (o *ThreeDSVersion) SetAvailableVersionsNil(b bool)`

 SetAvailableVersionsNil sets the value for AvailableVersions to be an explicit nil

### UnsetAvailableVersions
`func (o *ThreeDSVersion) UnsetAvailableVersions()`

UnsetAvailableVersions ensures that no value is present for AvailableVersions, not even an explicit nil
### GetEarliestAcsSupportedVersion

`func (o *ThreeDSVersion) GetEarliestAcsSupportedVersion() string`

GetEarliestAcsSupportedVersion returns the EarliestAcsSupportedVersion field if non-nil, zero value otherwise.

### GetEarliestAcsSupportedVersionOk

`func (o *ThreeDSVersion) GetEarliestAcsSupportedVersionOk() (*string, bool)`

GetEarliestAcsSupportedVersionOk returns a tuple with the EarliestAcsSupportedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestAcsSupportedVersion

`func (o *ThreeDSVersion) SetEarliestAcsSupportedVersion(v string)`

SetEarliestAcsSupportedVersion sets EarliestAcsSupportedVersion field to given value.

### HasEarliestAcsSupportedVersion

`func (o *ThreeDSVersion) HasEarliestAcsSupportedVersion() bool`

HasEarliestAcsSupportedVersion returns a boolean if a field has been set.

### SetEarliestAcsSupportedVersionNil

`func (o *ThreeDSVersion) SetEarliestAcsSupportedVersionNil(b bool)`

 SetEarliestAcsSupportedVersionNil sets the value for EarliestAcsSupportedVersion to be an explicit nil

### UnsetEarliestAcsSupportedVersion
`func (o *ThreeDSVersion) UnsetEarliestAcsSupportedVersion()`

UnsetEarliestAcsSupportedVersion ensures that no value is present for EarliestAcsSupportedVersion, not even an explicit nil
### GetEarliestDsSupportedVersion

`func (o *ThreeDSVersion) GetEarliestDsSupportedVersion() string`

GetEarliestDsSupportedVersion returns the EarliestDsSupportedVersion field if non-nil, zero value otherwise.

### GetEarliestDsSupportedVersionOk

`func (o *ThreeDSVersion) GetEarliestDsSupportedVersionOk() (*string, bool)`

GetEarliestDsSupportedVersionOk returns a tuple with the EarliestDsSupportedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestDsSupportedVersion

`func (o *ThreeDSVersion) SetEarliestDsSupportedVersion(v string)`

SetEarliestDsSupportedVersion sets EarliestDsSupportedVersion field to given value.

### HasEarliestDsSupportedVersion

`func (o *ThreeDSVersion) HasEarliestDsSupportedVersion() bool`

HasEarliestDsSupportedVersion returns a boolean if a field has been set.

### SetEarliestDsSupportedVersionNil

`func (o *ThreeDSVersion) SetEarliestDsSupportedVersionNil(b bool)`

 SetEarliestDsSupportedVersionNil sets the value for EarliestDsSupportedVersion to be an explicit nil

### UnsetEarliestDsSupportedVersion
`func (o *ThreeDSVersion) UnsetEarliestDsSupportedVersion()`

UnsetEarliestDsSupportedVersion ensures that no value is present for EarliestDsSupportedVersion, not even an explicit nil
### GetLatestAcsSupportedVersion

`func (o *ThreeDSVersion) GetLatestAcsSupportedVersion() string`

GetLatestAcsSupportedVersion returns the LatestAcsSupportedVersion field if non-nil, zero value otherwise.

### GetLatestAcsSupportedVersionOk

`func (o *ThreeDSVersion) GetLatestAcsSupportedVersionOk() (*string, bool)`

GetLatestAcsSupportedVersionOk returns a tuple with the LatestAcsSupportedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAcsSupportedVersion

`func (o *ThreeDSVersion) SetLatestAcsSupportedVersion(v string)`

SetLatestAcsSupportedVersion sets LatestAcsSupportedVersion field to given value.

### HasLatestAcsSupportedVersion

`func (o *ThreeDSVersion) HasLatestAcsSupportedVersion() bool`

HasLatestAcsSupportedVersion returns a boolean if a field has been set.

### SetLatestAcsSupportedVersionNil

`func (o *ThreeDSVersion) SetLatestAcsSupportedVersionNil(b bool)`

 SetLatestAcsSupportedVersionNil sets the value for LatestAcsSupportedVersion to be an explicit nil

### UnsetLatestAcsSupportedVersion
`func (o *ThreeDSVersion) UnsetLatestAcsSupportedVersion()`

UnsetLatestAcsSupportedVersion ensures that no value is present for LatestAcsSupportedVersion, not even an explicit nil
### GetLatestDsSupportedVersion

`func (o *ThreeDSVersion) GetLatestDsSupportedVersion() string`

GetLatestDsSupportedVersion returns the LatestDsSupportedVersion field if non-nil, zero value otherwise.

### GetLatestDsSupportedVersionOk

`func (o *ThreeDSVersion) GetLatestDsSupportedVersionOk() (*string, bool)`

GetLatestDsSupportedVersionOk returns a tuple with the LatestDsSupportedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDsSupportedVersion

`func (o *ThreeDSVersion) SetLatestDsSupportedVersion(v string)`

SetLatestDsSupportedVersion sets LatestDsSupportedVersion field to given value.

### HasLatestDsSupportedVersion

`func (o *ThreeDSVersion) HasLatestDsSupportedVersion() bool`

HasLatestDsSupportedVersion returns a boolean if a field has been set.

### SetLatestDsSupportedVersionNil

`func (o *ThreeDSVersion) SetLatestDsSupportedVersionNil(b bool)`

 SetLatestDsSupportedVersionNil sets the value for LatestDsSupportedVersion to be an explicit nil

### UnsetLatestDsSupportedVersion
`func (o *ThreeDSVersion) UnsetLatestDsSupportedVersion()`

UnsetLatestDsSupportedVersion ensures that no value is present for LatestDsSupportedVersion, not even an explicit nil
### GetAcsInformation

`func (o *ThreeDSVersion) GetAcsInformation() []string`

GetAcsInformation returns the AcsInformation field if non-nil, zero value otherwise.

### GetAcsInformationOk

`func (o *ThreeDSVersion) GetAcsInformationOk() (*[]string, bool)`

GetAcsInformationOk returns a tuple with the AcsInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsInformation

`func (o *ThreeDSVersion) SetAcsInformation(v []string)`

SetAcsInformation sets AcsInformation field to given value.

### HasAcsInformation

`func (o *ThreeDSVersion) HasAcsInformation() bool`

HasAcsInformation returns a boolean if a field has been set.

### SetAcsInformationNil

`func (o *ThreeDSVersion) SetAcsInformationNil(b bool)`

 SetAcsInformationNil sets the value for AcsInformation to be an explicit nil

### UnsetAcsInformation
`func (o *ThreeDSVersion) UnsetAcsInformation()`

UnsetAcsInformation ensures that no value is present for AcsInformation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



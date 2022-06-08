# Privacy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classification** | Pointer to **NullableString** |  | [optional] 
**ImpactLevel** | Pointer to **NullableString** |  | [optional] 
**RestrictionPolicy** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPrivacy

`func NewPrivacy() *Privacy`

NewPrivacy instantiates a new Privacy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivacyWithDefaults

`func NewPrivacyWithDefaults() *Privacy`

NewPrivacyWithDefaults instantiates a new Privacy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassification

`func (o *Privacy) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *Privacy) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *Privacy) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *Privacy) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassificationNil

`func (o *Privacy) SetClassificationNil(b bool)`

 SetClassificationNil sets the value for Classification to be an explicit nil

### UnsetClassification
`func (o *Privacy) UnsetClassification()`

UnsetClassification ensures that no value is present for Classification, not even an explicit nil
### GetImpactLevel

`func (o *Privacy) GetImpactLevel() string`

GetImpactLevel returns the ImpactLevel field if non-nil, zero value otherwise.

### GetImpactLevelOk

`func (o *Privacy) GetImpactLevelOk() (*string, bool)`

GetImpactLevelOk returns a tuple with the ImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactLevel

`func (o *Privacy) SetImpactLevel(v string)`

SetImpactLevel sets ImpactLevel field to given value.

### HasImpactLevel

`func (o *Privacy) HasImpactLevel() bool`

HasImpactLevel returns a boolean if a field has been set.

### SetImpactLevelNil

`func (o *Privacy) SetImpactLevelNil(b bool)`

 SetImpactLevelNil sets the value for ImpactLevel to be an explicit nil

### UnsetImpactLevel
`func (o *Privacy) UnsetImpactLevel()`

UnsetImpactLevel ensures that no value is present for ImpactLevel, not even an explicit nil
### GetRestrictionPolicy

`func (o *Privacy) GetRestrictionPolicy() string`

GetRestrictionPolicy returns the RestrictionPolicy field if non-nil, zero value otherwise.

### GetRestrictionPolicyOk

`func (o *Privacy) GetRestrictionPolicyOk() (*string, bool)`

GetRestrictionPolicyOk returns a tuple with the RestrictionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionPolicy

`func (o *Privacy) SetRestrictionPolicy(v string)`

SetRestrictionPolicy sets RestrictionPolicy field to given value.

### HasRestrictionPolicy

`func (o *Privacy) HasRestrictionPolicy() bool`

HasRestrictionPolicy returns a boolean if a field has been set.

### SetRestrictionPolicyNil

`func (o *Privacy) SetRestrictionPolicyNil(b bool)`

 SetRestrictionPolicyNil sets the value for RestrictionPolicy to be an explicit nil

### UnsetRestrictionPolicy
`func (o *Privacy) UnsetRestrictionPolicy()`

UnsetRestrictionPolicy ensures that no value is present for RestrictionPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



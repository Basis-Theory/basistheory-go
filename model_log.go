/*
Basis Theory API

## Getting Started * Sign-in to [Basis Theory](https://basistheory.com) and go to [Applications](https://portal.basistheory.com/applications) * Create a Basis Theory Private Application * All permissions should be selected * Paste the API Key into the `BT-API-KEY` variable

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package basistheory

import (
	"encoding/json"
	"time"
)

// checks if the Log type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Log{}

// Log struct for Log
type Log struct {
	Id         NullableString `json:"id,omitempty"`
	TenantId   *string        `json:"tenant_id,omitempty"`
	ActorId    NullableString `json:"actor_id,omitempty"`
	ActorType  NullableString `json:"actor_type,omitempty"`
	EntityType NullableString `json:"entity_type,omitempty"`
	EntityId   NullableString `json:"entity_id,omitempty"`
	Operation  NullableString `json:"operation,omitempty"`
	Message    NullableString `json:"message,omitempty"`
	CreatedAt  NullableTime   `json:"created_at,omitempty"`
}

// NewLog instantiates a new Log object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLog() *Log {
	this := Log{}
	return &this
}

// NewLogWithDefaults instantiates a new Log object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogWithDefaults() *Log {
	this := Log{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field is not nil.
func (o *Log) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *Log) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *Log) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Log) UnsetId() {
	o.Id.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *Log) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field is not nil.
func (o *Log) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *Log) SetTenantId(v string) {
	o.TenantId = &v
}

// GetActorId returns the ActorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetActorId() string {
	if o == nil || IsNil(o.ActorId.Get()) {
		var ret string
		return ret
	}
	return *o.ActorId.Get()
}

// GetActorIdOk returns a tuple with the ActorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetActorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActorId.Get(), o.ActorId.IsSet()
}

// HasActorId returns a boolean if a field is not nil.
func (o *Log) HasActorId() bool {
	if o != nil && !IsNil(o.ActorId) {
		return true
	}

	return false
}

// SetActorId gets a reference to the given NullableString and assigns it to the ActorId field.
func (o *Log) SetActorId(v string) {
	o.ActorId.Set(&v)
}

// SetActorIdNil sets the value for ActorId to be an explicit nil
func (o *Log) SetActorIdNil() {
	o.ActorId.Set(nil)
}

// UnsetActorId ensures that no value is present for ActorId, not even an explicit nil
func (o *Log) UnsetActorId() {
	o.ActorId.Unset()
}

// GetActorType returns the ActorType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetActorType() string {
	if o == nil || IsNil(o.ActorType.Get()) {
		var ret string
		return ret
	}
	return *o.ActorType.Get()
}

// GetActorTypeOk returns a tuple with the ActorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetActorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActorType.Get(), o.ActorType.IsSet()
}

// HasActorType returns a boolean if a field is not nil.
func (o *Log) HasActorType() bool {
	if o != nil && !IsNil(o.ActorType) {
		return true
	}

	return false
}

// SetActorType gets a reference to the given NullableString and assigns it to the ActorType field.
func (o *Log) SetActorType(v string) {
	o.ActorType.Set(&v)
}

// SetActorTypeNil sets the value for ActorType to be an explicit nil
func (o *Log) SetActorTypeNil() {
	o.ActorType.Set(nil)
}

// UnsetActorType ensures that no value is present for ActorType, not even an explicit nil
func (o *Log) UnsetActorType() {
	o.ActorType.Unset()
}

// GetEntityType returns the EntityType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetEntityType() string {
	if o == nil || IsNil(o.EntityType.Get()) {
		var ret string
		return ret
	}
	return *o.EntityType.Get()
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntityType.Get(), o.EntityType.IsSet()
}

// HasEntityType returns a boolean if a field is not nil.
func (o *Log) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given NullableString and assigns it to the EntityType field.
func (o *Log) SetEntityType(v string) {
	o.EntityType.Set(&v)
}

// SetEntityTypeNil sets the value for EntityType to be an explicit nil
func (o *Log) SetEntityTypeNil() {
	o.EntityType.Set(nil)
}

// UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
func (o *Log) UnsetEntityType() {
	o.EntityType.Unset()
}

// GetEntityId returns the EntityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetEntityId() string {
	if o == nil || IsNil(o.EntityId.Get()) {
		var ret string
		return ret
	}
	return *o.EntityId.Get()
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntityId.Get(), o.EntityId.IsSet()
}

// HasEntityId returns a boolean if a field is not nil.
func (o *Log) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given NullableString and assigns it to the EntityId field.
func (o *Log) SetEntityId(v string) {
	o.EntityId.Set(&v)
}

// SetEntityIdNil sets the value for EntityId to be an explicit nil
func (o *Log) SetEntityIdNil() {
	o.EntityId.Set(nil)
}

// UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
func (o *Log) UnsetEntityId() {
	o.EntityId.Unset()
}

// GetOperation returns the Operation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetOperation() string {
	if o == nil || IsNil(o.Operation.Get()) {
		var ret string
		return ret
	}
	return *o.Operation.Get()
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operation.Get(), o.Operation.IsSet()
}

// HasOperation returns a boolean if a field is not nil.
func (o *Log) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given NullableString and assigns it to the Operation field.
func (o *Log) SetOperation(v string) {
	o.Operation.Set(&v)
}

// SetOperationNil sets the value for Operation to be an explicit nil
func (o *Log) SetOperationNil() {
	o.Operation.Set(nil)
}

// UnsetOperation ensures that no value is present for Operation, not even an explicit nil
func (o *Log) UnsetOperation() {
	o.Operation.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field is not nil.
func (o *Log) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *Log) SetMessage(v string) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil
func (o *Log) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *Log) UnsetMessage() {
	o.Message.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field is not nil.
func (o *Log) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *Log) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *Log) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *Log) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

func (o Log) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Log) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenant_id"] = o.TenantId
	}
	if o.ActorId.IsSet() {
		toSerialize["actor_id"] = o.ActorId.Get()
	}
	if o.ActorType.IsSet() {
		toSerialize["actor_type"] = o.ActorType.Get()
	}
	if o.EntityType.IsSet() {
		toSerialize["entity_type"] = o.EntityType.Get()
	}
	if o.EntityId.IsSet() {
		toSerialize["entity_id"] = o.EntityId.Get()
	}
	if o.Operation.IsSet() {
		toSerialize["operation"] = o.Operation.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	return toSerialize, nil
}

type NullableLog struct {
	value *Log
	isSet bool
}

func (v NullableLog) Get() *Log {
	return v.value
}

func (v *NullableLog) Set(val *Log) {
	v.value = val
	v.isSet = true
}

func (v NullableLog) IsSet() bool {
	return v.isSet
}

func (v *NullableLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLog(val *Log) *NullableLog {
	return &NullableLog{value: val, isSet: true}
}

func (v NullableLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Redfish OAPI specification

Partial Redfish OAPI specification for a limited client

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ComputerSystemReset struct for ComputerSystemReset
type ComputerSystemReset struct {
	// The unique identifier for a resource.
	Target *string `json:"target,omitempty"`
	ResetTypeRedfishAllowableValues []ResetType `json:"ResetType@Redfish.AllowableValues,omitempty"`
}

// NewComputerSystemReset instantiates a new ComputerSystemReset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerSystemReset() *ComputerSystemReset {
	this := ComputerSystemReset{}
	return &this
}

// NewComputerSystemResetWithDefaults instantiates a new ComputerSystemReset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerSystemResetWithDefaults() *ComputerSystemReset {
	this := ComputerSystemReset{}
	return &this
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ComputerSystemReset) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerSystemReset) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ComputerSystemReset) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *ComputerSystemReset) SetTarget(v string) {
	o.Target = &v
}

// GetResetTypeRedfishAllowableValues returns the ResetTypeRedfishAllowableValues field value if set, zero value otherwise.
func (o *ComputerSystemReset) GetResetTypeRedfishAllowableValues() []ResetType {
	if o == nil || o.ResetTypeRedfishAllowableValues == nil {
		var ret []ResetType
		return ret
	}
	return o.ResetTypeRedfishAllowableValues
}

// GetResetTypeRedfishAllowableValuesOk returns a tuple with the ResetTypeRedfishAllowableValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerSystemReset) GetResetTypeRedfishAllowableValuesOk() ([]ResetType, bool) {
	if o == nil || o.ResetTypeRedfishAllowableValues == nil {
		return nil, false
	}
	return o.ResetTypeRedfishAllowableValues, true
}

// HasResetTypeRedfishAllowableValues returns a boolean if a field has been set.
func (o *ComputerSystemReset) HasResetTypeRedfishAllowableValues() bool {
	if o != nil && o.ResetTypeRedfishAllowableValues != nil {
		return true
	}

	return false
}

// SetResetTypeRedfishAllowableValues gets a reference to the given []ResetType and assigns it to the ResetTypeRedfishAllowableValues field.
func (o *ComputerSystemReset) SetResetTypeRedfishAllowableValues(v []ResetType) {
	o.ResetTypeRedfishAllowableValues = v
}

func (o ComputerSystemReset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.ResetTypeRedfishAllowableValues != nil {
		toSerialize["ResetType@Redfish.AllowableValues"] = o.ResetTypeRedfishAllowableValues
	}
	return json.Marshal(toSerialize)
}

type NullableComputerSystemReset struct {
	value *ComputerSystemReset
	isSet bool
}

func (v NullableComputerSystemReset) Get() *ComputerSystemReset {
	return v.value
}

func (v *NullableComputerSystemReset) Set(val *ComputerSystemReset) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerSystemReset) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerSystemReset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerSystemReset(val *ComputerSystemReset) *NullableComputerSystemReset {
	return &NullableComputerSystemReset{value: val, isSet: true}
}

func (v NullableComputerSystemReset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerSystemReset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



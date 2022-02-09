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

// Status struct for Status
type Status struct {
	Health *Health `json:"Health,omitempty"`
	HealthRollup *Health `json:"HealthRollup,omitempty"`
	State *State `json:"State,omitempty"`
}

// NewStatus instantiates a new Status object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatus() *Status {
	this := Status{}
	return &this
}

// NewStatusWithDefaults instantiates a new Status object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusWithDefaults() *Status {
	this := Status{}
	return &this
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *Status) GetHealth() Health {
	if o == nil || o.Health == nil {
		var ret Health
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Status) GetHealthOk() (*Health, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *Status) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given Health and assigns it to the Health field.
func (o *Status) SetHealth(v Health) {
	o.Health = &v
}

// GetHealthRollup returns the HealthRollup field value if set, zero value otherwise.
func (o *Status) GetHealthRollup() Health {
	if o == nil || o.HealthRollup == nil {
		var ret Health
		return ret
	}
	return *o.HealthRollup
}

// GetHealthRollupOk returns a tuple with the HealthRollup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Status) GetHealthRollupOk() (*Health, bool) {
	if o == nil || o.HealthRollup == nil {
		return nil, false
	}
	return o.HealthRollup, true
}

// HasHealthRollup returns a boolean if a field has been set.
func (o *Status) HasHealthRollup() bool {
	if o != nil && o.HealthRollup != nil {
		return true
	}

	return false
}

// SetHealthRollup gets a reference to the given Health and assigns it to the HealthRollup field.
func (o *Status) SetHealthRollup(v Health) {
	o.HealthRollup = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Status) GetState() State {
	if o == nil || o.State == nil {
		var ret State
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Status) GetStateOk() (*State, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Status) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given State and assigns it to the State field.
func (o *Status) SetState(v State) {
	o.State = &v
}

func (o Status) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Health != nil {
		toSerialize["Health"] = o.Health
	}
	if o.HealthRollup != nil {
		toSerialize["HealthRollup"] = o.HealthRollup
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableStatus struct {
	value *Status
	isSet bool
}

func (v NullableStatus) Get() *Status {
	return v.value
}

func (v *NullableStatus) Set(val *Status) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus(val *Status) *NullableStatus {
	return &NullableStatus{value: val, isSet: true}
}

func (v NullableStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



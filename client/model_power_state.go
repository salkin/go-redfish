/*
Redfish OAPI specification

Partial Redfish OAPI specification for a limited client

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// PowerState the model 'PowerState'
type PowerState string

// List of PowerState
const (
	POWERSTATE_TRUE PowerState = "true"
	POWERSTATE_FALSE PowerState = "false"
	POWERSTATE_POWERING_ON PowerState = "PoweringOn"
	POWERSTATE_POWERING_OFF PowerState = "PoweringOff"
)

// All allowed values of PowerState enum
var AllowedPowerStateEnumValues = []PowerState{
	"true",
	"false",
	"PoweringOn",
	"PoweringOff",
}

func (v *PowerState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PowerState(value)
	for _, existing := range AllowedPowerStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PowerState", value)
}

// NewPowerStateFromValue returns a pointer to a valid PowerState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerStateFromValue(v string) (*PowerState, error) {
	ev := PowerState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PowerState: valid values are %v", v, AllowedPowerStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerState) IsValid() bool {
	for _, existing := range AllowedPowerStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerState value
func (v PowerState) Ptr() *PowerState {
	return &v
}

type NullablePowerState struct {
	value *PowerState
	isSet bool
}

func (v NullablePowerState) Get() *PowerState {
	return v.value
}

func (v *NullablePowerState) Set(val *PowerState) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerState) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerState(val *PowerState) *NullablePowerState {
	return &NullablePowerState{value: val, isSet: true}
}

func (v NullablePowerState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// TaskState the model 'TaskState'
type TaskState string

// List of TaskState
const (
	TASKSTATE_NEW TaskState = "New"
	TASKSTATE_STARTING TaskState = "Starting"
	TASKSTATE_RUNNING TaskState = "Running"
	TASKSTATE_SUSPENDED TaskState = "Suspended"
	TASKSTATE_INTERRUPTED TaskState = "Interrupted"
	TASKSTATE_PENDING TaskState = "Pending"
	TASKSTATE_STOPPING TaskState = "Stopping"
	TASKSTATE_COMPLETED TaskState = "Completed"
	TASKSTATE_KILLED TaskState = "Killed"
	TASKSTATE_EXCEPTION TaskState = "Exception"
	TASKSTATE_SERVICE TaskState = "Service"
	TASKSTATE_CANCELLING TaskState = "Cancelling"
	TASKSTATE_CANCELLED TaskState = "Cancelled"
)

// All allowed values of TaskState enum
var AllowedTaskStateEnumValues = []TaskState{
	"New",
	"Starting",
	"Running",
	"Suspended",
	"Interrupted",
	"Pending",
	"Stopping",
	"Completed",
	"Killed",
	"Exception",
	"Service",
	"Cancelling",
	"Cancelled",
}

func (v *TaskState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TaskState(value)
	for _, existing := range AllowedTaskStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TaskState", value)
}

// NewTaskStateFromValue returns a pointer to a valid TaskState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTaskStateFromValue(v string) (*TaskState, error) {
	ev := TaskState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TaskState: valid values are %v", v, AllowedTaskStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TaskState) IsValid() bool {
	for _, existing := range AllowedTaskStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaskState value
func (v TaskState) Ptr() *TaskState {
	return &v
}

type NullableTaskState struct {
	value *TaskState
	isSet bool
}

func (v NullableTaskState) Get() *TaskState {
	return v.value
}

func (v *NullableTaskState) Set(val *TaskState) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskState) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskState(val *TaskState) *NullableTaskState {
	return &NullableTaskState{value: val, isSet: true}
}

func (v NullableTaskState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


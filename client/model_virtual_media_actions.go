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

// VirtualMediaActions struct for VirtualMediaActions
type VirtualMediaActions struct {
	VirtualMediaEjectMedia *VirtualMediaActionsVirtualMediaEjectMedia `json:"#VirtualMedia.EjectMedia,omitempty"`
	VirtualMediaInsertMedia *VirtualMediaActionsVirtualMediaEjectMedia `json:"#VirtualMedia.InsertMedia,omitempty"`
}

// NewVirtualMediaActions instantiates a new VirtualMediaActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualMediaActions() *VirtualMediaActions {
	this := VirtualMediaActions{}
	return &this
}

// NewVirtualMediaActionsWithDefaults instantiates a new VirtualMediaActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualMediaActionsWithDefaults() *VirtualMediaActions {
	this := VirtualMediaActions{}
	return &this
}

// GetVirtualMediaEjectMedia returns the VirtualMediaEjectMedia field value if set, zero value otherwise.
func (o *VirtualMediaActions) GetVirtualMediaEjectMedia() VirtualMediaActionsVirtualMediaEjectMedia {
	if o == nil || o.VirtualMediaEjectMedia == nil {
		var ret VirtualMediaActionsVirtualMediaEjectMedia
		return ret
	}
	return *o.VirtualMediaEjectMedia
}

// GetVirtualMediaEjectMediaOk returns a tuple with the VirtualMediaEjectMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMediaActions) GetVirtualMediaEjectMediaOk() (*VirtualMediaActionsVirtualMediaEjectMedia, bool) {
	if o == nil || o.VirtualMediaEjectMedia == nil {
		return nil, false
	}
	return o.VirtualMediaEjectMedia, true
}

// HasVirtualMediaEjectMedia returns a boolean if a field has been set.
func (o *VirtualMediaActions) HasVirtualMediaEjectMedia() bool {
	if o != nil && o.VirtualMediaEjectMedia != nil {
		return true
	}

	return false
}

// SetVirtualMediaEjectMedia gets a reference to the given VirtualMediaActionsVirtualMediaEjectMedia and assigns it to the VirtualMediaEjectMedia field.
func (o *VirtualMediaActions) SetVirtualMediaEjectMedia(v VirtualMediaActionsVirtualMediaEjectMedia) {
	o.VirtualMediaEjectMedia = &v
}

// GetVirtualMediaInsertMedia returns the VirtualMediaInsertMedia field value if set, zero value otherwise.
func (o *VirtualMediaActions) GetVirtualMediaInsertMedia() VirtualMediaActionsVirtualMediaEjectMedia {
	if o == nil || o.VirtualMediaInsertMedia == nil {
		var ret VirtualMediaActionsVirtualMediaEjectMedia
		return ret
	}
	return *o.VirtualMediaInsertMedia
}

// GetVirtualMediaInsertMediaOk returns a tuple with the VirtualMediaInsertMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMediaActions) GetVirtualMediaInsertMediaOk() (*VirtualMediaActionsVirtualMediaEjectMedia, bool) {
	if o == nil || o.VirtualMediaInsertMedia == nil {
		return nil, false
	}
	return o.VirtualMediaInsertMedia, true
}

// HasVirtualMediaInsertMedia returns a boolean if a field has been set.
func (o *VirtualMediaActions) HasVirtualMediaInsertMedia() bool {
	if o != nil && o.VirtualMediaInsertMedia != nil {
		return true
	}

	return false
}

// SetVirtualMediaInsertMedia gets a reference to the given VirtualMediaActionsVirtualMediaEjectMedia and assigns it to the VirtualMediaInsertMedia field.
func (o *VirtualMediaActions) SetVirtualMediaInsertMedia(v VirtualMediaActionsVirtualMediaEjectMedia) {
	o.VirtualMediaInsertMedia = &v
}

func (o VirtualMediaActions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VirtualMediaEjectMedia != nil {
		toSerialize["#VirtualMedia.EjectMedia"] = o.VirtualMediaEjectMedia
	}
	if o.VirtualMediaInsertMedia != nil {
		toSerialize["#VirtualMedia.InsertMedia"] = o.VirtualMediaInsertMedia
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualMediaActions struct {
	value *VirtualMediaActions
	isSet bool
}

func (v NullableVirtualMediaActions) Get() *VirtualMediaActions {
	return v.value
}

func (v *NullableVirtualMediaActions) Set(val *VirtualMediaActions) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMediaActions) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMediaActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMediaActions(val *VirtualMediaActions) *NullableVirtualMediaActions {
	return &NullableVirtualMediaActions{value: val, isSet: true}
}

func (v NullableVirtualMediaActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMediaActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



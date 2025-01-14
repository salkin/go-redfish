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

// Root Root redfish path.
type Root struct {
	// The name of the resource.
	Id *string `json:"Id,omitempty"`
	// The name of the resource.
	Name string `json:"Name"`
	// redfish version
	RedfishVersion *string `json:"RedfishVersion,omitempty"`
	UUID *string `json:"UUID,omitempty"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// redfish copyright
	RedfishCopyright *string `json:"@Redfish.Copyright,omitempty"`
	Systems *IdRef `json:"Systems,omitempty"`
	Managers *IdRef `json:"Managers,omitempty"`
}

// NewRoot instantiates a new Root object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoot(name string, odataType string, odataId string) *Root {
	this := Root{}
	this.Name = name
	this.OdataType = odataType
	this.OdataId = odataId
	return &this
}

// NewRootWithDefaults instantiates a new Root object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRootWithDefaults() *Root {
	this := Root{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Root) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Root) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Root) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Root) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Root) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Root) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Root) SetName(v string) {
	o.Name = v
}

// GetRedfishVersion returns the RedfishVersion field value if set, zero value otherwise.
func (o *Root) GetRedfishVersion() string {
	if o == nil || o.RedfishVersion == nil {
		var ret string
		return ret
	}
	return *o.RedfishVersion
}

// GetRedfishVersionOk returns a tuple with the RedfishVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Root) GetRedfishVersionOk() (*string, bool) {
	if o == nil || o.RedfishVersion == nil {
		return nil, false
	}
	return o.RedfishVersion, true
}

// HasRedfishVersion returns a boolean if a field has been set.
func (o *Root) HasRedfishVersion() bool {
	if o != nil && o.RedfishVersion != nil {
		return true
	}

	return false
}

// SetRedfishVersion gets a reference to the given string and assigns it to the RedfishVersion field.
func (o *Root) SetRedfishVersion(v string) {
	o.RedfishVersion = &v
}

// GetUUID returns the UUID field value if set, zero value otherwise.
func (o *Root) GetUUID() string {
	if o == nil || o.UUID == nil {
		var ret string
		return ret
	}
	return *o.UUID
}

// GetUUIDOk returns a tuple with the UUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Root) GetUUIDOk() (*string, bool) {
	if o == nil || o.UUID == nil {
		return nil, false
	}
	return o.UUID, true
}

// HasUUID returns a boolean if a field has been set.
func (o *Root) HasUUID() bool {
	if o != nil && o.UUID != nil {
		return true
	}

	return false
}

// SetUUID gets a reference to the given string and assigns it to the UUID field.
func (o *Root) SetUUID(v string) {
	o.UUID = &v
}

// GetOdataType returns the OdataType field value
func (o *Root) GetOdataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataType
}

// GetOdataTypeOk returns a tuple with the OdataType field value
// and a boolean to check if the value has been set.
func (o *Root) GetOdataTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OdataType, true
}

// SetOdataType sets field value
func (o *Root) SetOdataType(v string) {
	o.OdataType = v
}

// GetOdataId returns the OdataId field value
func (o *Root) GetOdataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataId
}

// GetOdataIdOk returns a tuple with the OdataId field value
// and a boolean to check if the value has been set.
func (o *Root) GetOdataIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OdataId, true
}

// SetOdataId sets field value
func (o *Root) SetOdataId(v string) {
	o.OdataId = v
}

// GetRedfishCopyright returns the RedfishCopyright field value if set, zero value otherwise.
func (o *Root) GetRedfishCopyright() string {
	if o == nil || o.RedfishCopyright == nil {
		var ret string
		return ret
	}
	return *o.RedfishCopyright
}

// GetRedfishCopyrightOk returns a tuple with the RedfishCopyright field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Root) GetRedfishCopyrightOk() (*string, bool) {
	if o == nil || o.RedfishCopyright == nil {
		return nil, false
	}
	return o.RedfishCopyright, true
}

// HasRedfishCopyright returns a boolean if a field has been set.
func (o *Root) HasRedfishCopyright() bool {
	if o != nil && o.RedfishCopyright != nil {
		return true
	}

	return false
}

// SetRedfishCopyright gets a reference to the given string and assigns it to the RedfishCopyright field.
func (o *Root) SetRedfishCopyright(v string) {
	o.RedfishCopyright = &v
}

// GetSystems returns the Systems field value if set, zero value otherwise.
func (o *Root) GetSystems() IdRef {
	if o == nil || o.Systems == nil {
		var ret IdRef
		return ret
	}
	return *o.Systems
}

// GetSystemsOk returns a tuple with the Systems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Root) GetSystemsOk() (*IdRef, bool) {
	if o == nil || o.Systems == nil {
		return nil, false
	}
	return o.Systems, true
}

// HasSystems returns a boolean if a field has been set.
func (o *Root) HasSystems() bool {
	if o != nil && o.Systems != nil {
		return true
	}

	return false
}

// SetSystems gets a reference to the given IdRef and assigns it to the Systems field.
func (o *Root) SetSystems(v IdRef) {
	o.Systems = &v
}

// GetManagers returns the Managers field value if set, zero value otherwise.
func (o *Root) GetManagers() IdRef {
	if o == nil || o.Managers == nil {
		var ret IdRef
		return ret
	}
	return *o.Managers
}

// GetManagersOk returns a tuple with the Managers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Root) GetManagersOk() (*IdRef, bool) {
	if o == nil || o.Managers == nil {
		return nil, false
	}
	return o.Managers, true
}

// HasManagers returns a boolean if a field has been set.
func (o *Root) HasManagers() bool {
	if o != nil && o.Managers != nil {
		return true
	}

	return false
}

// SetManagers gets a reference to the given IdRef and assigns it to the Managers field.
func (o *Root) SetManagers(v IdRef) {
	o.Managers = &v
}

func (o Root) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if true {
		toSerialize["Name"] = o.Name
	}
	if o.RedfishVersion != nil {
		toSerialize["RedfishVersion"] = o.RedfishVersion
	}
	if o.UUID != nil {
		toSerialize["UUID"] = o.UUID
	}
	if true {
		toSerialize["@odata.type"] = o.OdataType
	}
	if true {
		toSerialize["@odata.id"] = o.OdataId
	}
	if o.RedfishCopyright != nil {
		toSerialize["@Redfish.Copyright"] = o.RedfishCopyright
	}
	if o.Systems != nil {
		toSerialize["Systems"] = o.Systems
	}
	if o.Managers != nil {
		toSerialize["Managers"] = o.Managers
	}
	return json.Marshal(toSerialize)
}

type NullableRoot struct {
	value *Root
	isSet bool
}

func (v NullableRoot) Get() *Root {
	return v.value
}

func (v *NullableRoot) Set(val *Root) {
	v.value = val
	v.isSet = true
}

func (v NullableRoot) IsSet() bool {
	return v.isSet
}

func (v *NullableRoot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoot(val *Root) *NullableRoot {
	return &NullableRoot{value: val, isSet: true}
}

func (v NullableRoot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



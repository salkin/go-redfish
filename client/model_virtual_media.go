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

// VirtualMedia Redfish virtual media resource for manager.
type VirtualMedia struct {
	// The name of the resource.
	Id *string `json:"Id,omitempty"`
	// The name of the resource.
	Name string `json:"Name"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The OData description of a payload.
	OdataContext *string `json:"@odata.context,omitempty"`
	// redfish copyright
	RedfishCopyright *string `json:"@Redfish.Copyright,omitempty"`
	// description
	Description NullableString `json:"Description,omitempty"`
	Image NullableString `json:"Image,omitempty"`
	ImageName NullableString `json:"ImageName,omitempty"`
	Inserted NullableBool `json:"Inserted,omitempty"`
	ConnectedVia *ConnectedVia `json:"ConnectedVia,omitempty"`
	MediaTypes []string `json:"MediaTypes,omitempty"`
	WriteProtected NullableBool `json:"WriteProtected,omitempty"`
	UserName NullableString `json:"UserName,omitempty"`
	Password NullableString `json:"Password,omitempty"`
	TransferMethod *TransferMethod `json:"TransferMethod,omitempty"`
	TransferProtocolType *TransferProtocolType `json:"TransferProtocolType,omitempty"`
	Actions *VirtualMediaActions `json:"Actions,omitempty"`
}

// NewVirtualMedia instantiates a new VirtualMedia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualMedia(name string, odataType string, odataId string) *VirtualMedia {
	this := VirtualMedia{}
	this.Name = name
	this.OdataType = odataType
	this.OdataId = odataId
	return &this
}

// NewVirtualMediaWithDefaults instantiates a new VirtualMedia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualMediaWithDefaults() *VirtualMedia {
	this := VirtualMedia{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VirtualMedia) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMedia) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VirtualMedia) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VirtualMedia) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *VirtualMedia) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VirtualMedia) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VirtualMedia) SetName(v string) {
	o.Name = v
}

// GetOdataType returns the OdataType field value
func (o *VirtualMedia) GetOdataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataType
}

// GetOdataTypeOk returns a tuple with the OdataType field value
// and a boolean to check if the value has been set.
func (o *VirtualMedia) GetOdataTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OdataType, true
}

// SetOdataType sets field value
func (o *VirtualMedia) SetOdataType(v string) {
	o.OdataType = v
}

// GetOdataId returns the OdataId field value
func (o *VirtualMedia) GetOdataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataId
}

// GetOdataIdOk returns a tuple with the OdataId field value
// and a boolean to check if the value has been set.
func (o *VirtualMedia) GetOdataIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OdataId, true
}

// SetOdataId sets field value
func (o *VirtualMedia) SetOdataId(v string) {
	o.OdataId = v
}

// GetOdataContext returns the OdataContext field value if set, zero value otherwise.
func (o *VirtualMedia) GetOdataContext() string {
	if o == nil || o.OdataContext == nil {
		var ret string
		return ret
	}
	return *o.OdataContext
}

// GetOdataContextOk returns a tuple with the OdataContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMedia) GetOdataContextOk() (*string, bool) {
	if o == nil || o.OdataContext == nil {
		return nil, false
	}
	return o.OdataContext, true
}

// HasOdataContext returns a boolean if a field has been set.
func (o *VirtualMedia) HasOdataContext() bool {
	if o != nil && o.OdataContext != nil {
		return true
	}

	return false
}

// SetOdataContext gets a reference to the given string and assigns it to the OdataContext field.
func (o *VirtualMedia) SetOdataContext(v string) {
	o.OdataContext = &v
}

// GetRedfishCopyright returns the RedfishCopyright field value if set, zero value otherwise.
func (o *VirtualMedia) GetRedfishCopyright() string {
	if o == nil || o.RedfishCopyright == nil {
		var ret string
		return ret
	}
	return *o.RedfishCopyright
}

// GetRedfishCopyrightOk returns a tuple with the RedfishCopyright field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMedia) GetRedfishCopyrightOk() (*string, bool) {
	if o == nil || o.RedfishCopyright == nil {
		return nil, false
	}
	return o.RedfishCopyright, true
}

// HasRedfishCopyright returns a boolean if a field has been set.
func (o *VirtualMedia) HasRedfishCopyright() bool {
	if o != nil && o.RedfishCopyright != nil {
		return true
	}

	return false
}

// SetRedfishCopyright gets a reference to the given string and assigns it to the RedfishCopyright field.
func (o *VirtualMedia) SetRedfishCopyright(v string) {
	o.RedfishCopyright = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMedia) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMedia) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *VirtualMedia) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *VirtualMedia) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *VirtualMedia) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *VirtualMedia) UnsetDescription() {
	o.Description.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMedia) GetImage() string {
	if o == nil || o.Image.Get() == nil {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMedia) GetImageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *VirtualMedia) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *VirtualMedia) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *VirtualMedia) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *VirtualMedia) UnsetImage() {
	o.Image.Unset()
}

// GetImageName returns the ImageName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMedia) GetImageName() string {
	if o == nil || o.ImageName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ImageName.Get()
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMedia) GetImageNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ImageName.Get(), o.ImageName.IsSet()
}

// HasImageName returns a boolean if a field has been set.
func (o *VirtualMedia) HasImageName() bool {
	if o != nil && o.ImageName.IsSet() {
		return true
	}

	return false
}

// SetImageName gets a reference to the given NullableString and assigns it to the ImageName field.
func (o *VirtualMedia) SetImageName(v string) {
	o.ImageName.Set(&v)
}
// SetImageNameNil sets the value for ImageName to be an explicit nil
func (o *VirtualMedia) SetImageNameNil() {
	o.ImageName.Set(nil)
}

// UnsetImageName ensures that no value is present for ImageName, not even an explicit nil
func (o *VirtualMedia) UnsetImageName() {
	o.ImageName.Unset()
}

// GetInserted returns the Inserted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMedia) GetInserted() bool {
	if o == nil || o.Inserted.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Inserted.Get()
}

// GetInsertedOk returns a tuple with the Inserted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMedia) GetInsertedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Inserted.Get(), o.Inserted.IsSet()
}

// HasInserted returns a boolean if a field has been set.
func (o *VirtualMedia) HasInserted() bool {
	if o != nil && o.Inserted.IsSet() {
		return true
	}

	return false
}

// SetInserted gets a reference to the given NullableBool and assigns it to the Inserted field.
func (o *VirtualMedia) SetInserted(v bool) {
	o.Inserted.Set(&v)
}
// SetInsertedNil sets the value for Inserted to be an explicit nil
func (o *VirtualMedia) SetInsertedNil() {
	o.Inserted.Set(nil)
}

// UnsetInserted ensures that no value is present for Inserted, not even an explicit nil
func (o *VirtualMedia) UnsetInserted() {
	o.Inserted.Unset()
}

// GetConnectedVia returns the ConnectedVia field value if set, zero value otherwise.
func (o *VirtualMedia) GetConnectedVia() ConnectedVia {
	if o == nil || o.ConnectedVia == nil {
		var ret ConnectedVia
		return ret
	}
	return *o.ConnectedVia
}

// GetConnectedViaOk returns a tuple with the ConnectedVia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMedia) GetConnectedViaOk() (*ConnectedVia, bool) {
	if o == nil || o.ConnectedVia == nil {
		return nil, false
	}
	return o.ConnectedVia, true
}

// HasConnectedVia returns a boolean if a field has been set.
func (o *VirtualMedia) HasConnectedVia() bool {
	if o != nil && o.ConnectedVia != nil {
		return true
	}

	return false
}

// SetConnectedVia gets a reference to the given ConnectedVia and assigns it to the ConnectedVia field.
func (o *VirtualMedia) SetConnectedVia(v ConnectedVia) {
	o.ConnectedVia = &v
}

// GetMediaTypes returns the MediaTypes field value if set, zero value otherwise.
func (o *VirtualMedia) GetMediaTypes() []string {
	if o == nil || o.MediaTypes == nil {
		var ret []string
		return ret
	}
	return o.MediaTypes
}

// GetMediaTypesOk returns a tuple with the MediaTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMedia) GetMediaTypesOk() ([]string, bool) {
	if o == nil || o.MediaTypes == nil {
		return nil, false
	}
	return o.MediaTypes, true
}

// HasMediaTypes returns a boolean if a field has been set.
func (o *VirtualMedia) HasMediaTypes() bool {
	if o != nil && o.MediaTypes != nil {
		return true
	}

	return false
}

// SetMediaTypes gets a reference to the given []string and assigns it to the MediaTypes field.
func (o *VirtualMedia) SetMediaTypes(v []string) {
	o.MediaTypes = v
}

// GetWriteProtected returns the WriteProtected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMedia) GetWriteProtected() bool {
	if o == nil || o.WriteProtected.Get() == nil {
		var ret bool
		return ret
	}
	return *o.WriteProtected.Get()
}

// GetWriteProtectedOk returns a tuple with the WriteProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMedia) GetWriteProtectedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WriteProtected.Get(), o.WriteProtected.IsSet()
}

// HasWriteProtected returns a boolean if a field has been set.
func (o *VirtualMedia) HasWriteProtected() bool {
	if o != nil && o.WriteProtected.IsSet() {
		return true
	}

	return false
}

// SetWriteProtected gets a reference to the given NullableBool and assigns it to the WriteProtected field.
func (o *VirtualMedia) SetWriteProtected(v bool) {
	o.WriteProtected.Set(&v)
}
// SetWriteProtectedNil sets the value for WriteProtected to be an explicit nil
func (o *VirtualMedia) SetWriteProtectedNil() {
	o.WriteProtected.Set(nil)
}

// UnsetWriteProtected ensures that no value is present for WriteProtected, not even an explicit nil
func (o *VirtualMedia) UnsetWriteProtected() {
	o.WriteProtected.Unset()
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMedia) GetUserName() string {
	if o == nil || o.UserName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMedia) GetUserNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *VirtualMedia) HasUserName() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *VirtualMedia) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *VirtualMedia) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *VirtualMedia) UnsetUserName() {
	o.UserName.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMedia) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMedia) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *VirtualMedia) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *VirtualMedia) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *VirtualMedia) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *VirtualMedia) UnsetPassword() {
	o.Password.Unset()
}

// GetTransferMethod returns the TransferMethod field value if set, zero value otherwise.
func (o *VirtualMedia) GetTransferMethod() TransferMethod {
	if o == nil || o.TransferMethod == nil {
		var ret TransferMethod
		return ret
	}
	return *o.TransferMethod
}

// GetTransferMethodOk returns a tuple with the TransferMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMedia) GetTransferMethodOk() (*TransferMethod, bool) {
	if o == nil || o.TransferMethod == nil {
		return nil, false
	}
	return o.TransferMethod, true
}

// HasTransferMethod returns a boolean if a field has been set.
func (o *VirtualMedia) HasTransferMethod() bool {
	if o != nil && o.TransferMethod != nil {
		return true
	}

	return false
}

// SetTransferMethod gets a reference to the given TransferMethod and assigns it to the TransferMethod field.
func (o *VirtualMedia) SetTransferMethod(v TransferMethod) {
	o.TransferMethod = &v
}

// GetTransferProtocolType returns the TransferProtocolType field value if set, zero value otherwise.
func (o *VirtualMedia) GetTransferProtocolType() TransferProtocolType {
	if o == nil || o.TransferProtocolType == nil {
		var ret TransferProtocolType
		return ret
	}
	return *o.TransferProtocolType
}

// GetTransferProtocolTypeOk returns a tuple with the TransferProtocolType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMedia) GetTransferProtocolTypeOk() (*TransferProtocolType, bool) {
	if o == nil || o.TransferProtocolType == nil {
		return nil, false
	}
	return o.TransferProtocolType, true
}

// HasTransferProtocolType returns a boolean if a field has been set.
func (o *VirtualMedia) HasTransferProtocolType() bool {
	if o != nil && o.TransferProtocolType != nil {
		return true
	}

	return false
}

// SetTransferProtocolType gets a reference to the given TransferProtocolType and assigns it to the TransferProtocolType field.
func (o *VirtualMedia) SetTransferProtocolType(v TransferProtocolType) {
	o.TransferProtocolType = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *VirtualMedia) GetActions() VirtualMediaActions {
	if o == nil || o.Actions == nil {
		var ret VirtualMediaActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMedia) GetActionsOk() (*VirtualMediaActions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *VirtualMedia) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given VirtualMediaActions and assigns it to the Actions field.
func (o *VirtualMedia) SetActions(v VirtualMediaActions) {
	o.Actions = &v
}

func (o VirtualMedia) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if true {
		toSerialize["Name"] = o.Name
	}
	if true {
		toSerialize["@odata.type"] = o.OdataType
	}
	if true {
		toSerialize["@odata.id"] = o.OdataId
	}
	if o.OdataContext != nil {
		toSerialize["@odata.context"] = o.OdataContext
	}
	if o.RedfishCopyright != nil {
		toSerialize["@Redfish.Copyright"] = o.RedfishCopyright
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.Image.IsSet() {
		toSerialize["Image"] = o.Image.Get()
	}
	if o.ImageName.IsSet() {
		toSerialize["ImageName"] = o.ImageName.Get()
	}
	if o.Inserted.IsSet() {
		toSerialize["Inserted"] = o.Inserted.Get()
	}
	if o.ConnectedVia != nil {
		toSerialize["ConnectedVia"] = o.ConnectedVia
	}
	if o.MediaTypes != nil {
		toSerialize["MediaTypes"] = o.MediaTypes
	}
	if o.WriteProtected.IsSet() {
		toSerialize["WriteProtected"] = o.WriteProtected.Get()
	}
	if o.UserName.IsSet() {
		toSerialize["UserName"] = o.UserName.Get()
	}
	if o.Password.IsSet() {
		toSerialize["Password"] = o.Password.Get()
	}
	if o.TransferMethod != nil {
		toSerialize["TransferMethod"] = o.TransferMethod
	}
	if o.TransferProtocolType != nil {
		toSerialize["TransferProtocolType"] = o.TransferProtocolType
	}
	if o.Actions != nil {
		toSerialize["Actions"] = o.Actions
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualMedia struct {
	value *VirtualMedia
	isSet bool
}

func (v NullableVirtualMedia) Get() *VirtualMedia {
	return v.value
}

func (v *NullableVirtualMedia) Set(val *VirtualMedia) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMedia) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMedia(val *VirtualMedia) *NullableVirtualMedia {
	return &NullableVirtualMedia{value: val, isSet: true}
}

func (v NullableVirtualMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



# Manager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The name of the resource. | [optional] [readonly] 
**Name** | **string** | The name of the resource. | [readonly] 
**UUID** | Pointer to **string** |  | [optional] 
**ServiceEntryPointUUID** | Pointer to **string** |  | [optional] 
**OdataType** | **string** | The type of a resource. | [readonly] 
**OdataId** | **string** | The unique identifier for a resource. | [readonly] 
**OdataContext** | Pointer to **string** | The OData description of a payload. | [optional] [readonly] 
**RedfishCopyright** | Pointer to **string** | redfish copyright | [optional] 
**Model** | Pointer to **NullableString** |  | [optional] [readonly] 
**ManagerType** | Pointer to [**ManagerType**](ManagerType.md) |  | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**DateTime** | Pointer to **NullableString** |  | [optional] 
**DateTimeLocalOffset** | Pointer to **NullableString** | The time offset from UTC that the DateTime property is set to in format: +06:00 . | [optional] 
**Description** | Pointer to **NullableString** | description | [optional] [readonly] 
**EthernetInterfaces** | Pointer to [**IdRef**](IdRef.md) |  | [optional] 
**FirmwareVersion** | Pointer to **NullableString** |  | [optional] [readonly] 
**Links** | Pointer to [**ManagerLinks**](ManagerLinks.md) |  | [optional] 
**PowerState** | Pointer to [**PowerState**](PowerState.md) |  | [optional] 
**VirtualMedia** | Pointer to [**IdRef**](IdRef.md) |  | [optional] 

## Methods

### NewManager

`func NewManager(name string, odataType string, odataId string, ) *Manager`

NewManager instantiates a new Manager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagerWithDefaults

`func NewManagerWithDefaults() *Manager`

NewManagerWithDefaults instantiates a new Manager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Manager) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Manager) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Manager) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Manager) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Manager) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Manager) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Manager) SetName(v string)`

SetName sets Name field to given value.


### GetUUID

`func (o *Manager) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *Manager) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *Manager) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *Manager) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetServiceEntryPointUUID

`func (o *Manager) GetServiceEntryPointUUID() string`

GetServiceEntryPointUUID returns the ServiceEntryPointUUID field if non-nil, zero value otherwise.

### GetServiceEntryPointUUIDOk

`func (o *Manager) GetServiceEntryPointUUIDOk() (*string, bool)`

GetServiceEntryPointUUIDOk returns a tuple with the ServiceEntryPointUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEntryPointUUID

`func (o *Manager) SetServiceEntryPointUUID(v string)`

SetServiceEntryPointUUID sets ServiceEntryPointUUID field to given value.

### HasServiceEntryPointUUID

`func (o *Manager) HasServiceEntryPointUUID() bool`

HasServiceEntryPointUUID returns a boolean if a field has been set.

### GetOdataType

`func (o *Manager) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *Manager) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *Manager) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataId

`func (o *Manager) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *Manager) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *Manager) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataContext

`func (o *Manager) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *Manager) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *Manager) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.

### HasOdataContext

`func (o *Manager) HasOdataContext() bool`

HasOdataContext returns a boolean if a field has been set.

### GetRedfishCopyright

`func (o *Manager) GetRedfishCopyright() string`

GetRedfishCopyright returns the RedfishCopyright field if non-nil, zero value otherwise.

### GetRedfishCopyrightOk

`func (o *Manager) GetRedfishCopyrightOk() (*string, bool)`

GetRedfishCopyrightOk returns a tuple with the RedfishCopyright field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedfishCopyright

`func (o *Manager) SetRedfishCopyright(v string)`

SetRedfishCopyright sets RedfishCopyright field to given value.

### HasRedfishCopyright

`func (o *Manager) HasRedfishCopyright() bool`

HasRedfishCopyright returns a boolean if a field has been set.

### GetModel

`func (o *Manager) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Manager) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Manager) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Manager) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *Manager) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *Manager) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetManagerType

`func (o *Manager) GetManagerType() ManagerType`

GetManagerType returns the ManagerType field if non-nil, zero value otherwise.

### GetManagerTypeOk

`func (o *Manager) GetManagerTypeOk() (*ManagerType, bool)`

GetManagerTypeOk returns a tuple with the ManagerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerType

`func (o *Manager) SetManagerType(v ManagerType)`

SetManagerType sets ManagerType field to given value.

### HasManagerType

`func (o *Manager) HasManagerType() bool`

HasManagerType returns a boolean if a field has been set.

### GetStatus

`func (o *Manager) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Manager) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Manager) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Manager) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateTime

`func (o *Manager) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *Manager) GetDateTimeOk() (*string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *Manager) SetDateTime(v string)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *Manager) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### SetDateTimeNil

`func (o *Manager) SetDateTimeNil(b bool)`

 SetDateTimeNil sets the value for DateTime to be an explicit nil

### UnsetDateTime
`func (o *Manager) UnsetDateTime()`

UnsetDateTime ensures that no value is present for DateTime, not even an explicit nil
### GetDateTimeLocalOffset

`func (o *Manager) GetDateTimeLocalOffset() string`

GetDateTimeLocalOffset returns the DateTimeLocalOffset field if non-nil, zero value otherwise.

### GetDateTimeLocalOffsetOk

`func (o *Manager) GetDateTimeLocalOffsetOk() (*string, bool)`

GetDateTimeLocalOffsetOk returns a tuple with the DateTimeLocalOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeLocalOffset

`func (o *Manager) SetDateTimeLocalOffset(v string)`

SetDateTimeLocalOffset sets DateTimeLocalOffset field to given value.

### HasDateTimeLocalOffset

`func (o *Manager) HasDateTimeLocalOffset() bool`

HasDateTimeLocalOffset returns a boolean if a field has been set.

### SetDateTimeLocalOffsetNil

`func (o *Manager) SetDateTimeLocalOffsetNil(b bool)`

 SetDateTimeLocalOffsetNil sets the value for DateTimeLocalOffset to be an explicit nil

### UnsetDateTimeLocalOffset
`func (o *Manager) UnsetDateTimeLocalOffset()`

UnsetDateTimeLocalOffset ensures that no value is present for DateTimeLocalOffset, not even an explicit nil
### GetDescription

`func (o *Manager) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Manager) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Manager) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Manager) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Manager) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Manager) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEthernetInterfaces

`func (o *Manager) GetEthernetInterfaces() IdRef`

GetEthernetInterfaces returns the EthernetInterfaces field if non-nil, zero value otherwise.

### GetEthernetInterfacesOk

`func (o *Manager) GetEthernetInterfacesOk() (*IdRef, bool)`

GetEthernetInterfacesOk returns a tuple with the EthernetInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetInterfaces

`func (o *Manager) SetEthernetInterfaces(v IdRef)`

SetEthernetInterfaces sets EthernetInterfaces field to given value.

### HasEthernetInterfaces

`func (o *Manager) HasEthernetInterfaces() bool`

HasEthernetInterfaces returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *Manager) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *Manager) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *Manager) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *Manager) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### SetFirmwareVersionNil

`func (o *Manager) SetFirmwareVersionNil(b bool)`

 SetFirmwareVersionNil sets the value for FirmwareVersion to be an explicit nil

### UnsetFirmwareVersion
`func (o *Manager) UnsetFirmwareVersion()`

UnsetFirmwareVersion ensures that no value is present for FirmwareVersion, not even an explicit nil
### GetLinks

`func (o *Manager) GetLinks() ManagerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Manager) GetLinksOk() (*ManagerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Manager) SetLinks(v ManagerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Manager) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetPowerState

`func (o *Manager) GetPowerState() PowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *Manager) GetPowerStateOk() (*PowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *Manager) SetPowerState(v PowerState)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *Manager) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetVirtualMedia

`func (o *Manager) GetVirtualMedia() IdRef`

GetVirtualMedia returns the VirtualMedia field if non-nil, zero value otherwise.

### GetVirtualMediaOk

`func (o *Manager) GetVirtualMediaOk() (*IdRef, bool)`

GetVirtualMediaOk returns a tuple with the VirtualMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMedia

`func (o *Manager) SetVirtualMedia(v IdRef)`

SetVirtualMedia sets VirtualMedia field to given value.

### HasVirtualMedia

`func (o *Manager) HasVirtualMedia() bool`

HasVirtualMedia returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



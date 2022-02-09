# ComputerSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The name of the resource. | [optional] [readonly] 
**Name** | **string** | The name of the resource. | [readonly] 
**RedfishVersion** | Pointer to **string** | redfish version | [optional] 
**UUID** | Pointer to **string** |  | [optional] 
**OdataType** | **string** | The type of a resource. | [readonly] 
**OdataId** | **string** | The unique identifier for a resource. | [readonly] 
**OdataContext** | Pointer to **string** | The OData description of a payload. | [optional] [readonly] 
**RedfishCopyright** | Pointer to **string** | redfish copyright | [optional] 
**Bios** | Pointer to [**IdRef**](IdRef.md) |  | [optional] 
**Processors** | Pointer to [**IdRef**](IdRef.md) |  | [optional] 
**Memory** | Pointer to [**IdRef**](IdRef.md) |  | [optional] 
**EthernetInterfaces** | Pointer to [**IdRef**](IdRef.md) |  | [optional] 
**SimpleStorage** | Pointer to [**IdRef**](IdRef.md) |  | [optional] 
**PowerState** | Pointer to [**PowerState**](PowerState.md) |  | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Boot** | Pointer to [**Boot**](Boot.md) |  | [optional] 
**ProcessorSummary** | Pointer to [**ProcessorSummary**](ProcessorSummary.md) |  | [optional] 
**MemorySummary** | Pointer to [**MemorySummary**](MemorySummary.md) |  | [optional] 
**IndicatorLED** | Pointer to [**IndicatorLED**](IndicatorLED.md) |  | [optional] 
**Links** | Pointer to [**SystemLinks**](SystemLinks.md) |  | [optional] 
**Actions** | Pointer to [**ComputerSystemActions**](ComputerSystemActions.md) |  | [optional] 

## Methods

### NewComputerSystem

`func NewComputerSystem(name string, odataType string, odataId string, ) *ComputerSystem`

NewComputerSystem instantiates a new ComputerSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerSystemWithDefaults

`func NewComputerSystemWithDefaults() *ComputerSystem`

NewComputerSystemWithDefaults instantiates a new ComputerSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComputerSystem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputerSystem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputerSystem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComputerSystem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ComputerSystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputerSystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputerSystem) SetName(v string)`

SetName sets Name field to given value.


### GetRedfishVersion

`func (o *ComputerSystem) GetRedfishVersion() string`

GetRedfishVersion returns the RedfishVersion field if non-nil, zero value otherwise.

### GetRedfishVersionOk

`func (o *ComputerSystem) GetRedfishVersionOk() (*string, bool)`

GetRedfishVersionOk returns a tuple with the RedfishVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedfishVersion

`func (o *ComputerSystem) SetRedfishVersion(v string)`

SetRedfishVersion sets RedfishVersion field to given value.

### HasRedfishVersion

`func (o *ComputerSystem) HasRedfishVersion() bool`

HasRedfishVersion returns a boolean if a field has been set.

### GetUUID

`func (o *ComputerSystem) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *ComputerSystem) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *ComputerSystem) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *ComputerSystem) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetOdataType

`func (o *ComputerSystem) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *ComputerSystem) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *ComputerSystem) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataId

`func (o *ComputerSystem) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *ComputerSystem) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *ComputerSystem) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataContext

`func (o *ComputerSystem) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *ComputerSystem) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *ComputerSystem) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.

### HasOdataContext

`func (o *ComputerSystem) HasOdataContext() bool`

HasOdataContext returns a boolean if a field has been set.

### GetRedfishCopyright

`func (o *ComputerSystem) GetRedfishCopyright() string`

GetRedfishCopyright returns the RedfishCopyright field if non-nil, zero value otherwise.

### GetRedfishCopyrightOk

`func (o *ComputerSystem) GetRedfishCopyrightOk() (*string, bool)`

GetRedfishCopyrightOk returns a tuple with the RedfishCopyright field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedfishCopyright

`func (o *ComputerSystem) SetRedfishCopyright(v string)`

SetRedfishCopyright sets RedfishCopyright field to given value.

### HasRedfishCopyright

`func (o *ComputerSystem) HasRedfishCopyright() bool`

HasRedfishCopyright returns a boolean if a field has been set.

### GetBios

`func (o *ComputerSystem) GetBios() IdRef`

GetBios returns the Bios field if non-nil, zero value otherwise.

### GetBiosOk

`func (o *ComputerSystem) GetBiosOk() (*IdRef, bool)`

GetBiosOk returns a tuple with the Bios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBios

`func (o *ComputerSystem) SetBios(v IdRef)`

SetBios sets Bios field to given value.

### HasBios

`func (o *ComputerSystem) HasBios() bool`

HasBios returns a boolean if a field has been set.

### GetProcessors

`func (o *ComputerSystem) GetProcessors() IdRef`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *ComputerSystem) GetProcessorsOk() (*IdRef, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *ComputerSystem) SetProcessors(v IdRef)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *ComputerSystem) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### GetMemory

`func (o *ComputerSystem) GetMemory() IdRef`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ComputerSystem) GetMemoryOk() (*IdRef, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ComputerSystem) SetMemory(v IdRef)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ComputerSystem) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetEthernetInterfaces

`func (o *ComputerSystem) GetEthernetInterfaces() IdRef`

GetEthernetInterfaces returns the EthernetInterfaces field if non-nil, zero value otherwise.

### GetEthernetInterfacesOk

`func (o *ComputerSystem) GetEthernetInterfacesOk() (*IdRef, bool)`

GetEthernetInterfacesOk returns a tuple with the EthernetInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetInterfaces

`func (o *ComputerSystem) SetEthernetInterfaces(v IdRef)`

SetEthernetInterfaces sets EthernetInterfaces field to given value.

### HasEthernetInterfaces

`func (o *ComputerSystem) HasEthernetInterfaces() bool`

HasEthernetInterfaces returns a boolean if a field has been set.

### GetSimpleStorage

`func (o *ComputerSystem) GetSimpleStorage() IdRef`

GetSimpleStorage returns the SimpleStorage field if non-nil, zero value otherwise.

### GetSimpleStorageOk

`func (o *ComputerSystem) GetSimpleStorageOk() (*IdRef, bool)`

GetSimpleStorageOk returns a tuple with the SimpleStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleStorage

`func (o *ComputerSystem) SetSimpleStorage(v IdRef)`

SetSimpleStorage sets SimpleStorage field to given value.

### HasSimpleStorage

`func (o *ComputerSystem) HasSimpleStorage() bool`

HasSimpleStorage returns a boolean if a field has been set.

### GetPowerState

`func (o *ComputerSystem) GetPowerState() PowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *ComputerSystem) GetPowerStateOk() (*PowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *ComputerSystem) SetPowerState(v PowerState)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *ComputerSystem) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetStatus

`func (o *ComputerSystem) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComputerSystem) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComputerSystem) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ComputerSystem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBoot

`func (o *ComputerSystem) GetBoot() Boot`

GetBoot returns the Boot field if non-nil, zero value otherwise.

### GetBootOk

`func (o *ComputerSystem) GetBootOk() (*Boot, bool)`

GetBootOk returns a tuple with the Boot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoot

`func (o *ComputerSystem) SetBoot(v Boot)`

SetBoot sets Boot field to given value.

### HasBoot

`func (o *ComputerSystem) HasBoot() bool`

HasBoot returns a boolean if a field has been set.

### GetProcessorSummary

`func (o *ComputerSystem) GetProcessorSummary() ProcessorSummary`

GetProcessorSummary returns the ProcessorSummary field if non-nil, zero value otherwise.

### GetProcessorSummaryOk

`func (o *ComputerSystem) GetProcessorSummaryOk() (*ProcessorSummary, bool)`

GetProcessorSummaryOk returns a tuple with the ProcessorSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorSummary

`func (o *ComputerSystem) SetProcessorSummary(v ProcessorSummary)`

SetProcessorSummary sets ProcessorSummary field to given value.

### HasProcessorSummary

`func (o *ComputerSystem) HasProcessorSummary() bool`

HasProcessorSummary returns a boolean if a field has been set.

### GetMemorySummary

`func (o *ComputerSystem) GetMemorySummary() MemorySummary`

GetMemorySummary returns the MemorySummary field if non-nil, zero value otherwise.

### GetMemorySummaryOk

`func (o *ComputerSystem) GetMemorySummaryOk() (*MemorySummary, bool)`

GetMemorySummaryOk returns a tuple with the MemorySummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySummary

`func (o *ComputerSystem) SetMemorySummary(v MemorySummary)`

SetMemorySummary sets MemorySummary field to given value.

### HasMemorySummary

`func (o *ComputerSystem) HasMemorySummary() bool`

HasMemorySummary returns a boolean if a field has been set.

### GetIndicatorLED

`func (o *ComputerSystem) GetIndicatorLED() IndicatorLED`

GetIndicatorLED returns the IndicatorLED field if non-nil, zero value otherwise.

### GetIndicatorLEDOk

`func (o *ComputerSystem) GetIndicatorLEDOk() (*IndicatorLED, bool)`

GetIndicatorLEDOk returns a tuple with the IndicatorLED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatorLED

`func (o *ComputerSystem) SetIndicatorLED(v IndicatorLED)`

SetIndicatorLED sets IndicatorLED field to given value.

### HasIndicatorLED

`func (o *ComputerSystem) HasIndicatorLED() bool`

HasIndicatorLED returns a boolean if a field has been set.

### GetLinks

`func (o *ComputerSystem) GetLinks() SystemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ComputerSystem) GetLinksOk() (*SystemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ComputerSystem) SetLinks(v SystemLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ComputerSystem) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetActions

`func (o *ComputerSystem) GetActions() ComputerSystemActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ComputerSystem) GetActionsOk() (*ComputerSystemActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ComputerSystem) SetActions(v ComputerSystemActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ComputerSystem) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



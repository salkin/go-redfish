# UpdateService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirmwareInventory** | Pointer to [**FirmwareInventory**](FirmwareInventory.md) |  | [optional] 
**Id** | Pointer to **string** | The name of the resource. | [optional] [readonly] 
**Name** | **string** | The name of the resource. | [readonly] 
**OdataType** | **string** | The type of a resource. | [readonly] 
**OdataId** | **string** | The unique identifier for a resource. | [readonly] 
**OdataContext** | Pointer to **string** | The OData description of a payload. | [optional] [readonly] 
**Description** | Pointer to **NullableString** | description | [optional] [readonly] 
**HttpPushUri** | Pointer to **string** |  | [optional] [readonly] 
**ServiceEnabled** | Pointer to **NullableBool** |  | [optional] 
**Actions** | Pointer to [**UpdateServiceActions**](UpdateServiceActions.md) |  | [optional] 

## Methods

### NewUpdateService

`func NewUpdateService(name string, odataType string, odataId string, ) *UpdateService`

NewUpdateService instantiates a new UpdateService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServiceWithDefaults

`func NewUpdateServiceWithDefaults() *UpdateService`

NewUpdateServiceWithDefaults instantiates a new UpdateService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirmwareInventory

`func (o *UpdateService) GetFirmwareInventory() FirmwareInventory`

GetFirmwareInventory returns the FirmwareInventory field if non-nil, zero value otherwise.

### GetFirmwareInventoryOk

`func (o *UpdateService) GetFirmwareInventoryOk() (*FirmwareInventory, bool)`

GetFirmwareInventoryOk returns a tuple with the FirmwareInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareInventory

`func (o *UpdateService) SetFirmwareInventory(v FirmwareInventory)`

SetFirmwareInventory sets FirmwareInventory field to given value.

### HasFirmwareInventory

`func (o *UpdateService) HasFirmwareInventory() bool`

HasFirmwareInventory returns a boolean if a field has been set.

### GetId

`func (o *UpdateService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateService) SetName(v string)`

SetName sets Name field to given value.


### GetOdataType

`func (o *UpdateService) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *UpdateService) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *UpdateService) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataId

`func (o *UpdateService) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *UpdateService) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *UpdateService) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataContext

`func (o *UpdateService) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *UpdateService) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *UpdateService) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.

### HasOdataContext

`func (o *UpdateService) HasOdataContext() bool`

HasOdataContext returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateService) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateService) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateService) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateService) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateService) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateService) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHttpPushUri

`func (o *UpdateService) GetHttpPushUri() string`

GetHttpPushUri returns the HttpPushUri field if non-nil, zero value otherwise.

### GetHttpPushUriOk

`func (o *UpdateService) GetHttpPushUriOk() (*string, bool)`

GetHttpPushUriOk returns a tuple with the HttpPushUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPushUri

`func (o *UpdateService) SetHttpPushUri(v string)`

SetHttpPushUri sets HttpPushUri field to given value.

### HasHttpPushUri

`func (o *UpdateService) HasHttpPushUri() bool`

HasHttpPushUri returns a boolean if a field has been set.

### GetServiceEnabled

`func (o *UpdateService) GetServiceEnabled() bool`

GetServiceEnabled returns the ServiceEnabled field if non-nil, zero value otherwise.

### GetServiceEnabledOk

`func (o *UpdateService) GetServiceEnabledOk() (*bool, bool)`

GetServiceEnabledOk returns a tuple with the ServiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnabled

`func (o *UpdateService) SetServiceEnabled(v bool)`

SetServiceEnabled sets ServiceEnabled field to given value.

### HasServiceEnabled

`func (o *UpdateService) HasServiceEnabled() bool`

HasServiceEnabled returns a boolean if a field has been set.

### SetServiceEnabledNil

`func (o *UpdateService) SetServiceEnabledNil(b bool)`

 SetServiceEnabledNil sets the value for ServiceEnabled to be an explicit nil

### UnsetServiceEnabled
`func (o *UpdateService) UnsetServiceEnabled()`

UnsetServiceEnabled ensures that no value is present for ServiceEnabled, not even an explicit nil
### GetActions

`func (o *UpdateService) GetActions() UpdateServiceActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *UpdateService) GetActionsOk() (*UpdateServiceActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *UpdateService) SetActions(v UpdateServiceActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *UpdateService) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



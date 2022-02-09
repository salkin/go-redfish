# Root

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The name of the resource. | [optional] [readonly] 
**Name** | **string** | The name of the resource. | [readonly] 
**RedfishVersion** | Pointer to **string** | redfish version | [optional] 
**UUID** | Pointer to **string** |  | [optional] 
**OdataType** | **string** | The type of a resource. | [readonly] 
**OdataId** | **string** | The unique identifier for a resource. | [readonly] 
**RedfishCopyright** | Pointer to **string** | redfish copyright | [optional] 
**Systems** | Pointer to [**IdRef**](IdRef.md) |  | [optional] 
**Managers** | Pointer to [**IdRef**](IdRef.md) |  | [optional] 

## Methods

### NewRoot

`func NewRoot(name string, odataType string, odataId string, ) *Root`

NewRoot instantiates a new Root object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootWithDefaults

`func NewRootWithDefaults() *Root`

NewRootWithDefaults instantiates a new Root object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Root) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Root) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Root) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Root) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Root) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Root) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Root) SetName(v string)`

SetName sets Name field to given value.


### GetRedfishVersion

`func (o *Root) GetRedfishVersion() string`

GetRedfishVersion returns the RedfishVersion field if non-nil, zero value otherwise.

### GetRedfishVersionOk

`func (o *Root) GetRedfishVersionOk() (*string, bool)`

GetRedfishVersionOk returns a tuple with the RedfishVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedfishVersion

`func (o *Root) SetRedfishVersion(v string)`

SetRedfishVersion sets RedfishVersion field to given value.

### HasRedfishVersion

`func (o *Root) HasRedfishVersion() bool`

HasRedfishVersion returns a boolean if a field has been set.

### GetUUID

`func (o *Root) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *Root) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *Root) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *Root) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetOdataType

`func (o *Root) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *Root) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *Root) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataId

`func (o *Root) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *Root) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *Root) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetRedfishCopyright

`func (o *Root) GetRedfishCopyright() string`

GetRedfishCopyright returns the RedfishCopyright field if non-nil, zero value otherwise.

### GetRedfishCopyrightOk

`func (o *Root) GetRedfishCopyrightOk() (*string, bool)`

GetRedfishCopyrightOk returns a tuple with the RedfishCopyright field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedfishCopyright

`func (o *Root) SetRedfishCopyright(v string)`

SetRedfishCopyright sets RedfishCopyright field to given value.

### HasRedfishCopyright

`func (o *Root) HasRedfishCopyright() bool`

HasRedfishCopyright returns a boolean if a field has been set.

### GetSystems

`func (o *Root) GetSystems() IdRef`

GetSystems returns the Systems field if non-nil, zero value otherwise.

### GetSystemsOk

`func (o *Root) GetSystemsOk() (*IdRef, bool)`

GetSystemsOk returns a tuple with the Systems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystems

`func (o *Root) SetSystems(v IdRef)`

SetSystems sets Systems field to given value.

### HasSystems

`func (o *Root) HasSystems() bool`

HasSystems returns a boolean if a field has been set.

### GetManagers

`func (o *Root) GetManagers() IdRef`

GetManagers returns the Managers field if non-nil, zero value otherwise.

### GetManagersOk

`func (o *Root) GetManagersOk() (*IdRef, bool)`

GetManagersOk returns a tuple with the Managers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagers

`func (o *Root) SetManagers(v IdRef)`

SetManagers sets Managers field to given value.

### HasManagers

`func (o *Root) HasManagers() bool`

HasManagers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



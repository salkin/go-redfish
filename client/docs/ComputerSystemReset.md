# ComputerSystemReset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | Pointer to **string** | The unique identifier for a resource. | [optional] [readonly] 
**ResetTypeRedfishAllowableValues** | Pointer to [**[]ResetType**](ResetType.md) |  | [optional] 

## Methods

### NewComputerSystemReset

`func NewComputerSystemReset() *ComputerSystemReset`

NewComputerSystemReset instantiates a new ComputerSystemReset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerSystemResetWithDefaults

`func NewComputerSystemResetWithDefaults() *ComputerSystemReset`

NewComputerSystemResetWithDefaults instantiates a new ComputerSystemReset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *ComputerSystemReset) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ComputerSystemReset) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ComputerSystemReset) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ComputerSystemReset) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetResetTypeRedfishAllowableValues

`func (o *ComputerSystemReset) GetResetTypeRedfishAllowableValues() []ResetType`

GetResetTypeRedfishAllowableValues returns the ResetTypeRedfishAllowableValues field if non-nil, zero value otherwise.

### GetResetTypeRedfishAllowableValuesOk

`func (o *ComputerSystemReset) GetResetTypeRedfishAllowableValuesOk() (*[]ResetType, bool)`

GetResetTypeRedfishAllowableValuesOk returns a tuple with the ResetTypeRedfishAllowableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetTypeRedfishAllowableValues

`func (o *ComputerSystemReset) SetResetTypeRedfishAllowableValues(v []ResetType)`

SetResetTypeRedfishAllowableValues sets ResetTypeRedfishAllowableValues field to given value.

### HasResetTypeRedfishAllowableValues

`func (o *ComputerSystemReset) HasResetTypeRedfishAllowableValues() bool`

HasResetTypeRedfishAllowableValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



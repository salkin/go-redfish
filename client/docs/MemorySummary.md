# MemorySummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalSystemMemoryGiB** | Pointer to **NullableFloat32** |  | [optional] [readonly] 
**TotalSystemPersistentMemoryGiB** | Pointer to **NullableFloat32** |  | [optional] [readonly] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 

## Methods

### NewMemorySummary

`func NewMemorySummary() *MemorySummary`

NewMemorySummary instantiates a new MemorySummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemorySummaryWithDefaults

`func NewMemorySummaryWithDefaults() *MemorySummary`

NewMemorySummaryWithDefaults instantiates a new MemorySummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalSystemMemoryGiB

`func (o *MemorySummary) GetTotalSystemMemoryGiB() float32`

GetTotalSystemMemoryGiB returns the TotalSystemMemoryGiB field if non-nil, zero value otherwise.

### GetTotalSystemMemoryGiBOk

`func (o *MemorySummary) GetTotalSystemMemoryGiBOk() (*float32, bool)`

GetTotalSystemMemoryGiBOk returns a tuple with the TotalSystemMemoryGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSystemMemoryGiB

`func (o *MemorySummary) SetTotalSystemMemoryGiB(v float32)`

SetTotalSystemMemoryGiB sets TotalSystemMemoryGiB field to given value.

### HasTotalSystemMemoryGiB

`func (o *MemorySummary) HasTotalSystemMemoryGiB() bool`

HasTotalSystemMemoryGiB returns a boolean if a field has been set.

### SetTotalSystemMemoryGiBNil

`func (o *MemorySummary) SetTotalSystemMemoryGiBNil(b bool)`

 SetTotalSystemMemoryGiBNil sets the value for TotalSystemMemoryGiB to be an explicit nil

### UnsetTotalSystemMemoryGiB
`func (o *MemorySummary) UnsetTotalSystemMemoryGiB()`

UnsetTotalSystemMemoryGiB ensures that no value is present for TotalSystemMemoryGiB, not even an explicit nil
### GetTotalSystemPersistentMemoryGiB

`func (o *MemorySummary) GetTotalSystemPersistentMemoryGiB() float32`

GetTotalSystemPersistentMemoryGiB returns the TotalSystemPersistentMemoryGiB field if non-nil, zero value otherwise.

### GetTotalSystemPersistentMemoryGiBOk

`func (o *MemorySummary) GetTotalSystemPersistentMemoryGiBOk() (*float32, bool)`

GetTotalSystemPersistentMemoryGiBOk returns a tuple with the TotalSystemPersistentMemoryGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSystemPersistentMemoryGiB

`func (o *MemorySummary) SetTotalSystemPersistentMemoryGiB(v float32)`

SetTotalSystemPersistentMemoryGiB sets TotalSystemPersistentMemoryGiB field to given value.

### HasTotalSystemPersistentMemoryGiB

`func (o *MemorySummary) HasTotalSystemPersistentMemoryGiB() bool`

HasTotalSystemPersistentMemoryGiB returns a boolean if a field has been set.

### SetTotalSystemPersistentMemoryGiBNil

`func (o *MemorySummary) SetTotalSystemPersistentMemoryGiBNil(b bool)`

 SetTotalSystemPersistentMemoryGiBNil sets the value for TotalSystemPersistentMemoryGiB to be an explicit nil

### UnsetTotalSystemPersistentMemoryGiB
`func (o *MemorySummary) UnsetTotalSystemPersistentMemoryGiB()`

UnsetTotalSystemPersistentMemoryGiB ensures that no value is present for TotalSystemPersistentMemoryGiB, not even an explicit nil
### GetStatus

`func (o *MemorySummary) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MemorySummary) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MemorySummary) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MemorySummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



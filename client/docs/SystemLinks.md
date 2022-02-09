# SystemLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chassis** | Pointer to [**[]IdRef**](IdRef.md) |  | [optional] 
**ManagedBy** | Pointer to [**[]IdRef**](IdRef.md) |  | [optional] 

## Methods

### NewSystemLinks

`func NewSystemLinks() *SystemLinks`

NewSystemLinks instantiates a new SystemLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemLinksWithDefaults

`func NewSystemLinksWithDefaults() *SystemLinks`

NewSystemLinksWithDefaults instantiates a new SystemLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassis

`func (o *SystemLinks) GetChassis() []IdRef`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *SystemLinks) GetChassisOk() (*[]IdRef, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *SystemLinks) SetChassis(v []IdRef)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *SystemLinks) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### GetManagedBy

`func (o *SystemLinks) GetManagedBy() []IdRef`

GetManagedBy returns the ManagedBy field if non-nil, zero value otherwise.

### GetManagedByOk

`func (o *SystemLinks) GetManagedByOk() (*[]IdRef, bool)`

GetManagedByOk returns a tuple with the ManagedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBy

`func (o *SystemLinks) SetManagedBy(v []IdRef)`

SetManagedBy sets ManagedBy field to given value.

### HasManagedBy

`func (o *SystemLinks) HasManagedBy() bool`

HasManagedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



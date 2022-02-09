# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | Pointer to [**Health**](Health.md) |  | [optional] 
**HealthRollup** | Pointer to [**Health**](Health.md) |  | [optional] 
**State** | Pointer to [**State**](State.md) |  | [optional] 

## Methods

### NewStatus

`func NewStatus() *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *Status) GetHealth() Health`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *Status) GetHealthOk() (*Health, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *Status) SetHealth(v Health)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *Status) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHealthRollup

`func (o *Status) GetHealthRollup() Health`

GetHealthRollup returns the HealthRollup field if non-nil, zero value otherwise.

### GetHealthRollupOk

`func (o *Status) GetHealthRollupOk() (*Health, bool)`

GetHealthRollupOk returns a tuple with the HealthRollup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthRollup

`func (o *Status) SetHealthRollup(v Health)`

SetHealthRollup sets HealthRollup field to given value.

### HasHealthRollup

`func (o *Status) HasHealthRollup() bool`

HasHealthRollup returns a boolean if a field has been set.

### GetState

`func (o *Status) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Status) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Status) SetState(v State)`

SetState sets State field to given value.

### HasState

`func (o *Status) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



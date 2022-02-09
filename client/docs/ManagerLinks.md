# ManagerLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagerForServers** | Pointer to [**[]IdRef**](IdRef.md) |  | [optional] 
**ManagerForChassis** | Pointer to [**[]IdRef**](IdRef.md) |  | [optional] 
**ManagerForSwitches** | Pointer to [**[]IdRef**](IdRef.md) |  | [optional] 
**ManagerInChassis** | Pointer to [**[]IdRef**](IdRef.md) |  | [optional] 

## Methods

### NewManagerLinks

`func NewManagerLinks() *ManagerLinks`

NewManagerLinks instantiates a new ManagerLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagerLinksWithDefaults

`func NewManagerLinksWithDefaults() *ManagerLinks`

NewManagerLinksWithDefaults instantiates a new ManagerLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagerForServers

`func (o *ManagerLinks) GetManagerForServers() []IdRef`

GetManagerForServers returns the ManagerForServers field if non-nil, zero value otherwise.

### GetManagerForServersOk

`func (o *ManagerLinks) GetManagerForServersOk() (*[]IdRef, bool)`

GetManagerForServersOk returns a tuple with the ManagerForServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerForServers

`func (o *ManagerLinks) SetManagerForServers(v []IdRef)`

SetManagerForServers sets ManagerForServers field to given value.

### HasManagerForServers

`func (o *ManagerLinks) HasManagerForServers() bool`

HasManagerForServers returns a boolean if a field has been set.

### GetManagerForChassis

`func (o *ManagerLinks) GetManagerForChassis() []IdRef`

GetManagerForChassis returns the ManagerForChassis field if non-nil, zero value otherwise.

### GetManagerForChassisOk

`func (o *ManagerLinks) GetManagerForChassisOk() (*[]IdRef, bool)`

GetManagerForChassisOk returns a tuple with the ManagerForChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerForChassis

`func (o *ManagerLinks) SetManagerForChassis(v []IdRef)`

SetManagerForChassis sets ManagerForChassis field to given value.

### HasManagerForChassis

`func (o *ManagerLinks) HasManagerForChassis() bool`

HasManagerForChassis returns a boolean if a field has been set.

### GetManagerForSwitches

`func (o *ManagerLinks) GetManagerForSwitches() []IdRef`

GetManagerForSwitches returns the ManagerForSwitches field if non-nil, zero value otherwise.

### GetManagerForSwitchesOk

`func (o *ManagerLinks) GetManagerForSwitchesOk() (*[]IdRef, bool)`

GetManagerForSwitchesOk returns a tuple with the ManagerForSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerForSwitches

`func (o *ManagerLinks) SetManagerForSwitches(v []IdRef)`

SetManagerForSwitches sets ManagerForSwitches field to given value.

### HasManagerForSwitches

`func (o *ManagerLinks) HasManagerForSwitches() bool`

HasManagerForSwitches returns a boolean if a field has been set.

### GetManagerInChassis

`func (o *ManagerLinks) GetManagerInChassis() []IdRef`

GetManagerInChassis returns the ManagerInChassis field if non-nil, zero value otherwise.

### GetManagerInChassisOk

`func (o *ManagerLinks) GetManagerInChassisOk() (*[]IdRef, bool)`

GetManagerInChassisOk returns a tuple with the ManagerInChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerInChassis

`func (o *ManagerLinks) SetManagerInChassis(v []IdRef)`

SetManagerInChassis sets ManagerInChassis field to given value.

### HasManagerInChassis

`func (o *ManagerLinks) HasManagerInChassis() bool`

HasManagerInChassis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



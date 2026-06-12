# ActivityFeedDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**ActivitiesCount** | Pointer to **int32** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewActivityFeedDto

`func NewActivityFeedDto() *ActivityFeedDto`

NewActivityFeedDto instantiates a new ActivityFeedDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityFeedDtoWithDefaults

`func NewActivityFeedDtoWithDefaults() *ActivityFeedDto`

NewActivityFeedDtoWithDefaults instantiates a new ActivityFeedDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityFeedDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityFeedDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityFeedDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityFeedDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ActivityFeedDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ActivityFeedDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ActivityFeedDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ActivityFeedDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ActivityFeedDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ActivityFeedDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ActivityFeedDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ActivityFeedDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetActivitiesCount

`func (o *ActivityFeedDto) GetActivitiesCount() int32`

GetActivitiesCount returns the ActivitiesCount field if non-nil, zero value otherwise.

### GetActivitiesCountOk

`func (o *ActivityFeedDto) GetActivitiesCountOk() (*int32, bool)`

GetActivitiesCountOk returns a tuple with the ActivitiesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivitiesCount

`func (o *ActivityFeedDto) SetActivitiesCount(v int32)`

SetActivitiesCount sets ActivitiesCount field to given value.

### HasActivitiesCount

`func (o *ActivityFeedDto) HasActivitiesCount() bool`

HasActivitiesCount returns a boolean if a field has been set.

### GetTenantId

`func (o *ActivityFeedDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ActivityFeedDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ActivityFeedDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ActivityFeedDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ActivityFeedDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ActivityFeedDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



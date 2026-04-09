# ServiceQueueDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**ServiceId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewServiceQueueDto

`func NewServiceQueueDto() *ServiceQueueDto`

NewServiceQueueDto instantiates a new ServiceQueueDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceQueueDtoWithDefaults

`func NewServiceQueueDtoWithDefaults() *ServiceQueueDto`

NewServiceQueueDtoWithDefaults instantiates a new ServiceQueueDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceQueueDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceQueueDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceQueueDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceQueueDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ServiceQueueDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ServiceQueueDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ServiceQueueDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ServiceQueueDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ServiceQueueDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ServiceQueueDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ServiceQueueDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ServiceQueueDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetServiceId

`func (o *ServiceQueueDto) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceQueueDto) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceQueueDto) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ServiceQueueDto) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### SetServiceIdNil

`func (o *ServiceQueueDto) SetServiceIdNil(b bool)`

 SetServiceIdNil sets the value for ServiceId to be an explicit nil

### UnsetServiceId
`func (o *ServiceQueueDto) UnsetServiceId()`

UnsetServiceId ensures that no value is present for ServiceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



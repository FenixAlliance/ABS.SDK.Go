# TimezoneDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] [readonly] 
**UtcOffset** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTimezoneDto

`func NewTimezoneDto() *TimezoneDto`

NewTimezoneDto instantiates a new TimezoneDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimezoneDtoWithDefaults

`func NewTimezoneDtoWithDefaults() *TimezoneDto`

NewTimezoneDtoWithDefaults instantiates a new TimezoneDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *TimezoneDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TimezoneDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TimezoneDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TimezoneDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TimezoneDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TimezoneDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetId

`func (o *TimezoneDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimezoneDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimezoneDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TimezoneDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TimezoneDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TimezoneDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *TimezoneDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimezoneDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimezoneDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TimezoneDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TimezoneDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TimezoneDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisplayName

`func (o *TimezoneDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TimezoneDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TimezoneDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TimezoneDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *TimezoneDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *TimezoneDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetUtcOffset

`func (o *TimezoneDto) GetUtcOffset() string`

GetUtcOffset returns the UtcOffset field if non-nil, zero value otherwise.

### GetUtcOffsetOk

`func (o *TimezoneDto) GetUtcOffsetOk() (*string, bool)`

GetUtcOffsetOk returns a tuple with the UtcOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcOffset

`func (o *TimezoneDto) SetUtcOffset(v string)`

SetUtcOffset sets UtcOffset field to given value.

### HasUtcOffset

`func (o *TimezoneDto) HasUtcOffset() bool`

HasUtcOffset returns a boolean if a field has been set.

### SetUtcOffsetNil

`func (o *TimezoneDto) SetUtcOffsetNil(b bool)`

 SetUtcOffsetNil sets the value for UtcOffset to be an explicit nil

### UnsetUtcOffset
`func (o *TimezoneDto) UnsetUtcOffset()`

UnsetUtcOffset ensures that no value is present for UtcOffset, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



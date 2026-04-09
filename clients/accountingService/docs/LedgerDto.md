# LedgerDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DateTime** | Pointer to **time.Time** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**LedgerTypeId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLedgerDto

`func NewLedgerDto() *LedgerDto`

NewLedgerDto instantiates a new LedgerDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerDtoWithDefaults

`func NewLedgerDtoWithDefaults() *LedgerDto`

NewLedgerDtoWithDefaults instantiates a new LedgerDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LedgerDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LedgerDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LedgerDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LedgerDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *LedgerDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LedgerDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *LedgerDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LedgerDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LedgerDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LedgerDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *LedgerDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *LedgerDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *LedgerDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LedgerDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LedgerDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LedgerDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LedgerDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LedgerDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *LedgerDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LedgerDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LedgerDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LedgerDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LedgerDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LedgerDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDateTime

`func (o *LedgerDto) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *LedgerDto) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *LedgerDto) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *LedgerDto) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetTenantId

`func (o *LedgerDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *LedgerDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *LedgerDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *LedgerDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *LedgerDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *LedgerDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *LedgerDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *LedgerDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *LedgerDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *LedgerDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *LedgerDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *LedgerDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetLedgerTypeId

`func (o *LedgerDto) GetLedgerTypeId() string`

GetLedgerTypeId returns the LedgerTypeId field if non-nil, zero value otherwise.

### GetLedgerTypeIdOk

`func (o *LedgerDto) GetLedgerTypeIdOk() (*string, bool)`

GetLedgerTypeIdOk returns a tuple with the LedgerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerTypeId

`func (o *LedgerDto) SetLedgerTypeId(v string)`

SetLedgerTypeId sets LedgerTypeId field to given value.

### HasLedgerTypeId

`func (o *LedgerDto) HasLedgerTypeId() bool`

HasLedgerTypeId returns a boolean if a field has been set.

### SetLedgerTypeIdNil

`func (o *LedgerDto) SetLedgerTypeIdNil(b bool)`

 SetLedgerTypeIdNil sets the value for LedgerTypeId to be an explicit nil

### UnsetLedgerTypeId
`func (o *LedgerDto) UnsetLedgerTypeId()`

UnsetLedgerTypeId ensures that no value is present for LedgerTypeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



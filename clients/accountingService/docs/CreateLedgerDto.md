# CreateLedgerDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DateTime** | Pointer to **time.Time** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**LedgerTypeId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateLedgerDto

`func NewCreateLedgerDto() *CreateLedgerDto`

NewCreateLedgerDto instantiates a new CreateLedgerDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLedgerDtoWithDefaults

`func NewCreateLedgerDtoWithDefaults() *CreateLedgerDto`

NewCreateLedgerDtoWithDefaults instantiates a new CreateLedgerDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateLedgerDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateLedgerDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateLedgerDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateLedgerDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CreateLedgerDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CreateLedgerDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CreateLedgerDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CreateLedgerDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *CreateLedgerDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLedgerDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLedgerDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateLedgerDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateLedgerDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateLedgerDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CreateLedgerDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLedgerDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLedgerDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLedgerDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateLedgerDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateLedgerDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDateTime

`func (o *CreateLedgerDto) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *CreateLedgerDto) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *CreateLedgerDto) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *CreateLedgerDto) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetTenantId

`func (o *CreateLedgerDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateLedgerDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateLedgerDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CreateLedgerDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CreateLedgerDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CreateLedgerDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *CreateLedgerDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *CreateLedgerDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *CreateLedgerDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *CreateLedgerDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *CreateLedgerDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *CreateLedgerDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetLedgerTypeId

`func (o *CreateLedgerDto) GetLedgerTypeId() string`

GetLedgerTypeId returns the LedgerTypeId field if non-nil, zero value otherwise.

### GetLedgerTypeIdOk

`func (o *CreateLedgerDto) GetLedgerTypeIdOk() (*string, bool)`

GetLedgerTypeIdOk returns a tuple with the LedgerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerTypeId

`func (o *CreateLedgerDto) SetLedgerTypeId(v string)`

SetLedgerTypeId sets LedgerTypeId field to given value.

### HasLedgerTypeId

`func (o *CreateLedgerDto) HasLedgerTypeId() bool`

HasLedgerTypeId returns a boolean if a field has been set.

### SetLedgerTypeIdNil

`func (o *CreateLedgerDto) SetLedgerTypeIdNil(b bool)`

 SetLedgerTypeIdNil sets the value for LedgerTypeId to be an explicit nil

### UnsetLedgerTypeId
`func (o *CreateLedgerDto) UnsetLedgerTypeId()`

UnsetLedgerTypeId ensures that no value is present for LedgerTypeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



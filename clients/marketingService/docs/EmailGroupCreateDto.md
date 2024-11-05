# EmailGroupCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrolmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEmailGroupCreateDto

`func NewEmailGroupCreateDto() *EmailGroupCreateDto`

NewEmailGroupCreateDto instantiates a new EmailGroupCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailGroupCreateDtoWithDefaults

`func NewEmailGroupCreateDtoWithDefaults() *EmailGroupCreateDto`

NewEmailGroupCreateDtoWithDefaults instantiates a new EmailGroupCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailGroupCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailGroupCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailGroupCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmailGroupCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *EmailGroupCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EmailGroupCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EmailGroupCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EmailGroupCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *EmailGroupCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailGroupCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailGroupCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmailGroupCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EmailGroupCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EmailGroupCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *EmailGroupCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmailGroupCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmailGroupCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EmailGroupCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EmailGroupCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EmailGroupCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *EmailGroupCreateDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EmailGroupCreateDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EmailGroupCreateDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EmailGroupCreateDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTenantId

`func (o *EmailGroupCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *EmailGroupCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *EmailGroupCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *EmailGroupCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *EmailGroupCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *EmailGroupCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrolmentId

`func (o *EmailGroupCreateDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *EmailGroupCreateDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *EmailGroupCreateDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *EmailGroupCreateDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *EmailGroupCreateDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *EmailGroupCreateDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



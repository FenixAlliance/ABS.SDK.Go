# LocalizationStringDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Base** | Pointer to **NullableString** |  | [optional] 
**Comments** | Pointer to **NullableString** |  | [optional] 
**CountryLanguageId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLocalizationStringDto

`func NewLocalizationStringDto() *LocalizationStringDto`

NewLocalizationStringDto instantiates a new LocalizationStringDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalizationStringDtoWithDefaults

`func NewLocalizationStringDtoWithDefaults() *LocalizationStringDto`

NewLocalizationStringDtoWithDefaults instantiates a new LocalizationStringDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocalizationStringDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocalizationStringDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocalizationStringDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LocalizationStringDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *LocalizationStringDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LocalizationStringDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *LocalizationStringDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LocalizationStringDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LocalizationStringDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LocalizationStringDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *LocalizationStringDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *LocalizationStringDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetBase

`func (o *LocalizationStringDto) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *LocalizationStringDto) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *LocalizationStringDto) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *LocalizationStringDto) HasBase() bool`

HasBase returns a boolean if a field has been set.

### SetBaseNil

`func (o *LocalizationStringDto) SetBaseNil(b bool)`

 SetBaseNil sets the value for Base to be an explicit nil

### UnsetBase
`func (o *LocalizationStringDto) UnsetBase()`

UnsetBase ensures that no value is present for Base, not even an explicit nil
### GetComments

`func (o *LocalizationStringDto) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *LocalizationStringDto) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *LocalizationStringDto) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *LocalizationStringDto) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *LocalizationStringDto) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *LocalizationStringDto) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil
### GetCountryLanguageId

`func (o *LocalizationStringDto) GetCountryLanguageId() string`

GetCountryLanguageId returns the CountryLanguageId field if non-nil, zero value otherwise.

### GetCountryLanguageIdOk

`func (o *LocalizationStringDto) GetCountryLanguageIdOk() (*string, bool)`

GetCountryLanguageIdOk returns a tuple with the CountryLanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryLanguageId

`func (o *LocalizationStringDto) SetCountryLanguageId(v string)`

SetCountryLanguageId sets CountryLanguageId field to given value.

### HasCountryLanguageId

`func (o *LocalizationStringDto) HasCountryLanguageId() bool`

HasCountryLanguageId returns a boolean if a field has been set.

### SetCountryLanguageIdNil

`func (o *LocalizationStringDto) SetCountryLanguageIdNil(b bool)`

 SetCountryLanguageIdNil sets the value for CountryLanguageId to be an explicit nil

### UnsetCountryLanguageId
`func (o *LocalizationStringDto) UnsetCountryLanguageId()`

UnsetCountryLanguageId ensures that no value is present for CountryLanguageId, not even an explicit nil
### GetTenantId

`func (o *LocalizationStringDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *LocalizationStringDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *LocalizationStringDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *LocalizationStringDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *LocalizationStringDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *LocalizationStringDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



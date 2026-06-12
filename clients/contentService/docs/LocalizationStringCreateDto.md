# LocalizationStringCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Base** | **string** |  | 
**Comments** | Pointer to **NullableString** |  | [optional] 
**CountryLanguageId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLocalizationStringCreateDto

`func NewLocalizationStringCreateDto(base string, ) *LocalizationStringCreateDto`

NewLocalizationStringCreateDto instantiates a new LocalizationStringCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalizationStringCreateDtoWithDefaults

`func NewLocalizationStringCreateDtoWithDefaults() *LocalizationStringCreateDto`

NewLocalizationStringCreateDtoWithDefaults instantiates a new LocalizationStringCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocalizationStringCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocalizationStringCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocalizationStringCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LocalizationStringCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *LocalizationStringCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LocalizationStringCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LocalizationStringCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LocalizationStringCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBase

`func (o *LocalizationStringCreateDto) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *LocalizationStringCreateDto) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *LocalizationStringCreateDto) SetBase(v string)`

SetBase sets Base field to given value.


### GetComments

`func (o *LocalizationStringCreateDto) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *LocalizationStringCreateDto) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *LocalizationStringCreateDto) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *LocalizationStringCreateDto) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *LocalizationStringCreateDto) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *LocalizationStringCreateDto) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil
### GetCountryLanguageId

`func (o *LocalizationStringCreateDto) GetCountryLanguageId() string`

GetCountryLanguageId returns the CountryLanguageId field if non-nil, zero value otherwise.

### GetCountryLanguageIdOk

`func (o *LocalizationStringCreateDto) GetCountryLanguageIdOk() (*string, bool)`

GetCountryLanguageIdOk returns a tuple with the CountryLanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryLanguageId

`func (o *LocalizationStringCreateDto) SetCountryLanguageId(v string)`

SetCountryLanguageId sets CountryLanguageId field to given value.

### HasCountryLanguageId

`func (o *LocalizationStringCreateDto) HasCountryLanguageId() bool`

HasCountryLanguageId returns a boolean if a field has been set.

### SetCountryLanguageIdNil

`func (o *LocalizationStringCreateDto) SetCountryLanguageIdNil(b bool)`

 SetCountryLanguageIdNil sets the value for CountryLanguageId to be an explicit nil

### UnsetCountryLanguageId
`func (o *LocalizationStringCreateDto) UnsetCountryLanguageId()`

UnsetCountryLanguageId ensures that no value is present for CountryLanguageId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



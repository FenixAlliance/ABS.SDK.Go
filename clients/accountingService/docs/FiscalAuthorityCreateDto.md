# FiscalAuthorityCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**LogoUrl** | Pointer to **NullableString** |  | [optional] 
**WebUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFiscalAuthorityCreateDto

`func NewFiscalAuthorityCreateDto() *FiscalAuthorityCreateDto`

NewFiscalAuthorityCreateDto instantiates a new FiscalAuthorityCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiscalAuthorityCreateDtoWithDefaults

`func NewFiscalAuthorityCreateDtoWithDefaults() *FiscalAuthorityCreateDto`

NewFiscalAuthorityCreateDtoWithDefaults instantiates a new FiscalAuthorityCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FiscalAuthorityCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FiscalAuthorityCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FiscalAuthorityCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FiscalAuthorityCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *FiscalAuthorityCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FiscalAuthorityCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FiscalAuthorityCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *FiscalAuthorityCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *FiscalAuthorityCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FiscalAuthorityCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FiscalAuthorityCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FiscalAuthorityCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FiscalAuthorityCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FiscalAuthorityCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *FiscalAuthorityCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FiscalAuthorityCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FiscalAuthorityCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FiscalAuthorityCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FiscalAuthorityCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FiscalAuthorityCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCountryId

`func (o *FiscalAuthorityCreateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *FiscalAuthorityCreateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *FiscalAuthorityCreateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *FiscalAuthorityCreateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *FiscalAuthorityCreateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *FiscalAuthorityCreateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetLogoUrl

`func (o *FiscalAuthorityCreateDto) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *FiscalAuthorityCreateDto) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *FiscalAuthorityCreateDto) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *FiscalAuthorityCreateDto) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *FiscalAuthorityCreateDto) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *FiscalAuthorityCreateDto) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetWebUrl

`func (o *FiscalAuthorityCreateDto) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *FiscalAuthorityCreateDto) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *FiscalAuthorityCreateDto) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *FiscalAuthorityCreateDto) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *FiscalAuthorityCreateDto) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *FiscalAuthorityCreateDto) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



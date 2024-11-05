# CountryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Iso3** | Pointer to **NullableString** |  | [optional] 
**Iso2** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**NativeName** | Pointer to **NullableString** |  | [optional] 
**FlagUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCountryDto

`func NewCountryDto() *CountryDto`

NewCountryDto instantiates a new CountryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryDtoWithDefaults

`func NewCountryDtoWithDefaults() *CountryDto`

NewCountryDtoWithDefaults instantiates a new CountryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CountryDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CountryDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CountryDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CountryDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CountryDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CountryDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CountryDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CountryDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CountryDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CountryDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CountryDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CountryDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetIso3

`func (o *CountryDto) GetIso3() string`

GetIso3 returns the Iso3 field if non-nil, zero value otherwise.

### GetIso3Ok

`func (o *CountryDto) GetIso3Ok() (*string, bool)`

GetIso3Ok returns a tuple with the Iso3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso3

`func (o *CountryDto) SetIso3(v string)`

SetIso3 sets Iso3 field to given value.

### HasIso3

`func (o *CountryDto) HasIso3() bool`

HasIso3 returns a boolean if a field has been set.

### SetIso3Nil

`func (o *CountryDto) SetIso3Nil(b bool)`

 SetIso3Nil sets the value for Iso3 to be an explicit nil

### UnsetIso3
`func (o *CountryDto) UnsetIso3()`

UnsetIso3 ensures that no value is present for Iso3, not even an explicit nil
### GetIso2

`func (o *CountryDto) GetIso2() string`

GetIso2 returns the Iso2 field if non-nil, zero value otherwise.

### GetIso2Ok

`func (o *CountryDto) GetIso2Ok() (*string, bool)`

GetIso2Ok returns a tuple with the Iso2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso2

`func (o *CountryDto) SetIso2(v string)`

SetIso2 sets Iso2 field to given value.

### HasIso2

`func (o *CountryDto) HasIso2() bool`

HasIso2 returns a boolean if a field has been set.

### SetIso2Nil

`func (o *CountryDto) SetIso2Nil(b bool)`

 SetIso2Nil sets the value for Iso2 to be an explicit nil

### UnsetIso2
`func (o *CountryDto) UnsetIso2()`

UnsetIso2 ensures that no value is present for Iso2, not even an explicit nil
### GetName

`func (o *CountryDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CountryDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CountryDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CountryDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CountryDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CountryDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNativeName

`func (o *CountryDto) GetNativeName() string`

GetNativeName returns the NativeName field if non-nil, zero value otherwise.

### GetNativeNameOk

`func (o *CountryDto) GetNativeNameOk() (*string, bool)`

GetNativeNameOk returns a tuple with the NativeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeName

`func (o *CountryDto) SetNativeName(v string)`

SetNativeName sets NativeName field to given value.

### HasNativeName

`func (o *CountryDto) HasNativeName() bool`

HasNativeName returns a boolean if a field has been set.

### SetNativeNameNil

`func (o *CountryDto) SetNativeNameNil(b bool)`

 SetNativeNameNil sets the value for NativeName to be an explicit nil

### UnsetNativeName
`func (o *CountryDto) UnsetNativeName()`

UnsetNativeName ensures that no value is present for NativeName, not even an explicit nil
### GetFlagUrl

`func (o *CountryDto) GetFlagUrl() string`

GetFlagUrl returns the FlagUrl field if non-nil, zero value otherwise.

### GetFlagUrlOk

`func (o *CountryDto) GetFlagUrlOk() (*string, bool)`

GetFlagUrlOk returns a tuple with the FlagUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagUrl

`func (o *CountryDto) SetFlagUrl(v string)`

SetFlagUrl sets FlagUrl field to given value.

### HasFlagUrl

`func (o *CountryDto) HasFlagUrl() bool`

HasFlagUrl returns a boolean if a field has been set.

### SetFlagUrlNil

`func (o *CountryDto) SetFlagUrlNil(b bool)`

 SetFlagUrlNil sets the value for FlagUrl to be an explicit nil

### UnsetFlagUrl
`func (o *CountryDto) UnsetFlagUrl()`

UnsetFlagUrl ensures that no value is present for FlagUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



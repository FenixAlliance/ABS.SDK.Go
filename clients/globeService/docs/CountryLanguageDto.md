# CountryLanguageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Iso6391** | Pointer to **NullableString** |  | [optional] 
**Iso6392** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**LanguageNativeName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCountryLanguageDto

`func NewCountryLanguageDto() *CountryLanguageDto`

NewCountryLanguageDto instantiates a new CountryLanguageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryLanguageDtoWithDefaults

`func NewCountryLanguageDtoWithDefaults() *CountryLanguageDto`

NewCountryLanguageDtoWithDefaults instantiates a new CountryLanguageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CountryLanguageDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CountryLanguageDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CountryLanguageDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CountryLanguageDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CountryLanguageDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CountryLanguageDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CountryLanguageDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CountryLanguageDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CountryLanguageDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CountryLanguageDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CountryLanguageDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CountryLanguageDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetIso6391

`func (o *CountryLanguageDto) GetIso6391() string`

GetIso6391 returns the Iso6391 field if non-nil, zero value otherwise.

### GetIso6391Ok

`func (o *CountryLanguageDto) GetIso6391Ok() (*string, bool)`

GetIso6391Ok returns a tuple with the Iso6391 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso6391

`func (o *CountryLanguageDto) SetIso6391(v string)`

SetIso6391 sets Iso6391 field to given value.

### HasIso6391

`func (o *CountryLanguageDto) HasIso6391() bool`

HasIso6391 returns a boolean if a field has been set.

### SetIso6391Nil

`func (o *CountryLanguageDto) SetIso6391Nil(b bool)`

 SetIso6391Nil sets the value for Iso6391 to be an explicit nil

### UnsetIso6391
`func (o *CountryLanguageDto) UnsetIso6391()`

UnsetIso6391 ensures that no value is present for Iso6391, not even an explicit nil
### GetIso6392

`func (o *CountryLanguageDto) GetIso6392() string`

GetIso6392 returns the Iso6392 field if non-nil, zero value otherwise.

### GetIso6392Ok

`func (o *CountryLanguageDto) GetIso6392Ok() (*string, bool)`

GetIso6392Ok returns a tuple with the Iso6392 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso6392

`func (o *CountryLanguageDto) SetIso6392(v string)`

SetIso6392 sets Iso6392 field to given value.

### HasIso6392

`func (o *CountryLanguageDto) HasIso6392() bool`

HasIso6392 returns a boolean if a field has been set.

### SetIso6392Nil

`func (o *CountryLanguageDto) SetIso6392Nil(b bool)`

 SetIso6392Nil sets the value for Iso6392 to be an explicit nil

### UnsetIso6392
`func (o *CountryLanguageDto) UnsetIso6392()`

UnsetIso6392 ensures that no value is present for Iso6392, not even an explicit nil
### GetEnabled

`func (o *CountryLanguageDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CountryLanguageDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CountryLanguageDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CountryLanguageDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *CountryLanguageDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CountryLanguageDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CountryLanguageDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CountryLanguageDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CountryLanguageDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CountryLanguageDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLanguageNativeName

`func (o *CountryLanguageDto) GetLanguageNativeName() string`

GetLanguageNativeName returns the LanguageNativeName field if non-nil, zero value otherwise.

### GetLanguageNativeNameOk

`func (o *CountryLanguageDto) GetLanguageNativeNameOk() (*string, bool)`

GetLanguageNativeNameOk returns a tuple with the LanguageNativeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageNativeName

`func (o *CountryLanguageDto) SetLanguageNativeName(v string)`

SetLanguageNativeName sets LanguageNativeName field to given value.

### HasLanguageNativeName

`func (o *CountryLanguageDto) HasLanguageNativeName() bool`

HasLanguageNativeName returns a boolean if a field has been set.

### SetLanguageNativeNameNil

`func (o *CountryLanguageDto) SetLanguageNativeNameNil(b bool)`

 SetLanguageNativeNameNil sets the value for LanguageNativeName to be an explicit nil

### UnsetLanguageNativeName
`func (o *CountryLanguageDto) UnsetLanguageNativeName()`

UnsetLanguageNativeName ensures that no value is present for LanguageNativeName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



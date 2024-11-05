# UserSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**DateFormat** | Pointer to **NullableString** |  | [optional] 
**CurrencyFormat** | Pointer to **NullableString** |  | [optional] 
**DateTimeFormat** | Pointer to **NullableString** |  | [optional] 
**SiteTheme** | Pointer to **int32** |  | [optional] 

## Methods

### NewUserSettingsDto

`func NewUserSettingsDto() *UserSettingsDto`

NewUserSettingsDto instantiates a new UserSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSettingsDtoWithDefaults

`func NewUserSettingsDtoWithDefaults() *UserSettingsDto`

NewUserSettingsDtoWithDefaults instantiates a new UserSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSettingsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSettingsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSettingsDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserSettingsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UserSettingsDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UserSettingsDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *UserSettingsDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UserSettingsDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UserSettingsDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *UserSettingsDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *UserSettingsDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *UserSettingsDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetPageSize

`func (o *UserSettingsDto) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *UserSettingsDto) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *UserSettingsDto) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *UserSettingsDto) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetDateFormat

`func (o *UserSettingsDto) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *UserSettingsDto) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *UserSettingsDto) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.

### HasDateFormat

`func (o *UserSettingsDto) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### SetDateFormatNil

`func (o *UserSettingsDto) SetDateFormatNil(b bool)`

 SetDateFormatNil sets the value for DateFormat to be an explicit nil

### UnsetDateFormat
`func (o *UserSettingsDto) UnsetDateFormat()`

UnsetDateFormat ensures that no value is present for DateFormat, not even an explicit nil
### GetCurrencyFormat

`func (o *UserSettingsDto) GetCurrencyFormat() string`

GetCurrencyFormat returns the CurrencyFormat field if non-nil, zero value otherwise.

### GetCurrencyFormatOk

`func (o *UserSettingsDto) GetCurrencyFormatOk() (*string, bool)`

GetCurrencyFormatOk returns a tuple with the CurrencyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyFormat

`func (o *UserSettingsDto) SetCurrencyFormat(v string)`

SetCurrencyFormat sets CurrencyFormat field to given value.

### HasCurrencyFormat

`func (o *UserSettingsDto) HasCurrencyFormat() bool`

HasCurrencyFormat returns a boolean if a field has been set.

### SetCurrencyFormatNil

`func (o *UserSettingsDto) SetCurrencyFormatNil(b bool)`

 SetCurrencyFormatNil sets the value for CurrencyFormat to be an explicit nil

### UnsetCurrencyFormat
`func (o *UserSettingsDto) UnsetCurrencyFormat()`

UnsetCurrencyFormat ensures that no value is present for CurrencyFormat, not even an explicit nil
### GetDateTimeFormat

`func (o *UserSettingsDto) GetDateTimeFormat() string`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *UserSettingsDto) GetDateTimeFormatOk() (*string, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *UserSettingsDto) SetDateTimeFormat(v string)`

SetDateTimeFormat sets DateTimeFormat field to given value.

### HasDateTimeFormat

`func (o *UserSettingsDto) HasDateTimeFormat() bool`

HasDateTimeFormat returns a boolean if a field has been set.

### SetDateTimeFormatNil

`func (o *UserSettingsDto) SetDateTimeFormatNil(b bool)`

 SetDateTimeFormatNil sets the value for DateTimeFormat to be an explicit nil

### UnsetDateTimeFormat
`func (o *UserSettingsDto) UnsetDateTimeFormat()`

UnsetDateTimeFormat ensures that no value is present for DateTimeFormat, not even an explicit nil
### GetSiteTheme

`func (o *UserSettingsDto) GetSiteTheme() int32`

GetSiteTheme returns the SiteTheme field if non-nil, zero value otherwise.

### GetSiteThemeOk

`func (o *UserSettingsDto) GetSiteThemeOk() (*int32, bool)`

GetSiteThemeOk returns a tuple with the SiteTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTheme

`func (o *UserSettingsDto) SetSiteTheme(v int32)`

SetSiteTheme sets SiteTheme field to given value.

### HasSiteTheme

`func (o *UserSettingsDto) HasSiteTheme() bool`

HasSiteTheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UserSettingsUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageSize** | Pointer to **NullableInt32** |  | [optional] 
**DateFormat** | **string** |  | 
**CurrencyFormat** | **string** |  | 
**DateTimeFormat** | **string** |  | 
**SiteTheme** | **int32** |  | 

## Methods

### NewUserSettingsUpdateDto

`func NewUserSettingsUpdateDto(dateFormat string, currencyFormat string, dateTimeFormat string, siteTheme int32, ) *UserSettingsUpdateDto`

NewUserSettingsUpdateDto instantiates a new UserSettingsUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSettingsUpdateDtoWithDefaults

`func NewUserSettingsUpdateDtoWithDefaults() *UserSettingsUpdateDto`

NewUserSettingsUpdateDtoWithDefaults instantiates a new UserSettingsUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageSize

`func (o *UserSettingsUpdateDto) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *UserSettingsUpdateDto) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *UserSettingsUpdateDto) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *UserSettingsUpdateDto) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSizeNil

`func (o *UserSettingsUpdateDto) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *UserSettingsUpdateDto) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
### GetDateFormat

`func (o *UserSettingsUpdateDto) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *UserSettingsUpdateDto) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *UserSettingsUpdateDto) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.


### GetCurrencyFormat

`func (o *UserSettingsUpdateDto) GetCurrencyFormat() string`

GetCurrencyFormat returns the CurrencyFormat field if non-nil, zero value otherwise.

### GetCurrencyFormatOk

`func (o *UserSettingsUpdateDto) GetCurrencyFormatOk() (*string, bool)`

GetCurrencyFormatOk returns a tuple with the CurrencyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyFormat

`func (o *UserSettingsUpdateDto) SetCurrencyFormat(v string)`

SetCurrencyFormat sets CurrencyFormat field to given value.


### GetDateTimeFormat

`func (o *UserSettingsUpdateDto) GetDateTimeFormat() string`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *UserSettingsUpdateDto) GetDateTimeFormatOk() (*string, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *UserSettingsUpdateDto) SetDateTimeFormat(v string)`

SetDateTimeFormat sets DateTimeFormat field to given value.


### GetSiteTheme

`func (o *UserSettingsUpdateDto) GetSiteTheme() int32`

GetSiteTheme returns the SiteTheme field if non-nil, zero value otherwise.

### GetSiteThemeOk

`func (o *UserSettingsUpdateDto) GetSiteThemeOk() (*int32, bool)`

GetSiteThemeOk returns a tuple with the SiteTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTheme

`func (o *UserSettingsUpdateDto) SetSiteTheme(v int32)`

SetSiteTheme sets SiteTheme field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



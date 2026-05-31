# ShippingZoneCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Default** | Pointer to **bool** |  | [optional] 
**Everywhere** | Pointer to **bool** |  | [optional] 
**PostalCodes** | Pointer to **NullableString** |  | [optional] 
**CountryCodes** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewShippingZoneCreateDto

`func NewShippingZoneCreateDto(name string, ) *ShippingZoneCreateDto`

NewShippingZoneCreateDto instantiates a new ShippingZoneCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingZoneCreateDtoWithDefaults

`func NewShippingZoneCreateDtoWithDefaults() *ShippingZoneCreateDto`

NewShippingZoneCreateDtoWithDefaults instantiates a new ShippingZoneCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShippingZoneCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingZoneCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingZoneCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShippingZoneCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ShippingZoneCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ShippingZoneCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ShippingZoneCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ShippingZoneCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *ShippingZoneCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShippingZoneCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShippingZoneCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetDefault

`func (o *ShippingZoneCreateDto) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ShippingZoneCreateDto) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ShippingZoneCreateDto) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ShippingZoneCreateDto) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEverywhere

`func (o *ShippingZoneCreateDto) GetEverywhere() bool`

GetEverywhere returns the Everywhere field if non-nil, zero value otherwise.

### GetEverywhereOk

`func (o *ShippingZoneCreateDto) GetEverywhereOk() (*bool, bool)`

GetEverywhereOk returns a tuple with the Everywhere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEverywhere

`func (o *ShippingZoneCreateDto) SetEverywhere(v bool)`

SetEverywhere sets Everywhere field to given value.

### HasEverywhere

`func (o *ShippingZoneCreateDto) HasEverywhere() bool`

HasEverywhere returns a boolean if a field has been set.

### GetPostalCodes

`func (o *ShippingZoneCreateDto) GetPostalCodes() string`

GetPostalCodes returns the PostalCodes field if non-nil, zero value otherwise.

### GetPostalCodesOk

`func (o *ShippingZoneCreateDto) GetPostalCodesOk() (*string, bool)`

GetPostalCodesOk returns a tuple with the PostalCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCodes

`func (o *ShippingZoneCreateDto) SetPostalCodes(v string)`

SetPostalCodes sets PostalCodes field to given value.

### HasPostalCodes

`func (o *ShippingZoneCreateDto) HasPostalCodes() bool`

HasPostalCodes returns a boolean if a field has been set.

### SetPostalCodesNil

`func (o *ShippingZoneCreateDto) SetPostalCodesNil(b bool)`

 SetPostalCodesNil sets the value for PostalCodes to be an explicit nil

### UnsetPostalCodes
`func (o *ShippingZoneCreateDto) UnsetPostalCodes()`

UnsetPostalCodes ensures that no value is present for PostalCodes, not even an explicit nil
### GetCountryCodes

`func (o *ShippingZoneCreateDto) GetCountryCodes() string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *ShippingZoneCreateDto) GetCountryCodesOk() (*string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *ShippingZoneCreateDto) SetCountryCodes(v string)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *ShippingZoneCreateDto) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### SetCountryCodesNil

`func (o *ShippingZoneCreateDto) SetCountryCodesNil(b bool)`

 SetCountryCodesNil sets the value for CountryCodes to be an explicit nil

### UnsetCountryCodes
`func (o *ShippingZoneCreateDto) UnsetCountryCodes()`

UnsetCountryCodes ensures that no value is present for CountryCodes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ShippingZoneDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**Everywhere** | Pointer to **bool** |  | [optional] 
**PostalCodes** | Pointer to **NullableString** |  | [optional] 
**CountryCodes** | Pointer to **NullableString** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewShippingZoneDto

`func NewShippingZoneDto() *ShippingZoneDto`

NewShippingZoneDto instantiates a new ShippingZoneDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingZoneDtoWithDefaults

`func NewShippingZoneDtoWithDefaults() *ShippingZoneDto`

NewShippingZoneDtoWithDefaults instantiates a new ShippingZoneDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShippingZoneDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingZoneDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingZoneDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShippingZoneDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ShippingZoneDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ShippingZoneDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ShippingZoneDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ShippingZoneDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ShippingZoneDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ShippingZoneDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ShippingZoneDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ShippingZoneDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *ShippingZoneDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShippingZoneDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShippingZoneDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShippingZoneDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ShippingZoneDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ShippingZoneDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDefault

`func (o *ShippingZoneDto) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ShippingZoneDto) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ShippingZoneDto) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ShippingZoneDto) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEverywhere

`func (o *ShippingZoneDto) GetEverywhere() bool`

GetEverywhere returns the Everywhere field if non-nil, zero value otherwise.

### GetEverywhereOk

`func (o *ShippingZoneDto) GetEverywhereOk() (*bool, bool)`

GetEverywhereOk returns a tuple with the Everywhere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEverywhere

`func (o *ShippingZoneDto) SetEverywhere(v bool)`

SetEverywhere sets Everywhere field to given value.

### HasEverywhere

`func (o *ShippingZoneDto) HasEverywhere() bool`

HasEverywhere returns a boolean if a field has been set.

### GetPostalCodes

`func (o *ShippingZoneDto) GetPostalCodes() string`

GetPostalCodes returns the PostalCodes field if non-nil, zero value otherwise.

### GetPostalCodesOk

`func (o *ShippingZoneDto) GetPostalCodesOk() (*string, bool)`

GetPostalCodesOk returns a tuple with the PostalCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCodes

`func (o *ShippingZoneDto) SetPostalCodes(v string)`

SetPostalCodes sets PostalCodes field to given value.

### HasPostalCodes

`func (o *ShippingZoneDto) HasPostalCodes() bool`

HasPostalCodes returns a boolean if a field has been set.

### SetPostalCodesNil

`func (o *ShippingZoneDto) SetPostalCodesNil(b bool)`

 SetPostalCodesNil sets the value for PostalCodes to be an explicit nil

### UnsetPostalCodes
`func (o *ShippingZoneDto) UnsetPostalCodes()`

UnsetPostalCodes ensures that no value is present for PostalCodes, not even an explicit nil
### GetCountryCodes

`func (o *ShippingZoneDto) GetCountryCodes() string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *ShippingZoneDto) GetCountryCodesOk() (*string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *ShippingZoneDto) SetCountryCodes(v string)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *ShippingZoneDto) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### SetCountryCodesNil

`func (o *ShippingZoneDto) SetCountryCodesNil(b bool)`

 SetCountryCodesNil sets the value for CountryCodes to be an explicit nil

### UnsetCountryCodes
`func (o *ShippingZoneDto) UnsetCountryCodes()`

UnsetCountryCodes ensures that no value is present for CountryCodes, not even an explicit nil
### GetBusinessID

`func (o *ShippingZoneDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *ShippingZoneDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *ShippingZoneDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *ShippingZoneDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *ShippingZoneDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *ShippingZoneDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ShippingCourierCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**LogoURL** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewShippingCourierCreateDto

`func NewShippingCourierCreateDto(name string, ) *ShippingCourierCreateDto`

NewShippingCourierCreateDto instantiates a new ShippingCourierCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingCourierCreateDtoWithDefaults

`func NewShippingCourierCreateDtoWithDefaults() *ShippingCourierCreateDto`

NewShippingCourierCreateDtoWithDefaults instantiates a new ShippingCourierCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShippingCourierCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingCourierCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingCourierCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShippingCourierCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ShippingCourierCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ShippingCourierCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ShippingCourierCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ShippingCourierCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *ShippingCourierCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShippingCourierCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShippingCourierCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetLogoURL

`func (o *ShippingCourierCreateDto) GetLogoURL() string`

GetLogoURL returns the LogoURL field if non-nil, zero value otherwise.

### GetLogoURLOk

`func (o *ShippingCourierCreateDto) GetLogoURLOk() (*string, bool)`

GetLogoURLOk returns a tuple with the LogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoURL

`func (o *ShippingCourierCreateDto) SetLogoURL(v string)`

SetLogoURL sets LogoURL field to given value.

### HasLogoURL

`func (o *ShippingCourierCreateDto) HasLogoURL() bool`

HasLogoURL returns a boolean if a field has been set.

### SetLogoURLNil

`func (o *ShippingCourierCreateDto) SetLogoURLNil(b bool)`

 SetLogoURLNil sets the value for LogoURL to be an explicit nil

### UnsetLogoURL
`func (o *ShippingCourierCreateDto) UnsetLogoURL()`

UnsetLogoURL ensures that no value is present for LogoURL, not even an explicit nil
### GetCountryId

`func (o *ShippingCourierCreateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *ShippingCourierCreateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *ShippingCourierCreateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *ShippingCourierCreateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *ShippingCourierCreateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *ShippingCourierCreateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ShippingRegionCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**PostalCodes** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewShippingRegionCreateDto

`func NewShippingRegionCreateDto(name string, ) *ShippingRegionCreateDto`

NewShippingRegionCreateDto instantiates a new ShippingRegionCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingRegionCreateDtoWithDefaults

`func NewShippingRegionCreateDtoWithDefaults() *ShippingRegionCreateDto`

NewShippingRegionCreateDtoWithDefaults instantiates a new ShippingRegionCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShippingRegionCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingRegionCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingRegionCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShippingRegionCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ShippingRegionCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ShippingRegionCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ShippingRegionCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ShippingRegionCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *ShippingRegionCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShippingRegionCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShippingRegionCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetPostalCodes

`func (o *ShippingRegionCreateDto) GetPostalCodes() string`

GetPostalCodes returns the PostalCodes field if non-nil, zero value otherwise.

### GetPostalCodesOk

`func (o *ShippingRegionCreateDto) GetPostalCodesOk() (*string, bool)`

GetPostalCodesOk returns a tuple with the PostalCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCodes

`func (o *ShippingRegionCreateDto) SetPostalCodes(v string)`

SetPostalCodes sets PostalCodes field to given value.

### HasPostalCodes

`func (o *ShippingRegionCreateDto) HasPostalCodes() bool`

HasPostalCodes returns a boolean if a field has been set.

### SetPostalCodesNil

`func (o *ShippingRegionCreateDto) SetPostalCodesNil(b bool)`

 SetPostalCodesNil sets the value for PostalCodes to be an explicit nil

### UnsetPostalCodes
`func (o *ShippingRegionCreateDto) UnsetPostalCodes()`

UnsetPostalCodes ensures that no value is present for PostalCodes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



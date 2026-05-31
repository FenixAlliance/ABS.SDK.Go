# TruckCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**PlateNumber** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**TruckType** | Pointer to **NullableString** |  | [optional] 
**MaxPayloadKg** | Pointer to **NullableFloat64** |  | [optional] 
**TeuCapacity** | Pointer to **NullableInt32** |  | [optional] 
**DriverName** | Pointer to **NullableString** |  | [optional] 
**DriverPhone** | Pointer to **NullableString** |  | [optional] 
**DriverLicenseNumber** | Pointer to **NullableString** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**IsRefrigerated** | Pointer to **bool** |  | [optional] 
**ShippingCourierId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTruckCreateDto

`func NewTruckCreateDto() *TruckCreateDto`

NewTruckCreateDto instantiates a new TruckCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTruckCreateDtoWithDefaults

`func NewTruckCreateDtoWithDefaults() *TruckCreateDto`

NewTruckCreateDtoWithDefaults instantiates a new TruckCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TruckCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TruckCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TruckCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TruckCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TruckCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TruckCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TruckCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TruckCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetPlateNumber

`func (o *TruckCreateDto) GetPlateNumber() string`

GetPlateNumber returns the PlateNumber field if non-nil, zero value otherwise.

### GetPlateNumberOk

`func (o *TruckCreateDto) GetPlateNumberOk() (*string, bool)`

GetPlateNumberOk returns a tuple with the PlateNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlateNumber

`func (o *TruckCreateDto) SetPlateNumber(v string)`

SetPlateNumber sets PlateNumber field to given value.

### HasPlateNumber

`func (o *TruckCreateDto) HasPlateNumber() bool`

HasPlateNumber returns a boolean if a field has been set.

### SetPlateNumberNil

`func (o *TruckCreateDto) SetPlateNumberNil(b bool)`

 SetPlateNumberNil sets the value for PlateNumber to be an explicit nil

### UnsetPlateNumber
`func (o *TruckCreateDto) UnsetPlateNumber()`

UnsetPlateNumber ensures that no value is present for PlateNumber, not even an explicit nil
### GetName

`func (o *TruckCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TruckCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TruckCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TruckCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TruckCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TruckCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTruckType

`func (o *TruckCreateDto) GetTruckType() string`

GetTruckType returns the TruckType field if non-nil, zero value otherwise.

### GetTruckTypeOk

`func (o *TruckCreateDto) GetTruckTypeOk() (*string, bool)`

GetTruckTypeOk returns a tuple with the TruckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckType

`func (o *TruckCreateDto) SetTruckType(v string)`

SetTruckType sets TruckType field to given value.

### HasTruckType

`func (o *TruckCreateDto) HasTruckType() bool`

HasTruckType returns a boolean if a field has been set.

### SetTruckTypeNil

`func (o *TruckCreateDto) SetTruckTypeNil(b bool)`

 SetTruckTypeNil sets the value for TruckType to be an explicit nil

### UnsetTruckType
`func (o *TruckCreateDto) UnsetTruckType()`

UnsetTruckType ensures that no value is present for TruckType, not even an explicit nil
### GetMaxPayloadKg

`func (o *TruckCreateDto) GetMaxPayloadKg() float64`

GetMaxPayloadKg returns the MaxPayloadKg field if non-nil, zero value otherwise.

### GetMaxPayloadKgOk

`func (o *TruckCreateDto) GetMaxPayloadKgOk() (*float64, bool)`

GetMaxPayloadKgOk returns a tuple with the MaxPayloadKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPayloadKg

`func (o *TruckCreateDto) SetMaxPayloadKg(v float64)`

SetMaxPayloadKg sets MaxPayloadKg field to given value.

### HasMaxPayloadKg

`func (o *TruckCreateDto) HasMaxPayloadKg() bool`

HasMaxPayloadKg returns a boolean if a field has been set.

### SetMaxPayloadKgNil

`func (o *TruckCreateDto) SetMaxPayloadKgNil(b bool)`

 SetMaxPayloadKgNil sets the value for MaxPayloadKg to be an explicit nil

### UnsetMaxPayloadKg
`func (o *TruckCreateDto) UnsetMaxPayloadKg()`

UnsetMaxPayloadKg ensures that no value is present for MaxPayloadKg, not even an explicit nil
### GetTeuCapacity

`func (o *TruckCreateDto) GetTeuCapacity() int32`

GetTeuCapacity returns the TeuCapacity field if non-nil, zero value otherwise.

### GetTeuCapacityOk

`func (o *TruckCreateDto) GetTeuCapacityOk() (*int32, bool)`

GetTeuCapacityOk returns a tuple with the TeuCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeuCapacity

`func (o *TruckCreateDto) SetTeuCapacity(v int32)`

SetTeuCapacity sets TeuCapacity field to given value.

### HasTeuCapacity

`func (o *TruckCreateDto) HasTeuCapacity() bool`

HasTeuCapacity returns a boolean if a field has been set.

### SetTeuCapacityNil

`func (o *TruckCreateDto) SetTeuCapacityNil(b bool)`

 SetTeuCapacityNil sets the value for TeuCapacity to be an explicit nil

### UnsetTeuCapacity
`func (o *TruckCreateDto) UnsetTeuCapacity()`

UnsetTeuCapacity ensures that no value is present for TeuCapacity, not even an explicit nil
### GetDriverName

`func (o *TruckCreateDto) GetDriverName() string`

GetDriverName returns the DriverName field if non-nil, zero value otherwise.

### GetDriverNameOk

`func (o *TruckCreateDto) GetDriverNameOk() (*string, bool)`

GetDriverNameOk returns a tuple with the DriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverName

`func (o *TruckCreateDto) SetDriverName(v string)`

SetDriverName sets DriverName field to given value.

### HasDriverName

`func (o *TruckCreateDto) HasDriverName() bool`

HasDriverName returns a boolean if a field has been set.

### SetDriverNameNil

`func (o *TruckCreateDto) SetDriverNameNil(b bool)`

 SetDriverNameNil sets the value for DriverName to be an explicit nil

### UnsetDriverName
`func (o *TruckCreateDto) UnsetDriverName()`

UnsetDriverName ensures that no value is present for DriverName, not even an explicit nil
### GetDriverPhone

`func (o *TruckCreateDto) GetDriverPhone() string`

GetDriverPhone returns the DriverPhone field if non-nil, zero value otherwise.

### GetDriverPhoneOk

`func (o *TruckCreateDto) GetDriverPhoneOk() (*string, bool)`

GetDriverPhoneOk returns a tuple with the DriverPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverPhone

`func (o *TruckCreateDto) SetDriverPhone(v string)`

SetDriverPhone sets DriverPhone field to given value.

### HasDriverPhone

`func (o *TruckCreateDto) HasDriverPhone() bool`

HasDriverPhone returns a boolean if a field has been set.

### SetDriverPhoneNil

`func (o *TruckCreateDto) SetDriverPhoneNil(b bool)`

 SetDriverPhoneNil sets the value for DriverPhone to be an explicit nil

### UnsetDriverPhone
`func (o *TruckCreateDto) UnsetDriverPhone()`

UnsetDriverPhone ensures that no value is present for DriverPhone, not even an explicit nil
### GetDriverLicenseNumber

`func (o *TruckCreateDto) GetDriverLicenseNumber() string`

GetDriverLicenseNumber returns the DriverLicenseNumber field if non-nil, zero value otherwise.

### GetDriverLicenseNumberOk

`func (o *TruckCreateDto) GetDriverLicenseNumberOk() (*string, bool)`

GetDriverLicenseNumberOk returns a tuple with the DriverLicenseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseNumber

`func (o *TruckCreateDto) SetDriverLicenseNumber(v string)`

SetDriverLicenseNumber sets DriverLicenseNumber field to given value.

### HasDriverLicenseNumber

`func (o *TruckCreateDto) HasDriverLicenseNumber() bool`

HasDriverLicenseNumber returns a boolean if a field has been set.

### SetDriverLicenseNumberNil

`func (o *TruckCreateDto) SetDriverLicenseNumberNil(b bool)`

 SetDriverLicenseNumberNil sets the value for DriverLicenseNumber to be an explicit nil

### UnsetDriverLicenseNumber
`func (o *TruckCreateDto) UnsetDriverLicenseNumber()`

UnsetDriverLicenseNumber ensures that no value is present for DriverLicenseNumber, not even an explicit nil
### GetIsActive

`func (o *TruckCreateDto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *TruckCreateDto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *TruckCreateDto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *TruckCreateDto) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsRefrigerated

`func (o *TruckCreateDto) GetIsRefrigerated() bool`

GetIsRefrigerated returns the IsRefrigerated field if non-nil, zero value otherwise.

### GetIsRefrigeratedOk

`func (o *TruckCreateDto) GetIsRefrigeratedOk() (*bool, bool)`

GetIsRefrigeratedOk returns a tuple with the IsRefrigerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRefrigerated

`func (o *TruckCreateDto) SetIsRefrigerated(v bool)`

SetIsRefrigerated sets IsRefrigerated field to given value.

### HasIsRefrigerated

`func (o *TruckCreateDto) HasIsRefrigerated() bool`

HasIsRefrigerated returns a boolean if a field has been set.

### GetShippingCourierId

`func (o *TruckCreateDto) GetShippingCourierId() string`

GetShippingCourierId returns the ShippingCourierId field if non-nil, zero value otherwise.

### GetShippingCourierIdOk

`func (o *TruckCreateDto) GetShippingCourierIdOk() (*string, bool)`

GetShippingCourierIdOk returns a tuple with the ShippingCourierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCourierId

`func (o *TruckCreateDto) SetShippingCourierId(v string)`

SetShippingCourierId sets ShippingCourierId field to given value.

### HasShippingCourierId

`func (o *TruckCreateDto) HasShippingCourierId() bool`

HasShippingCourierId returns a boolean if a field has been set.

### SetShippingCourierIdNil

`func (o *TruckCreateDto) SetShippingCourierIdNil(b bool)`

 SetShippingCourierIdNil sets the value for ShippingCourierId to be an explicit nil

### UnsetShippingCourierId
`func (o *TruckCreateDto) UnsetShippingCourierId()`

UnsetShippingCourierId ensures that no value is present for ShippingCourierId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



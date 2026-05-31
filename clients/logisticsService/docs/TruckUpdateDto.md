# TruckUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewTruckUpdateDto

`func NewTruckUpdateDto() *TruckUpdateDto`

NewTruckUpdateDto instantiates a new TruckUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTruckUpdateDtoWithDefaults

`func NewTruckUpdateDtoWithDefaults() *TruckUpdateDto`

NewTruckUpdateDtoWithDefaults instantiates a new TruckUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlateNumber

`func (o *TruckUpdateDto) GetPlateNumber() string`

GetPlateNumber returns the PlateNumber field if non-nil, zero value otherwise.

### GetPlateNumberOk

`func (o *TruckUpdateDto) GetPlateNumberOk() (*string, bool)`

GetPlateNumberOk returns a tuple with the PlateNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlateNumber

`func (o *TruckUpdateDto) SetPlateNumber(v string)`

SetPlateNumber sets PlateNumber field to given value.

### HasPlateNumber

`func (o *TruckUpdateDto) HasPlateNumber() bool`

HasPlateNumber returns a boolean if a field has been set.

### SetPlateNumberNil

`func (o *TruckUpdateDto) SetPlateNumberNil(b bool)`

 SetPlateNumberNil sets the value for PlateNumber to be an explicit nil

### UnsetPlateNumber
`func (o *TruckUpdateDto) UnsetPlateNumber()`

UnsetPlateNumber ensures that no value is present for PlateNumber, not even an explicit nil
### GetName

`func (o *TruckUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TruckUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TruckUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TruckUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TruckUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TruckUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTruckType

`func (o *TruckUpdateDto) GetTruckType() string`

GetTruckType returns the TruckType field if non-nil, zero value otherwise.

### GetTruckTypeOk

`func (o *TruckUpdateDto) GetTruckTypeOk() (*string, bool)`

GetTruckTypeOk returns a tuple with the TruckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckType

`func (o *TruckUpdateDto) SetTruckType(v string)`

SetTruckType sets TruckType field to given value.

### HasTruckType

`func (o *TruckUpdateDto) HasTruckType() bool`

HasTruckType returns a boolean if a field has been set.

### SetTruckTypeNil

`func (o *TruckUpdateDto) SetTruckTypeNil(b bool)`

 SetTruckTypeNil sets the value for TruckType to be an explicit nil

### UnsetTruckType
`func (o *TruckUpdateDto) UnsetTruckType()`

UnsetTruckType ensures that no value is present for TruckType, not even an explicit nil
### GetMaxPayloadKg

`func (o *TruckUpdateDto) GetMaxPayloadKg() float64`

GetMaxPayloadKg returns the MaxPayloadKg field if non-nil, zero value otherwise.

### GetMaxPayloadKgOk

`func (o *TruckUpdateDto) GetMaxPayloadKgOk() (*float64, bool)`

GetMaxPayloadKgOk returns a tuple with the MaxPayloadKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPayloadKg

`func (o *TruckUpdateDto) SetMaxPayloadKg(v float64)`

SetMaxPayloadKg sets MaxPayloadKg field to given value.

### HasMaxPayloadKg

`func (o *TruckUpdateDto) HasMaxPayloadKg() bool`

HasMaxPayloadKg returns a boolean if a field has been set.

### SetMaxPayloadKgNil

`func (o *TruckUpdateDto) SetMaxPayloadKgNil(b bool)`

 SetMaxPayloadKgNil sets the value for MaxPayloadKg to be an explicit nil

### UnsetMaxPayloadKg
`func (o *TruckUpdateDto) UnsetMaxPayloadKg()`

UnsetMaxPayloadKg ensures that no value is present for MaxPayloadKg, not even an explicit nil
### GetTeuCapacity

`func (o *TruckUpdateDto) GetTeuCapacity() int32`

GetTeuCapacity returns the TeuCapacity field if non-nil, zero value otherwise.

### GetTeuCapacityOk

`func (o *TruckUpdateDto) GetTeuCapacityOk() (*int32, bool)`

GetTeuCapacityOk returns a tuple with the TeuCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeuCapacity

`func (o *TruckUpdateDto) SetTeuCapacity(v int32)`

SetTeuCapacity sets TeuCapacity field to given value.

### HasTeuCapacity

`func (o *TruckUpdateDto) HasTeuCapacity() bool`

HasTeuCapacity returns a boolean if a field has been set.

### SetTeuCapacityNil

`func (o *TruckUpdateDto) SetTeuCapacityNil(b bool)`

 SetTeuCapacityNil sets the value for TeuCapacity to be an explicit nil

### UnsetTeuCapacity
`func (o *TruckUpdateDto) UnsetTeuCapacity()`

UnsetTeuCapacity ensures that no value is present for TeuCapacity, not even an explicit nil
### GetDriverName

`func (o *TruckUpdateDto) GetDriverName() string`

GetDriverName returns the DriverName field if non-nil, zero value otherwise.

### GetDriverNameOk

`func (o *TruckUpdateDto) GetDriverNameOk() (*string, bool)`

GetDriverNameOk returns a tuple with the DriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverName

`func (o *TruckUpdateDto) SetDriverName(v string)`

SetDriverName sets DriverName field to given value.

### HasDriverName

`func (o *TruckUpdateDto) HasDriverName() bool`

HasDriverName returns a boolean if a field has been set.

### SetDriverNameNil

`func (o *TruckUpdateDto) SetDriverNameNil(b bool)`

 SetDriverNameNil sets the value for DriverName to be an explicit nil

### UnsetDriverName
`func (o *TruckUpdateDto) UnsetDriverName()`

UnsetDriverName ensures that no value is present for DriverName, not even an explicit nil
### GetDriverPhone

`func (o *TruckUpdateDto) GetDriverPhone() string`

GetDriverPhone returns the DriverPhone field if non-nil, zero value otherwise.

### GetDriverPhoneOk

`func (o *TruckUpdateDto) GetDriverPhoneOk() (*string, bool)`

GetDriverPhoneOk returns a tuple with the DriverPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverPhone

`func (o *TruckUpdateDto) SetDriverPhone(v string)`

SetDriverPhone sets DriverPhone field to given value.

### HasDriverPhone

`func (o *TruckUpdateDto) HasDriverPhone() bool`

HasDriverPhone returns a boolean if a field has been set.

### SetDriverPhoneNil

`func (o *TruckUpdateDto) SetDriverPhoneNil(b bool)`

 SetDriverPhoneNil sets the value for DriverPhone to be an explicit nil

### UnsetDriverPhone
`func (o *TruckUpdateDto) UnsetDriverPhone()`

UnsetDriverPhone ensures that no value is present for DriverPhone, not even an explicit nil
### GetDriverLicenseNumber

`func (o *TruckUpdateDto) GetDriverLicenseNumber() string`

GetDriverLicenseNumber returns the DriverLicenseNumber field if non-nil, zero value otherwise.

### GetDriverLicenseNumberOk

`func (o *TruckUpdateDto) GetDriverLicenseNumberOk() (*string, bool)`

GetDriverLicenseNumberOk returns a tuple with the DriverLicenseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseNumber

`func (o *TruckUpdateDto) SetDriverLicenseNumber(v string)`

SetDriverLicenseNumber sets DriverLicenseNumber field to given value.

### HasDriverLicenseNumber

`func (o *TruckUpdateDto) HasDriverLicenseNumber() bool`

HasDriverLicenseNumber returns a boolean if a field has been set.

### SetDriverLicenseNumberNil

`func (o *TruckUpdateDto) SetDriverLicenseNumberNil(b bool)`

 SetDriverLicenseNumberNil sets the value for DriverLicenseNumber to be an explicit nil

### UnsetDriverLicenseNumber
`func (o *TruckUpdateDto) UnsetDriverLicenseNumber()`

UnsetDriverLicenseNumber ensures that no value is present for DriverLicenseNumber, not even an explicit nil
### GetIsActive

`func (o *TruckUpdateDto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *TruckUpdateDto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *TruckUpdateDto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *TruckUpdateDto) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsRefrigerated

`func (o *TruckUpdateDto) GetIsRefrigerated() bool`

GetIsRefrigerated returns the IsRefrigerated field if non-nil, zero value otherwise.

### GetIsRefrigeratedOk

`func (o *TruckUpdateDto) GetIsRefrigeratedOk() (*bool, bool)`

GetIsRefrigeratedOk returns a tuple with the IsRefrigerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRefrigerated

`func (o *TruckUpdateDto) SetIsRefrigerated(v bool)`

SetIsRefrigerated sets IsRefrigerated field to given value.

### HasIsRefrigerated

`func (o *TruckUpdateDto) HasIsRefrigerated() bool`

HasIsRefrigerated returns a boolean if a field has been set.

### GetShippingCourierId

`func (o *TruckUpdateDto) GetShippingCourierId() string`

GetShippingCourierId returns the ShippingCourierId field if non-nil, zero value otherwise.

### GetShippingCourierIdOk

`func (o *TruckUpdateDto) GetShippingCourierIdOk() (*string, bool)`

GetShippingCourierIdOk returns a tuple with the ShippingCourierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCourierId

`func (o *TruckUpdateDto) SetShippingCourierId(v string)`

SetShippingCourierId sets ShippingCourierId field to given value.

### HasShippingCourierId

`func (o *TruckUpdateDto) HasShippingCourierId() bool`

HasShippingCourierId returns a boolean if a field has been set.

### SetShippingCourierIdNil

`func (o *TruckUpdateDto) SetShippingCourierIdNil(b bool)`

 SetShippingCourierIdNil sets the value for ShippingCourierId to be an explicit nil

### UnsetShippingCourierId
`func (o *TruckUpdateDto) UnsetShippingCourierId()`

UnsetShippingCourierId ensures that no value is present for ShippingCourierId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



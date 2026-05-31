# TruckDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
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
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTruckDto

`func NewTruckDto() *TruckDto`

NewTruckDto instantiates a new TruckDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTruckDtoWithDefaults

`func NewTruckDtoWithDefaults() *TruckDto`

NewTruckDtoWithDefaults instantiates a new TruckDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TruckDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TruckDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TruckDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TruckDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TruckDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TruckDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *TruckDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TruckDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TruckDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TruckDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TruckDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TruckDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetPlateNumber

`func (o *TruckDto) GetPlateNumber() string`

GetPlateNumber returns the PlateNumber field if non-nil, zero value otherwise.

### GetPlateNumberOk

`func (o *TruckDto) GetPlateNumberOk() (*string, bool)`

GetPlateNumberOk returns a tuple with the PlateNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlateNumber

`func (o *TruckDto) SetPlateNumber(v string)`

SetPlateNumber sets PlateNumber field to given value.

### HasPlateNumber

`func (o *TruckDto) HasPlateNumber() bool`

HasPlateNumber returns a boolean if a field has been set.

### SetPlateNumberNil

`func (o *TruckDto) SetPlateNumberNil(b bool)`

 SetPlateNumberNil sets the value for PlateNumber to be an explicit nil

### UnsetPlateNumber
`func (o *TruckDto) UnsetPlateNumber()`

UnsetPlateNumber ensures that no value is present for PlateNumber, not even an explicit nil
### GetName

`func (o *TruckDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TruckDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TruckDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TruckDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TruckDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TruckDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTruckType

`func (o *TruckDto) GetTruckType() string`

GetTruckType returns the TruckType field if non-nil, zero value otherwise.

### GetTruckTypeOk

`func (o *TruckDto) GetTruckTypeOk() (*string, bool)`

GetTruckTypeOk returns a tuple with the TruckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckType

`func (o *TruckDto) SetTruckType(v string)`

SetTruckType sets TruckType field to given value.

### HasTruckType

`func (o *TruckDto) HasTruckType() bool`

HasTruckType returns a boolean if a field has been set.

### SetTruckTypeNil

`func (o *TruckDto) SetTruckTypeNil(b bool)`

 SetTruckTypeNil sets the value for TruckType to be an explicit nil

### UnsetTruckType
`func (o *TruckDto) UnsetTruckType()`

UnsetTruckType ensures that no value is present for TruckType, not even an explicit nil
### GetMaxPayloadKg

`func (o *TruckDto) GetMaxPayloadKg() float64`

GetMaxPayloadKg returns the MaxPayloadKg field if non-nil, zero value otherwise.

### GetMaxPayloadKgOk

`func (o *TruckDto) GetMaxPayloadKgOk() (*float64, bool)`

GetMaxPayloadKgOk returns a tuple with the MaxPayloadKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPayloadKg

`func (o *TruckDto) SetMaxPayloadKg(v float64)`

SetMaxPayloadKg sets MaxPayloadKg field to given value.

### HasMaxPayloadKg

`func (o *TruckDto) HasMaxPayloadKg() bool`

HasMaxPayloadKg returns a boolean if a field has been set.

### SetMaxPayloadKgNil

`func (o *TruckDto) SetMaxPayloadKgNil(b bool)`

 SetMaxPayloadKgNil sets the value for MaxPayloadKg to be an explicit nil

### UnsetMaxPayloadKg
`func (o *TruckDto) UnsetMaxPayloadKg()`

UnsetMaxPayloadKg ensures that no value is present for MaxPayloadKg, not even an explicit nil
### GetTeuCapacity

`func (o *TruckDto) GetTeuCapacity() int32`

GetTeuCapacity returns the TeuCapacity field if non-nil, zero value otherwise.

### GetTeuCapacityOk

`func (o *TruckDto) GetTeuCapacityOk() (*int32, bool)`

GetTeuCapacityOk returns a tuple with the TeuCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeuCapacity

`func (o *TruckDto) SetTeuCapacity(v int32)`

SetTeuCapacity sets TeuCapacity field to given value.

### HasTeuCapacity

`func (o *TruckDto) HasTeuCapacity() bool`

HasTeuCapacity returns a boolean if a field has been set.

### SetTeuCapacityNil

`func (o *TruckDto) SetTeuCapacityNil(b bool)`

 SetTeuCapacityNil sets the value for TeuCapacity to be an explicit nil

### UnsetTeuCapacity
`func (o *TruckDto) UnsetTeuCapacity()`

UnsetTeuCapacity ensures that no value is present for TeuCapacity, not even an explicit nil
### GetDriverName

`func (o *TruckDto) GetDriverName() string`

GetDriverName returns the DriverName field if non-nil, zero value otherwise.

### GetDriverNameOk

`func (o *TruckDto) GetDriverNameOk() (*string, bool)`

GetDriverNameOk returns a tuple with the DriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverName

`func (o *TruckDto) SetDriverName(v string)`

SetDriverName sets DriverName field to given value.

### HasDriverName

`func (o *TruckDto) HasDriverName() bool`

HasDriverName returns a boolean if a field has been set.

### SetDriverNameNil

`func (o *TruckDto) SetDriverNameNil(b bool)`

 SetDriverNameNil sets the value for DriverName to be an explicit nil

### UnsetDriverName
`func (o *TruckDto) UnsetDriverName()`

UnsetDriverName ensures that no value is present for DriverName, not even an explicit nil
### GetDriverPhone

`func (o *TruckDto) GetDriverPhone() string`

GetDriverPhone returns the DriverPhone field if non-nil, zero value otherwise.

### GetDriverPhoneOk

`func (o *TruckDto) GetDriverPhoneOk() (*string, bool)`

GetDriverPhoneOk returns a tuple with the DriverPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverPhone

`func (o *TruckDto) SetDriverPhone(v string)`

SetDriverPhone sets DriverPhone field to given value.

### HasDriverPhone

`func (o *TruckDto) HasDriverPhone() bool`

HasDriverPhone returns a boolean if a field has been set.

### SetDriverPhoneNil

`func (o *TruckDto) SetDriverPhoneNil(b bool)`

 SetDriverPhoneNil sets the value for DriverPhone to be an explicit nil

### UnsetDriverPhone
`func (o *TruckDto) UnsetDriverPhone()`

UnsetDriverPhone ensures that no value is present for DriverPhone, not even an explicit nil
### GetDriverLicenseNumber

`func (o *TruckDto) GetDriverLicenseNumber() string`

GetDriverLicenseNumber returns the DriverLicenseNumber field if non-nil, zero value otherwise.

### GetDriverLicenseNumberOk

`func (o *TruckDto) GetDriverLicenseNumberOk() (*string, bool)`

GetDriverLicenseNumberOk returns a tuple with the DriverLicenseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseNumber

`func (o *TruckDto) SetDriverLicenseNumber(v string)`

SetDriverLicenseNumber sets DriverLicenseNumber field to given value.

### HasDriverLicenseNumber

`func (o *TruckDto) HasDriverLicenseNumber() bool`

HasDriverLicenseNumber returns a boolean if a field has been set.

### SetDriverLicenseNumberNil

`func (o *TruckDto) SetDriverLicenseNumberNil(b bool)`

 SetDriverLicenseNumberNil sets the value for DriverLicenseNumber to be an explicit nil

### UnsetDriverLicenseNumber
`func (o *TruckDto) UnsetDriverLicenseNumber()`

UnsetDriverLicenseNumber ensures that no value is present for DriverLicenseNumber, not even an explicit nil
### GetIsActive

`func (o *TruckDto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *TruckDto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *TruckDto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *TruckDto) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsRefrigerated

`func (o *TruckDto) GetIsRefrigerated() bool`

GetIsRefrigerated returns the IsRefrigerated field if non-nil, zero value otherwise.

### GetIsRefrigeratedOk

`func (o *TruckDto) GetIsRefrigeratedOk() (*bool, bool)`

GetIsRefrigeratedOk returns a tuple with the IsRefrigerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRefrigerated

`func (o *TruckDto) SetIsRefrigerated(v bool)`

SetIsRefrigerated sets IsRefrigerated field to given value.

### HasIsRefrigerated

`func (o *TruckDto) HasIsRefrigerated() bool`

HasIsRefrigerated returns a boolean if a field has been set.

### GetShippingCourierId

`func (o *TruckDto) GetShippingCourierId() string`

GetShippingCourierId returns the ShippingCourierId field if non-nil, zero value otherwise.

### GetShippingCourierIdOk

`func (o *TruckDto) GetShippingCourierIdOk() (*string, bool)`

GetShippingCourierIdOk returns a tuple with the ShippingCourierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCourierId

`func (o *TruckDto) SetShippingCourierId(v string)`

SetShippingCourierId sets ShippingCourierId field to given value.

### HasShippingCourierId

`func (o *TruckDto) HasShippingCourierId() bool`

HasShippingCourierId returns a boolean if a field has been set.

### SetShippingCourierIdNil

`func (o *TruckDto) SetShippingCourierIdNil(b bool)`

 SetShippingCourierIdNil sets the value for ShippingCourierId to be an explicit nil

### UnsetShippingCourierId
`func (o *TruckDto) UnsetShippingCourierId()`

UnsetShippingCourierId ensures that no value is present for ShippingCourierId, not even an explicit nil
### GetTenantId

`func (o *TruckDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TruckDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TruckDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TruckDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TruckDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TruckDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *TruckDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *TruckDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *TruckDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *TruckDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *TruckDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *TruckDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



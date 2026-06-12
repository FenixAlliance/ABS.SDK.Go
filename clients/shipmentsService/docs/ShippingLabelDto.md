# ShippingLabelDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TrackingCode** | Pointer to **NullableString** |  | [optional] 
**ExpectedDelivery** | Pointer to **time.Time** |  | [optional] 
**LocationId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**ShipmentId** | Pointer to **NullableString** |  | [optional] 
**ShippingCourierId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewShippingLabelDto

`func NewShippingLabelDto() *ShippingLabelDto`

NewShippingLabelDto instantiates a new ShippingLabelDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingLabelDtoWithDefaults

`func NewShippingLabelDtoWithDefaults() *ShippingLabelDto`

NewShippingLabelDtoWithDefaults instantiates a new ShippingLabelDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShippingLabelDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingLabelDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingLabelDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShippingLabelDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ShippingLabelDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ShippingLabelDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ShippingLabelDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ShippingLabelDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ShippingLabelDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ShippingLabelDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ShippingLabelDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ShippingLabelDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTrackingCode

`func (o *ShippingLabelDto) GetTrackingCode() string`

GetTrackingCode returns the TrackingCode field if non-nil, zero value otherwise.

### GetTrackingCodeOk

`func (o *ShippingLabelDto) GetTrackingCodeOk() (*string, bool)`

GetTrackingCodeOk returns a tuple with the TrackingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCode

`func (o *ShippingLabelDto) SetTrackingCode(v string)`

SetTrackingCode sets TrackingCode field to given value.

### HasTrackingCode

`func (o *ShippingLabelDto) HasTrackingCode() bool`

HasTrackingCode returns a boolean if a field has been set.

### SetTrackingCodeNil

`func (o *ShippingLabelDto) SetTrackingCodeNil(b bool)`

 SetTrackingCodeNil sets the value for TrackingCode to be an explicit nil

### UnsetTrackingCode
`func (o *ShippingLabelDto) UnsetTrackingCode()`

UnsetTrackingCode ensures that no value is present for TrackingCode, not even an explicit nil
### GetExpectedDelivery

`func (o *ShippingLabelDto) GetExpectedDelivery() time.Time`

GetExpectedDelivery returns the ExpectedDelivery field if non-nil, zero value otherwise.

### GetExpectedDeliveryOk

`func (o *ShippingLabelDto) GetExpectedDeliveryOk() (*time.Time, bool)`

GetExpectedDeliveryOk returns a tuple with the ExpectedDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDelivery

`func (o *ShippingLabelDto) SetExpectedDelivery(v time.Time)`

SetExpectedDelivery sets ExpectedDelivery field to given value.

### HasExpectedDelivery

`func (o *ShippingLabelDto) HasExpectedDelivery() bool`

HasExpectedDelivery returns a boolean if a field has been set.

### GetLocationId

`func (o *ShippingLabelDto) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *ShippingLabelDto) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *ShippingLabelDto) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *ShippingLabelDto) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *ShippingLabelDto) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *ShippingLabelDto) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetTenantId

`func (o *ShippingLabelDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ShippingLabelDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ShippingLabelDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ShippingLabelDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ShippingLabelDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ShippingLabelDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetShipmentId

`func (o *ShippingLabelDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *ShippingLabelDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *ShippingLabelDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *ShippingLabelDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *ShippingLabelDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *ShippingLabelDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil
### GetShippingCourierId

`func (o *ShippingLabelDto) GetShippingCourierId() string`

GetShippingCourierId returns the ShippingCourierId field if non-nil, zero value otherwise.

### GetShippingCourierIdOk

`func (o *ShippingLabelDto) GetShippingCourierIdOk() (*string, bool)`

GetShippingCourierIdOk returns a tuple with the ShippingCourierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCourierId

`func (o *ShippingLabelDto) SetShippingCourierId(v string)`

SetShippingCourierId sets ShippingCourierId field to given value.

### HasShippingCourierId

`func (o *ShippingLabelDto) HasShippingCourierId() bool`

HasShippingCourierId returns a boolean if a field has been set.

### SetShippingCourierIdNil

`func (o *ShippingLabelDto) SetShippingCourierIdNil(b bool)`

 SetShippingCourierIdNil sets the value for ShippingCourierId to be an explicit nil

### UnsetShippingCourierId
`func (o *ShippingLabelDto) UnsetShippingCourierId()`

UnsetShippingCourierId ensures that no value is present for ShippingCourierId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



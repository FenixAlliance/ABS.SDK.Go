# ShippingLabelCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**TrackingCode** | **string** |  | 
**ExpectedDelivery** | Pointer to **time.Time** |  | [optional] 
**LocationId** | Pointer to **NullableString** |  | [optional] 
**ShipmentId** | Pointer to **NullableString** |  | [optional] 
**ShippingCourierId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewShippingLabelCreateDto

`func NewShippingLabelCreateDto(trackingCode string, ) *ShippingLabelCreateDto`

NewShippingLabelCreateDto instantiates a new ShippingLabelCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingLabelCreateDtoWithDefaults

`func NewShippingLabelCreateDtoWithDefaults() *ShippingLabelCreateDto`

NewShippingLabelCreateDtoWithDefaults instantiates a new ShippingLabelCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShippingLabelCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingLabelCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingLabelCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShippingLabelCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ShippingLabelCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ShippingLabelCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ShippingLabelCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ShippingLabelCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTrackingCode

`func (o *ShippingLabelCreateDto) GetTrackingCode() string`

GetTrackingCode returns the TrackingCode field if non-nil, zero value otherwise.

### GetTrackingCodeOk

`func (o *ShippingLabelCreateDto) GetTrackingCodeOk() (*string, bool)`

GetTrackingCodeOk returns a tuple with the TrackingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCode

`func (o *ShippingLabelCreateDto) SetTrackingCode(v string)`

SetTrackingCode sets TrackingCode field to given value.


### GetExpectedDelivery

`func (o *ShippingLabelCreateDto) GetExpectedDelivery() time.Time`

GetExpectedDelivery returns the ExpectedDelivery field if non-nil, zero value otherwise.

### GetExpectedDeliveryOk

`func (o *ShippingLabelCreateDto) GetExpectedDeliveryOk() (*time.Time, bool)`

GetExpectedDeliveryOk returns a tuple with the ExpectedDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDelivery

`func (o *ShippingLabelCreateDto) SetExpectedDelivery(v time.Time)`

SetExpectedDelivery sets ExpectedDelivery field to given value.

### HasExpectedDelivery

`func (o *ShippingLabelCreateDto) HasExpectedDelivery() bool`

HasExpectedDelivery returns a boolean if a field has been set.

### GetLocationId

`func (o *ShippingLabelCreateDto) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *ShippingLabelCreateDto) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *ShippingLabelCreateDto) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *ShippingLabelCreateDto) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *ShippingLabelCreateDto) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *ShippingLabelCreateDto) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetShipmentId

`func (o *ShippingLabelCreateDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *ShippingLabelCreateDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *ShippingLabelCreateDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *ShippingLabelCreateDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *ShippingLabelCreateDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *ShippingLabelCreateDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil
### GetShippingCourierId

`func (o *ShippingLabelCreateDto) GetShippingCourierId() string`

GetShippingCourierId returns the ShippingCourierId field if non-nil, zero value otherwise.

### GetShippingCourierIdOk

`func (o *ShippingLabelCreateDto) GetShippingCourierIdOk() (*string, bool)`

GetShippingCourierIdOk returns a tuple with the ShippingCourierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCourierId

`func (o *ShippingLabelCreateDto) SetShippingCourierId(v string)`

SetShippingCourierId sets ShippingCourierId field to given value.

### HasShippingCourierId

`func (o *ShippingLabelCreateDto) HasShippingCourierId() bool`

HasShippingCourierId returns a boolean if a field has been set.

### SetShippingCourierIdNil

`func (o *ShippingLabelCreateDto) SetShippingCourierIdNil(b bool)`

 SetShippingCourierIdNil sets the value for ShippingCourierId to be an explicit nil

### UnsetShippingCourierId
`func (o *ShippingLabelCreateDto) UnsetShippingCourierId()`

UnsetShippingCourierId ensures that no value is present for ShippingCourierId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



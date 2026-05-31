# ShippingLabelDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TrackingCode** | Pointer to **NullableString** |  | [optional] 
**ExpectedDelivery** | Pointer to **time.Time** |  | [optional] 
**LocationID** | Pointer to **NullableString** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**ShipmentID** | Pointer to **NullableString** |  | [optional] 
**ShippingCourierID** | Pointer to **NullableString** |  | [optional] 

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

### GetLocationID

`func (o *ShippingLabelDto) GetLocationID() string`

GetLocationID returns the LocationID field if non-nil, zero value otherwise.

### GetLocationIDOk

`func (o *ShippingLabelDto) GetLocationIDOk() (*string, bool)`

GetLocationIDOk returns a tuple with the LocationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationID

`func (o *ShippingLabelDto) SetLocationID(v string)`

SetLocationID sets LocationID field to given value.

### HasLocationID

`func (o *ShippingLabelDto) HasLocationID() bool`

HasLocationID returns a boolean if a field has been set.

### SetLocationIDNil

`func (o *ShippingLabelDto) SetLocationIDNil(b bool)`

 SetLocationIDNil sets the value for LocationID to be an explicit nil

### UnsetLocationID
`func (o *ShippingLabelDto) UnsetLocationID()`

UnsetLocationID ensures that no value is present for LocationID, not even an explicit nil
### GetBusinessID

`func (o *ShippingLabelDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *ShippingLabelDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *ShippingLabelDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *ShippingLabelDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *ShippingLabelDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *ShippingLabelDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetShipmentID

`func (o *ShippingLabelDto) GetShipmentID() string`

GetShipmentID returns the ShipmentID field if non-nil, zero value otherwise.

### GetShipmentIDOk

`func (o *ShippingLabelDto) GetShipmentIDOk() (*string, bool)`

GetShipmentIDOk returns a tuple with the ShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentID

`func (o *ShippingLabelDto) SetShipmentID(v string)`

SetShipmentID sets ShipmentID field to given value.

### HasShipmentID

`func (o *ShippingLabelDto) HasShipmentID() bool`

HasShipmentID returns a boolean if a field has been set.

### SetShipmentIDNil

`func (o *ShippingLabelDto) SetShipmentIDNil(b bool)`

 SetShipmentIDNil sets the value for ShipmentID to be an explicit nil

### UnsetShipmentID
`func (o *ShippingLabelDto) UnsetShipmentID()`

UnsetShipmentID ensures that no value is present for ShipmentID, not even an explicit nil
### GetShippingCourierID

`func (o *ShippingLabelDto) GetShippingCourierID() string`

GetShippingCourierID returns the ShippingCourierID field if non-nil, zero value otherwise.

### GetShippingCourierIDOk

`func (o *ShippingLabelDto) GetShippingCourierIDOk() (*string, bool)`

GetShippingCourierIDOk returns a tuple with the ShippingCourierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCourierID

`func (o *ShippingLabelDto) SetShippingCourierID(v string)`

SetShippingCourierID sets ShippingCourierID field to given value.

### HasShippingCourierID

`func (o *ShippingLabelDto) HasShippingCourierID() bool`

HasShippingCourierID returns a boolean if a field has been set.

### SetShippingCourierIDNil

`func (o *ShippingLabelDto) SetShippingCourierIDNil(b bool)`

 SetShippingCourierIDNil sets the value for ShippingCourierID to be an explicit nil

### UnsetShippingCourierID
`func (o *ShippingLabelDto) UnsetShippingCourierID()`

UnsetShippingCourierID ensures that no value is present for ShippingCourierID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



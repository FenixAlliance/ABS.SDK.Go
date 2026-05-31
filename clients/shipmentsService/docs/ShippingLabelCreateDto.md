# ShippingLabelCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**TrackingCode** | **string** |  | 
**ExpectedDelivery** | Pointer to **time.Time** |  | [optional] 
**LocationID** | Pointer to **NullableString** |  | [optional] 
**ShipmentID** | Pointer to **NullableString** |  | [optional] 
**ShippingCourierID** | Pointer to **NullableString** |  | [optional] 

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

### GetLocationID

`func (o *ShippingLabelCreateDto) GetLocationID() string`

GetLocationID returns the LocationID field if non-nil, zero value otherwise.

### GetLocationIDOk

`func (o *ShippingLabelCreateDto) GetLocationIDOk() (*string, bool)`

GetLocationIDOk returns a tuple with the LocationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationID

`func (o *ShippingLabelCreateDto) SetLocationID(v string)`

SetLocationID sets LocationID field to given value.

### HasLocationID

`func (o *ShippingLabelCreateDto) HasLocationID() bool`

HasLocationID returns a boolean if a field has been set.

### SetLocationIDNil

`func (o *ShippingLabelCreateDto) SetLocationIDNil(b bool)`

 SetLocationIDNil sets the value for LocationID to be an explicit nil

### UnsetLocationID
`func (o *ShippingLabelCreateDto) UnsetLocationID()`

UnsetLocationID ensures that no value is present for LocationID, not even an explicit nil
### GetShipmentID

`func (o *ShippingLabelCreateDto) GetShipmentID() string`

GetShipmentID returns the ShipmentID field if non-nil, zero value otherwise.

### GetShipmentIDOk

`func (o *ShippingLabelCreateDto) GetShipmentIDOk() (*string, bool)`

GetShipmentIDOk returns a tuple with the ShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentID

`func (o *ShippingLabelCreateDto) SetShipmentID(v string)`

SetShipmentID sets ShipmentID field to given value.

### HasShipmentID

`func (o *ShippingLabelCreateDto) HasShipmentID() bool`

HasShipmentID returns a boolean if a field has been set.

### SetShipmentIDNil

`func (o *ShippingLabelCreateDto) SetShipmentIDNil(b bool)`

 SetShipmentIDNil sets the value for ShipmentID to be an explicit nil

### UnsetShipmentID
`func (o *ShippingLabelCreateDto) UnsetShipmentID()`

UnsetShipmentID ensures that no value is present for ShipmentID, not even an explicit nil
### GetShippingCourierID

`func (o *ShippingLabelCreateDto) GetShippingCourierID() string`

GetShippingCourierID returns the ShippingCourierID field if non-nil, zero value otherwise.

### GetShippingCourierIDOk

`func (o *ShippingLabelCreateDto) GetShippingCourierIDOk() (*string, bool)`

GetShippingCourierIDOk returns a tuple with the ShippingCourierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCourierID

`func (o *ShippingLabelCreateDto) SetShippingCourierID(v string)`

SetShippingCourierID sets ShippingCourierID field to given value.

### HasShippingCourierID

`func (o *ShippingLabelCreateDto) HasShippingCourierID() bool`

HasShippingCourierID returns a boolean if a field has been set.

### SetShippingCourierIDNil

`func (o *ShippingLabelCreateDto) SetShippingCourierIDNil(b bool)`

 SetShippingCourierIDNil sets the value for ShippingCourierID to be an explicit nil

### UnsetShippingCourierID
`func (o *ShippingLabelCreateDto) UnsetShippingCourierID()`

UnsetShippingCourierID ensures that no value is present for ShippingCourierID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



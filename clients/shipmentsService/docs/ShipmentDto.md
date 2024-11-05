# ShipmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**TrackingCode** | Pointer to **NullableString** |  | [optional] 
**IsInternational** | Pointer to **bool** |  | [optional] 
**ShipmentTimestamp** | Pointer to **time.Time** |  | [optional] 
**DeliveryTimestamp** | Pointer to **time.Time** |  | [optional] 
**ExpectedShippingDate** | Pointer to **time.Time** |  | [optional] 
**ExpectedDeliveryDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewShipmentDto

`func NewShipmentDto() *ShipmentDto`

NewShipmentDto instantiates a new ShipmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentDtoWithDefaults

`func NewShipmentDtoWithDefaults() *ShipmentDto`

NewShipmentDtoWithDefaults instantiates a new ShipmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShipmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipmentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShipmentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ShipmentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ShipmentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTrackingCode

`func (o *ShipmentDto) GetTrackingCode() string`

GetTrackingCode returns the TrackingCode field if non-nil, zero value otherwise.

### GetTrackingCodeOk

`func (o *ShipmentDto) GetTrackingCodeOk() (*string, bool)`

GetTrackingCodeOk returns a tuple with the TrackingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCode

`func (o *ShipmentDto) SetTrackingCode(v string)`

SetTrackingCode sets TrackingCode field to given value.

### HasTrackingCode

`func (o *ShipmentDto) HasTrackingCode() bool`

HasTrackingCode returns a boolean if a field has been set.

### SetTrackingCodeNil

`func (o *ShipmentDto) SetTrackingCodeNil(b bool)`

 SetTrackingCodeNil sets the value for TrackingCode to be an explicit nil

### UnsetTrackingCode
`func (o *ShipmentDto) UnsetTrackingCode()`

UnsetTrackingCode ensures that no value is present for TrackingCode, not even an explicit nil
### GetIsInternational

`func (o *ShipmentDto) GetIsInternational() bool`

GetIsInternational returns the IsInternational field if non-nil, zero value otherwise.

### GetIsInternationalOk

`func (o *ShipmentDto) GetIsInternationalOk() (*bool, bool)`

GetIsInternationalOk returns a tuple with the IsInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternational

`func (o *ShipmentDto) SetIsInternational(v bool)`

SetIsInternational sets IsInternational field to given value.

### HasIsInternational

`func (o *ShipmentDto) HasIsInternational() bool`

HasIsInternational returns a boolean if a field has been set.

### GetShipmentTimestamp

`func (o *ShipmentDto) GetShipmentTimestamp() time.Time`

GetShipmentTimestamp returns the ShipmentTimestamp field if non-nil, zero value otherwise.

### GetShipmentTimestampOk

`func (o *ShipmentDto) GetShipmentTimestampOk() (*time.Time, bool)`

GetShipmentTimestampOk returns a tuple with the ShipmentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentTimestamp

`func (o *ShipmentDto) SetShipmentTimestamp(v time.Time)`

SetShipmentTimestamp sets ShipmentTimestamp field to given value.

### HasShipmentTimestamp

`func (o *ShipmentDto) HasShipmentTimestamp() bool`

HasShipmentTimestamp returns a boolean if a field has been set.

### GetDeliveryTimestamp

`func (o *ShipmentDto) GetDeliveryTimestamp() time.Time`

GetDeliveryTimestamp returns the DeliveryTimestamp field if non-nil, zero value otherwise.

### GetDeliveryTimestampOk

`func (o *ShipmentDto) GetDeliveryTimestampOk() (*time.Time, bool)`

GetDeliveryTimestampOk returns a tuple with the DeliveryTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTimestamp

`func (o *ShipmentDto) SetDeliveryTimestamp(v time.Time)`

SetDeliveryTimestamp sets DeliveryTimestamp field to given value.

### HasDeliveryTimestamp

`func (o *ShipmentDto) HasDeliveryTimestamp() bool`

HasDeliveryTimestamp returns a boolean if a field has been set.

### GetExpectedShippingDate

`func (o *ShipmentDto) GetExpectedShippingDate() time.Time`

GetExpectedShippingDate returns the ExpectedShippingDate field if non-nil, zero value otherwise.

### GetExpectedShippingDateOk

`func (o *ShipmentDto) GetExpectedShippingDateOk() (*time.Time, bool)`

GetExpectedShippingDateOk returns a tuple with the ExpectedShippingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedShippingDate

`func (o *ShipmentDto) SetExpectedShippingDate(v time.Time)`

SetExpectedShippingDate sets ExpectedShippingDate field to given value.

### HasExpectedShippingDate

`func (o *ShipmentDto) HasExpectedShippingDate() bool`

HasExpectedShippingDate returns a boolean if a field has been set.

### GetExpectedDeliveryDate

`func (o *ShipmentDto) GetExpectedDeliveryDate() time.Time`

GetExpectedDeliveryDate returns the ExpectedDeliveryDate field if non-nil, zero value otherwise.

### GetExpectedDeliveryDateOk

`func (o *ShipmentDto) GetExpectedDeliveryDateOk() (*time.Time, bool)`

GetExpectedDeliveryDateOk returns a tuple with the ExpectedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDeliveryDate

`func (o *ShipmentDto) SetExpectedDeliveryDate(v time.Time)`

SetExpectedDeliveryDate sets ExpectedDeliveryDate field to given value.

### HasExpectedDeliveryDate

`func (o *ShipmentDto) HasExpectedDeliveryDate() bool`

HasExpectedDeliveryDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



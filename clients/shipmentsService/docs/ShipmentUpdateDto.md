# ShipmentUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingCode** | Pointer to **NullableString** |  | [optional] 
**IsInternational** | Pointer to **bool** |  | [optional] 
**Shipped** | Pointer to **bool** |  | [optional] 
**Delivered** | Pointer to **bool** |  | [optional] 
**ShipmentTimestamp** | Pointer to **time.Time** |  | [optional] 
**DeliveryTimestamp** | Pointer to **time.Time** |  | [optional] 
**ExpectedShippingDate** | Pointer to **time.Time** |  | [optional] 
**ExpectedDeliveryDate** | Pointer to **time.Time** |  | [optional] 
**ShippingTerms** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewShipmentUpdateDto

`func NewShipmentUpdateDto() *ShipmentUpdateDto`

NewShipmentUpdateDto instantiates a new ShipmentUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentUpdateDtoWithDefaults

`func NewShipmentUpdateDtoWithDefaults() *ShipmentUpdateDto`

NewShipmentUpdateDtoWithDefaults instantiates a new ShipmentUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingCode

`func (o *ShipmentUpdateDto) GetTrackingCode() string`

GetTrackingCode returns the TrackingCode field if non-nil, zero value otherwise.

### GetTrackingCodeOk

`func (o *ShipmentUpdateDto) GetTrackingCodeOk() (*string, bool)`

GetTrackingCodeOk returns a tuple with the TrackingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCode

`func (o *ShipmentUpdateDto) SetTrackingCode(v string)`

SetTrackingCode sets TrackingCode field to given value.

### HasTrackingCode

`func (o *ShipmentUpdateDto) HasTrackingCode() bool`

HasTrackingCode returns a boolean if a field has been set.

### SetTrackingCodeNil

`func (o *ShipmentUpdateDto) SetTrackingCodeNil(b bool)`

 SetTrackingCodeNil sets the value for TrackingCode to be an explicit nil

### UnsetTrackingCode
`func (o *ShipmentUpdateDto) UnsetTrackingCode()`

UnsetTrackingCode ensures that no value is present for TrackingCode, not even an explicit nil
### GetIsInternational

`func (o *ShipmentUpdateDto) GetIsInternational() bool`

GetIsInternational returns the IsInternational field if non-nil, zero value otherwise.

### GetIsInternationalOk

`func (o *ShipmentUpdateDto) GetIsInternationalOk() (*bool, bool)`

GetIsInternationalOk returns a tuple with the IsInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternational

`func (o *ShipmentUpdateDto) SetIsInternational(v bool)`

SetIsInternational sets IsInternational field to given value.

### HasIsInternational

`func (o *ShipmentUpdateDto) HasIsInternational() bool`

HasIsInternational returns a boolean if a field has been set.

### GetShipped

`func (o *ShipmentUpdateDto) GetShipped() bool`

GetShipped returns the Shipped field if non-nil, zero value otherwise.

### GetShippedOk

`func (o *ShipmentUpdateDto) GetShippedOk() (*bool, bool)`

GetShippedOk returns a tuple with the Shipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipped

`func (o *ShipmentUpdateDto) SetShipped(v bool)`

SetShipped sets Shipped field to given value.

### HasShipped

`func (o *ShipmentUpdateDto) HasShipped() bool`

HasShipped returns a boolean if a field has been set.

### GetDelivered

`func (o *ShipmentUpdateDto) GetDelivered() bool`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *ShipmentUpdateDto) GetDeliveredOk() (*bool, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *ShipmentUpdateDto) SetDelivered(v bool)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *ShipmentUpdateDto) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetShipmentTimestamp

`func (o *ShipmentUpdateDto) GetShipmentTimestamp() time.Time`

GetShipmentTimestamp returns the ShipmentTimestamp field if non-nil, zero value otherwise.

### GetShipmentTimestampOk

`func (o *ShipmentUpdateDto) GetShipmentTimestampOk() (*time.Time, bool)`

GetShipmentTimestampOk returns a tuple with the ShipmentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentTimestamp

`func (o *ShipmentUpdateDto) SetShipmentTimestamp(v time.Time)`

SetShipmentTimestamp sets ShipmentTimestamp field to given value.

### HasShipmentTimestamp

`func (o *ShipmentUpdateDto) HasShipmentTimestamp() bool`

HasShipmentTimestamp returns a boolean if a field has been set.

### GetDeliveryTimestamp

`func (o *ShipmentUpdateDto) GetDeliveryTimestamp() time.Time`

GetDeliveryTimestamp returns the DeliveryTimestamp field if non-nil, zero value otherwise.

### GetDeliveryTimestampOk

`func (o *ShipmentUpdateDto) GetDeliveryTimestampOk() (*time.Time, bool)`

GetDeliveryTimestampOk returns a tuple with the DeliveryTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTimestamp

`func (o *ShipmentUpdateDto) SetDeliveryTimestamp(v time.Time)`

SetDeliveryTimestamp sets DeliveryTimestamp field to given value.

### HasDeliveryTimestamp

`func (o *ShipmentUpdateDto) HasDeliveryTimestamp() bool`

HasDeliveryTimestamp returns a boolean if a field has been set.

### GetExpectedShippingDate

`func (o *ShipmentUpdateDto) GetExpectedShippingDate() time.Time`

GetExpectedShippingDate returns the ExpectedShippingDate field if non-nil, zero value otherwise.

### GetExpectedShippingDateOk

`func (o *ShipmentUpdateDto) GetExpectedShippingDateOk() (*time.Time, bool)`

GetExpectedShippingDateOk returns a tuple with the ExpectedShippingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedShippingDate

`func (o *ShipmentUpdateDto) SetExpectedShippingDate(v time.Time)`

SetExpectedShippingDate sets ExpectedShippingDate field to given value.

### HasExpectedShippingDate

`func (o *ShipmentUpdateDto) HasExpectedShippingDate() bool`

HasExpectedShippingDate returns a boolean if a field has been set.

### GetExpectedDeliveryDate

`func (o *ShipmentUpdateDto) GetExpectedDeliveryDate() time.Time`

GetExpectedDeliveryDate returns the ExpectedDeliveryDate field if non-nil, zero value otherwise.

### GetExpectedDeliveryDateOk

`func (o *ShipmentUpdateDto) GetExpectedDeliveryDateOk() (*time.Time, bool)`

GetExpectedDeliveryDateOk returns a tuple with the ExpectedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDeliveryDate

`func (o *ShipmentUpdateDto) SetExpectedDeliveryDate(v time.Time)`

SetExpectedDeliveryDate sets ExpectedDeliveryDate field to given value.

### HasExpectedDeliveryDate

`func (o *ShipmentUpdateDto) HasExpectedDeliveryDate() bool`

HasExpectedDeliveryDate returns a boolean if a field has been set.

### GetShippingTerms

`func (o *ShipmentUpdateDto) GetShippingTerms() string`

GetShippingTerms returns the ShippingTerms field if non-nil, zero value otherwise.

### GetShippingTermsOk

`func (o *ShipmentUpdateDto) GetShippingTermsOk() (*string, bool)`

GetShippingTermsOk returns a tuple with the ShippingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTerms

`func (o *ShipmentUpdateDto) SetShippingTerms(v string)`

SetShippingTerms sets ShippingTerms field to given value.

### HasShippingTerms

`func (o *ShipmentUpdateDto) HasShippingTerms() bool`

HasShippingTerms returns a boolean if a field has been set.

### GetOrderId

`func (o *ShipmentUpdateDto) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ShipmentUpdateDto) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ShipmentUpdateDto) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ShipmentUpdateDto) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *ShipmentUpdateDto) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *ShipmentUpdateDto) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



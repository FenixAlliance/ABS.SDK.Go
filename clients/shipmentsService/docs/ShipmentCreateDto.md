# ShipmentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**TrackingCode** | Pointer to **NullableString** |  | [optional] 
**IsInternational** | Pointer to **bool** |  | [optional] 
**ExpectedShippingDate** | Pointer to **time.Time** |  | [optional] 
**ExpectedDeliveryDate** | Pointer to **time.Time** |  | [optional] 
**ShippingTerms** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewShipmentCreateDto

`func NewShipmentCreateDto() *ShipmentCreateDto`

NewShipmentCreateDto instantiates a new ShipmentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentCreateDtoWithDefaults

`func NewShipmentCreateDtoWithDefaults() *ShipmentCreateDto`

NewShipmentCreateDtoWithDefaults instantiates a new ShipmentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShipmentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipmentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipmentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShipmentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ShipmentCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ShipmentCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ShipmentCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ShipmentCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTrackingCode

`func (o *ShipmentCreateDto) GetTrackingCode() string`

GetTrackingCode returns the TrackingCode field if non-nil, zero value otherwise.

### GetTrackingCodeOk

`func (o *ShipmentCreateDto) GetTrackingCodeOk() (*string, bool)`

GetTrackingCodeOk returns a tuple with the TrackingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCode

`func (o *ShipmentCreateDto) SetTrackingCode(v string)`

SetTrackingCode sets TrackingCode field to given value.

### HasTrackingCode

`func (o *ShipmentCreateDto) HasTrackingCode() bool`

HasTrackingCode returns a boolean if a field has been set.

### SetTrackingCodeNil

`func (o *ShipmentCreateDto) SetTrackingCodeNil(b bool)`

 SetTrackingCodeNil sets the value for TrackingCode to be an explicit nil

### UnsetTrackingCode
`func (o *ShipmentCreateDto) UnsetTrackingCode()`

UnsetTrackingCode ensures that no value is present for TrackingCode, not even an explicit nil
### GetIsInternational

`func (o *ShipmentCreateDto) GetIsInternational() bool`

GetIsInternational returns the IsInternational field if non-nil, zero value otherwise.

### GetIsInternationalOk

`func (o *ShipmentCreateDto) GetIsInternationalOk() (*bool, bool)`

GetIsInternationalOk returns a tuple with the IsInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternational

`func (o *ShipmentCreateDto) SetIsInternational(v bool)`

SetIsInternational sets IsInternational field to given value.

### HasIsInternational

`func (o *ShipmentCreateDto) HasIsInternational() bool`

HasIsInternational returns a boolean if a field has been set.

### GetExpectedShippingDate

`func (o *ShipmentCreateDto) GetExpectedShippingDate() time.Time`

GetExpectedShippingDate returns the ExpectedShippingDate field if non-nil, zero value otherwise.

### GetExpectedShippingDateOk

`func (o *ShipmentCreateDto) GetExpectedShippingDateOk() (*time.Time, bool)`

GetExpectedShippingDateOk returns a tuple with the ExpectedShippingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedShippingDate

`func (o *ShipmentCreateDto) SetExpectedShippingDate(v time.Time)`

SetExpectedShippingDate sets ExpectedShippingDate field to given value.

### HasExpectedShippingDate

`func (o *ShipmentCreateDto) HasExpectedShippingDate() bool`

HasExpectedShippingDate returns a boolean if a field has been set.

### GetExpectedDeliveryDate

`func (o *ShipmentCreateDto) GetExpectedDeliveryDate() time.Time`

GetExpectedDeliveryDate returns the ExpectedDeliveryDate field if non-nil, zero value otherwise.

### GetExpectedDeliveryDateOk

`func (o *ShipmentCreateDto) GetExpectedDeliveryDateOk() (*time.Time, bool)`

GetExpectedDeliveryDateOk returns a tuple with the ExpectedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDeliveryDate

`func (o *ShipmentCreateDto) SetExpectedDeliveryDate(v time.Time)`

SetExpectedDeliveryDate sets ExpectedDeliveryDate field to given value.

### HasExpectedDeliveryDate

`func (o *ShipmentCreateDto) HasExpectedDeliveryDate() bool`

HasExpectedDeliveryDate returns a boolean if a field has been set.

### GetShippingTerms

`func (o *ShipmentCreateDto) GetShippingTerms() string`

GetShippingTerms returns the ShippingTerms field if non-nil, zero value otherwise.

### GetShippingTermsOk

`func (o *ShipmentCreateDto) GetShippingTermsOk() (*string, bool)`

GetShippingTermsOk returns a tuple with the ShippingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTerms

`func (o *ShipmentCreateDto) SetShippingTerms(v string)`

SetShippingTerms sets ShippingTerms field to given value.

### HasShippingTerms

`func (o *ShipmentCreateDto) HasShippingTerms() bool`

HasShippingTerms returns a boolean if a field has been set.

### GetOrderId

`func (o *ShipmentCreateDto) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ShipmentCreateDto) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ShipmentCreateDto) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ShipmentCreateDto) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *ShipmentCreateDto) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *ShipmentCreateDto) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



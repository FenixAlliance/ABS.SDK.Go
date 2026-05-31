# ItemPackingSlipDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Instructions** | Pointer to **NullableString** |  | [optional] 
**DeliveryNoteId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**OrderId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemPackingSlipDto

`func NewItemPackingSlipDto() *ItemPackingSlipDto`

NewItemPackingSlipDto instantiates a new ItemPackingSlipDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemPackingSlipDtoWithDefaults

`func NewItemPackingSlipDtoWithDefaults() *ItemPackingSlipDto`

NewItemPackingSlipDtoWithDefaults instantiates a new ItemPackingSlipDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemPackingSlipDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemPackingSlipDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemPackingSlipDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemPackingSlipDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ItemPackingSlipDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ItemPackingSlipDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ItemPackingSlipDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemPackingSlipDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemPackingSlipDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemPackingSlipDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ItemPackingSlipDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ItemPackingSlipDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetInstructions

`func (o *ItemPackingSlipDto) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ItemPackingSlipDto) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ItemPackingSlipDto) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ItemPackingSlipDto) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *ItemPackingSlipDto) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *ItemPackingSlipDto) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetDeliveryNoteId

`func (o *ItemPackingSlipDto) GetDeliveryNoteId() string`

GetDeliveryNoteId returns the DeliveryNoteId field if non-nil, zero value otherwise.

### GetDeliveryNoteIdOk

`func (o *ItemPackingSlipDto) GetDeliveryNoteIdOk() (*string, bool)`

GetDeliveryNoteIdOk returns a tuple with the DeliveryNoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryNoteId

`func (o *ItemPackingSlipDto) SetDeliveryNoteId(v string)`

SetDeliveryNoteId sets DeliveryNoteId field to given value.

### HasDeliveryNoteId

`func (o *ItemPackingSlipDto) HasDeliveryNoteId() bool`

HasDeliveryNoteId returns a boolean if a field has been set.

### SetDeliveryNoteIdNil

`func (o *ItemPackingSlipDto) SetDeliveryNoteIdNil(b bool)`

 SetDeliveryNoteIdNil sets the value for DeliveryNoteId to be an explicit nil

### UnsetDeliveryNoteId
`func (o *ItemPackingSlipDto) UnsetDeliveryNoteId()`

UnsetDeliveryNoteId ensures that no value is present for DeliveryNoteId, not even an explicit nil
### GetTenantId

`func (o *ItemPackingSlipDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ItemPackingSlipDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ItemPackingSlipDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ItemPackingSlipDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ItemPackingSlipDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ItemPackingSlipDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetOrderId

`func (o *ItemPackingSlipDto) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ItemPackingSlipDto) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ItemPackingSlipDto) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ItemPackingSlipDto) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *ItemPackingSlipDto) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *ItemPackingSlipDto) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



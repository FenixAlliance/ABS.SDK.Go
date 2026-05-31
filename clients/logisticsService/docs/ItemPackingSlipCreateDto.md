# ItemPackingSlipCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Instructions** | Pointer to **NullableString** |  | [optional] 
**DeliveryNoteId** | Pointer to **NullableString** |  | [optional] 
**OrderId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemPackingSlipCreateDto

`func NewItemPackingSlipCreateDto() *ItemPackingSlipCreateDto`

NewItemPackingSlipCreateDto instantiates a new ItemPackingSlipCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemPackingSlipCreateDtoWithDefaults

`func NewItemPackingSlipCreateDtoWithDefaults() *ItemPackingSlipCreateDto`

NewItemPackingSlipCreateDtoWithDefaults instantiates a new ItemPackingSlipCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemPackingSlipCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemPackingSlipCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemPackingSlipCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemPackingSlipCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ItemPackingSlipCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemPackingSlipCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemPackingSlipCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemPackingSlipCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetInstructions

`func (o *ItemPackingSlipCreateDto) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ItemPackingSlipCreateDto) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ItemPackingSlipCreateDto) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ItemPackingSlipCreateDto) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *ItemPackingSlipCreateDto) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *ItemPackingSlipCreateDto) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetDeliveryNoteId

`func (o *ItemPackingSlipCreateDto) GetDeliveryNoteId() string`

GetDeliveryNoteId returns the DeliveryNoteId field if non-nil, zero value otherwise.

### GetDeliveryNoteIdOk

`func (o *ItemPackingSlipCreateDto) GetDeliveryNoteIdOk() (*string, bool)`

GetDeliveryNoteIdOk returns a tuple with the DeliveryNoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryNoteId

`func (o *ItemPackingSlipCreateDto) SetDeliveryNoteId(v string)`

SetDeliveryNoteId sets DeliveryNoteId field to given value.

### HasDeliveryNoteId

`func (o *ItemPackingSlipCreateDto) HasDeliveryNoteId() bool`

HasDeliveryNoteId returns a boolean if a field has been set.

### SetDeliveryNoteIdNil

`func (o *ItemPackingSlipCreateDto) SetDeliveryNoteIdNil(b bool)`

 SetDeliveryNoteIdNil sets the value for DeliveryNoteId to be an explicit nil

### UnsetDeliveryNoteId
`func (o *ItemPackingSlipCreateDto) UnsetDeliveryNoteId()`

UnsetDeliveryNoteId ensures that no value is present for DeliveryNoteId, not even an explicit nil
### GetOrderId

`func (o *ItemPackingSlipCreateDto) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ItemPackingSlipCreateDto) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ItemPackingSlipCreateDto) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ItemPackingSlipCreateDto) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *ItemPackingSlipCreateDto) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *ItemPackingSlipCreateDto) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



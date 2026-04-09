# ItemCartRecordCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**CartId** | Pointer to **NullableString** |  | [optional] 
**ProductId** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **float64** |  | [optional] 

## Methods

### NewItemCartRecordCreateDto

`func NewItemCartRecordCreateDto() *ItemCartRecordCreateDto`

NewItemCartRecordCreateDto instantiates a new ItemCartRecordCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemCartRecordCreateDtoWithDefaults

`func NewItemCartRecordCreateDtoWithDefaults() *ItemCartRecordCreateDto`

NewItemCartRecordCreateDtoWithDefaults instantiates a new ItemCartRecordCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemCartRecordCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemCartRecordCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemCartRecordCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemCartRecordCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ItemCartRecordCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemCartRecordCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemCartRecordCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemCartRecordCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetCartId

`func (o *ItemCartRecordCreateDto) GetCartId() string`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *ItemCartRecordCreateDto) GetCartIdOk() (*string, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *ItemCartRecordCreateDto) SetCartId(v string)`

SetCartId sets CartId field to given value.

### HasCartId

`func (o *ItemCartRecordCreateDto) HasCartId() bool`

HasCartId returns a boolean if a field has been set.

### SetCartIdNil

`func (o *ItemCartRecordCreateDto) SetCartIdNil(b bool)`

 SetCartIdNil sets the value for CartId to be an explicit nil

### UnsetCartId
`func (o *ItemCartRecordCreateDto) UnsetCartId()`

UnsetCartId ensures that no value is present for CartId, not even an explicit nil
### GetProductId

`func (o *ItemCartRecordCreateDto) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ItemCartRecordCreateDto) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ItemCartRecordCreateDto) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ItemCartRecordCreateDto) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *ItemCartRecordCreateDto) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *ItemCartRecordCreateDto) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetQuantity

`func (o *ItemCartRecordCreateDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ItemCartRecordCreateDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ItemCartRecordCreateDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ItemCartRecordCreateDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ItemRestockEntryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Quantity** | Pointer to **float64** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**WarehouseId** | Pointer to **NullableString** |  | [optional] 
**ItemRestockId** | Pointer to **NullableString** |  | [optional] 
**OrderItemRecordId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemRestockEntryDto

`func NewItemRestockEntryDto() *ItemRestockEntryDto`

NewItemRestockEntryDto instantiates a new ItemRestockEntryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemRestockEntryDtoWithDefaults

`func NewItemRestockEntryDtoWithDefaults() *ItemRestockEntryDto`

NewItemRestockEntryDtoWithDefaults instantiates a new ItemRestockEntryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemRestockEntryDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemRestockEntryDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemRestockEntryDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemRestockEntryDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ItemRestockEntryDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ItemRestockEntryDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ItemRestockEntryDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemRestockEntryDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemRestockEntryDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemRestockEntryDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ItemRestockEntryDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ItemRestockEntryDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetQuantity

`func (o *ItemRestockEntryDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ItemRestockEntryDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ItemRestockEntryDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ItemRestockEntryDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetItemId

`func (o *ItemRestockEntryDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemRestockEntryDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemRestockEntryDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ItemRestockEntryDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *ItemRestockEntryDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *ItemRestockEntryDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetWarehouseId

`func (o *ItemRestockEntryDto) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *ItemRestockEntryDto) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *ItemRestockEntryDto) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *ItemRestockEntryDto) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### SetWarehouseIdNil

`func (o *ItemRestockEntryDto) SetWarehouseIdNil(b bool)`

 SetWarehouseIdNil sets the value for WarehouseId to be an explicit nil

### UnsetWarehouseId
`func (o *ItemRestockEntryDto) UnsetWarehouseId()`

UnsetWarehouseId ensures that no value is present for WarehouseId, not even an explicit nil
### GetItemRestockId

`func (o *ItemRestockEntryDto) GetItemRestockId() string`

GetItemRestockId returns the ItemRestockId field if non-nil, zero value otherwise.

### GetItemRestockIdOk

`func (o *ItemRestockEntryDto) GetItemRestockIdOk() (*string, bool)`

GetItemRestockIdOk returns a tuple with the ItemRestockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemRestockId

`func (o *ItemRestockEntryDto) SetItemRestockId(v string)`

SetItemRestockId sets ItemRestockId field to given value.

### HasItemRestockId

`func (o *ItemRestockEntryDto) HasItemRestockId() bool`

HasItemRestockId returns a boolean if a field has been set.

### SetItemRestockIdNil

`func (o *ItemRestockEntryDto) SetItemRestockIdNil(b bool)`

 SetItemRestockIdNil sets the value for ItemRestockId to be an explicit nil

### UnsetItemRestockId
`func (o *ItemRestockEntryDto) UnsetItemRestockId()`

UnsetItemRestockId ensures that no value is present for ItemRestockId, not even an explicit nil
### GetOrderItemRecordId

`func (o *ItemRestockEntryDto) GetOrderItemRecordId() string`

GetOrderItemRecordId returns the OrderItemRecordId field if non-nil, zero value otherwise.

### GetOrderItemRecordIdOk

`func (o *ItemRestockEntryDto) GetOrderItemRecordIdOk() (*string, bool)`

GetOrderItemRecordIdOk returns a tuple with the OrderItemRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemRecordId

`func (o *ItemRestockEntryDto) SetOrderItemRecordId(v string)`

SetOrderItemRecordId sets OrderItemRecordId field to given value.

### HasOrderItemRecordId

`func (o *ItemRestockEntryDto) HasOrderItemRecordId() bool`

HasOrderItemRecordId returns a boolean if a field has been set.

### SetOrderItemRecordIdNil

`func (o *ItemRestockEntryDto) SetOrderItemRecordIdNil(b bool)`

 SetOrderItemRecordIdNil sets the value for OrderItemRecordId to be an explicit nil

### UnsetOrderItemRecordId
`func (o *ItemRestockEntryDto) UnsetOrderItemRecordId()`

UnsetOrderItemRecordId ensures that no value is present for OrderItemRecordId, not even an explicit nil
### GetTenantId

`func (o *ItemRestockEntryDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ItemRestockEntryDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ItemRestockEntryDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ItemRestockEntryDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ItemRestockEntryDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ItemRestockEntryDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


